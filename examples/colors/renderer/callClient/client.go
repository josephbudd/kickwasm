package call

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/colors/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/colors/domain/types"
	"github.com/josephbudd/kickwasm/examples/colors/site/notjs"
	"github.com/josephbudd/kickwasm/examples/colors/site/viewtools"
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
	OnConnectionBreakJS js.Callback
	OnConnectionBreak   func([]js.Value)
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
func (client *Client) SetOnConnectionBreak(f func([]js.Value)) {
	client.OnConnectionBreak = f
	client.OnConnectionBreakJS = client.notJS.RegisterCallBack(f)
}

func (client *Client) defaultOnConnectionBreak([]js.Value) {
	client.notJS.Alert("The connection to the main process has broken.")
}

// Connect connects to the server.
func (client *Client) Connect(callBack func()) bool {
	notJS := client.notJS
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
	client.connection.Set("onopen", notJS.RegisterCallBack(
		func(args []js.Value) {
			client.onOpen(args)
			callBack()
		}),
	)
	client.connection.Set("onclose", notJS.RegisterCallBack(client.onClose))
	client.connection.Set("onmessage", notJS.RegisterCallBack(client.onMessage))
	return true
}

func (client *Client) onOpen(args []js.Value) {
	client.connected = true
	client.notJS.ConsoleLog("Calls are connected.")
}

func (client *Client) onClose(args []js.Value) {
	client.connected = false
	client.notJS.ConsoleLog("Calls are unconnected.")
	client.OnConnectionBreak(nil)
}

func (client *Client) onMessage(args []js.Value) {
	e := args[0]
	data := e.Get("data").String()
	payload := types.Payload{}
	if err := json.Unmarshal([]byte(data), &payload); err != nil {
		message := fmt.Sprintf("client.onMessage: json.Unmarshal([]byte(data), payload) error is %q.", err.Error())
		client.notJS.Alert(message)
		return
	}
	client.queue = append(client.queue, payload)
	client.dispatch()
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
