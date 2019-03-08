package templates

// ClientGo is the renderer/wasm/viewtools/calllclient.go template.
const ClientGo = `package call

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"{{.ApplicationGitPath}}{{.ImportDomainInterfacesCallers}}"
	"{{.ApplicationGitPath}}{{.ImportDomainTypes}}"
	"{{.ApplicationGitPath}}{{.ImportRendererNotJS}}"
	"{{.ApplicationGitPath}}{{.ImportRendererViewTools}}"
)

// Client is a wasm local procedure call client.
type Client struct {
	host        string
	port        uint64
	location    string
	tools       *viewtools.Tools
	notJS       *notjs.NotJS
	connection  js.Value
	connected   bool
	dispatching bool
	queue       []types.Payload
	callMap     map[types.CallID]caller.Renderer
	initialCB   func()

	// handlers
	OnConnectionBreakJS js.Func
	OnConnectionBreak   func(this js.Value, args []js.Value) interface{}
}

// NewClient costructs a new Client.
func NewClient(host string, port uint64, tools *viewtools.Tools, notJS *notjs.NotJS) *Client {
	v := &Client{
		host:     host,
		port:     port,
		location: fmt.Sprintf("ws://%s:%d/ws", host, port),
		tools:    tools,
		notJS:    notJS,
		queue:    make([]types.Payload, 0, 10),
	}
	// handlers
	v.SetOnConnectionBreak(v.defaultOnConnectionBreak)
	return v
}

// SetCallMap sets the callMap and callsStruct.
func (client *Client) SetCallMap(callMap map[types.CallID]caller.Renderer) {
	client.callMap = callMap
}

// SetOnConnectionBreak set the handler for the connection break.
func (client *Client) SetOnConnectionBreak(fn func(this js.Value, args []js.Value) interface{}) {
	client.OnConnectionBreak = fn
	client.OnConnectionBreakJS = client.tools.RegisterCallBack(fn)
}

func (client *Client) defaultOnConnectionBreak(this js.Value, args []js.Value) interface{} {
	client.notJS.Alert("The connection to the main process has broken.")
	return nil
}

// Connect connects to the server.
func (client *Client) Connect(callBack func()) bool {
	notJS := client.notJS
	tools := client.tools
	if client.connected {
		return true
	}
	// setup the web socket
	ws := client.tools.Global.Get("WebSocket")
	client.connection = ws.New(client.location)
	if client.connection == js.Undefined() {
		notJS.ConsoleLog("client.connection is undefined")
		return false
	}
	rs := client.connection.Get("readyState")
	notJS.ConsoleLog(fmt.Sprintf("readyState is %s", rs.String()))
	if rs.String() == "undefined" {
		return false
	}
	client.connection.Set("onopen", tools.RegisterCallBack(
		func(this js.Value, args []js.Value) interface{} {
			client.connected = true
			client.notJS.ConsoleLog("Calls are connected.")
			callBack()
			return nil
		}),
	)
	client.connection.Set("onclose", tools.RegisterCallBack(client.onClose))
	client.connection.Set("onmessage", tools.RegisterCallBack(client.onMessage))
	return true
}

func (client *Client) onClose(this js.Value, args []js.Value) interface{} {
	client.connected = false
	client.notJS.ConsoleLog("Calls are unconnected.")
	client.OnConnectionBreak(js.Undefined(), nil)
	return nil
}

func (client *Client) onMessage(this js.Value, args []js.Value) interface{} {
	e := args[0]
	data := e.Get("data").String()
	payload := types.Payload{}
	if err := json.Unmarshal([]byte(data), &payload); err != nil {
		message := fmt.Sprintf("client.onMessage: json.Unmarshal([]byte(data), payload) error is %q.", err.Error())
		client.notJS.Alert(message)
		return nil
	}
	client.queue = append(client.queue, payload)
	client.dispatch()
	return nil
}

func (client *Client) dispatch() {
	if client.dispatching {
		return
	}
	client.dispatching = len(client.queue) > 0
	for client.dispatching {
		payload := client.queue[0]
		client.queue = client.queue[1:]
		call, found := client.callMap[payload.Procedure]
		if !found {
			message := fmt.Sprintf("No CB found for procedure %d.", payload.Procedure)
			client.notJS.Alert(message)
			return
		}
		call.Dispatch([]byte(payload.Params))
		client.dispatching = len(client.queue) > 0
	}
}

// SendPayload sends the payload to the connection.
func (client *Client) SendPayload(payload []byte) error {
	client.connection.Call("send", string(payload))
	return nil
}
`
