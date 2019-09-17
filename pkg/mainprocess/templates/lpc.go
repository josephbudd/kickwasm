package templates

// LPCChannelsGo is /mainprocess/lpc/channels.go
const LPCChannelsGo = `{{ $Dot := . }}package lpc

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/pkg/errors"

	"{{.ApplicationGitPath}}{{.ImportDomainLPC}}"
	"{{.ApplicationGitPath}}{{.ImportDomainLPCMessage}}"
)

// Sending is a chanel that sends to the renderer.
type Sending chan interface{}

// Receiving is a channel that receives from the renderer.
type Receiving chan interface{}

// EOJer signals lpc go process to quit.
type EOJer interface {
	Signal()
	NewEOJ() (ch chan struct{})
	Release()
}

// EOJing is has a channel with a dynami size.
// The channel signals go routines to stop.
// EOJing implements EOJer.
type EOJing struct {
	ch    chan struct{}
	count int
	signaled    bool
	countMutex  *sync.Mutex
	signalMutex *sync.Mutex
}

var (
	send    Sending
	receive Receiving
	eoj     EOJer
)

func init() {
	send = make(chan interface{}, 1024)
	receive = make(chan interface{})
	eoj = &EOJing{
		ch:    make(chan struct{}, 1024),
		count: 0,
		countMutex:  &sync.Mutex{},
		signalMutex: &sync.Mutex{},
	}
}

// Channels returns the renderer connection channels.
func Channels() (sendChan Sending, receiveChan Receiving, eojChan EOJer) {
	sendChan = send
	receiveChan = receive
	eojChan = eoj
	return
}

// Payload converts unmarshalled msg to the correct marshalled payload.
func (sending Sending) Payload(msg interface{}) (payload []byte, err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "sending.Payload")
		}
	}()

	var bb []byte
	var id uint64
	switch msg := msg.(type) {
	case *message.LogMainProcessToRenderer:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 0
	case *message.InitMainProcessToRenderer:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 1{{ range $index, $name := .LPCNames }}
case *message.{{ $name }}MainProcessToRenderer:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = {{ call $Dot.Inc2 $index }}{{ end }}
	default:
		bb = []byte("Unknown!")
		id = 999
	}
	pl := &lpc.Payload{
		ID:    id,
		Cargo: bb,
	}
	payload, err = json.Marshal(pl)
	return
}

// Cargo returns a marshalled payload's unmarshalled cargo.
func (receiving Receiving) Cargo(payloadbb []byte) (cargo interface{}, err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "receiving.Cargo")
		}
	}()

	payload := lpc.Payload{}
	if err = json.Unmarshal(payloadbb, &payload); err != nil {
		return
	}
	switch payload.ID {
	case 0:
		msg := &message.LogRendererToMainProcess{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 1:
		msg := &message.InitRendererToMainProcess{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg{{ range $index, $name := .LPCNames }}
	case {{ call $Dot.Inc2 $index }}:
		msg := &message.{{ $name }}RendererToMainProcess{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg{{ end }}
	default:
		errMsg := fmt.Sprintf("no case found for payload id %d", payload.ID)
		err = errors.New(errMsg)
	}
	return
}

// Signal sends on the eoj channel signaling lpc go funcs to quit.
func (eoj *EOJing) Signal() {
	eoj.signalMutex.Lock()
	if !eoj.signaled {
		eoj.signaled = true
		end := struct{}{}
		for i := 0; i < eoj.count; i++ {
			eoj.ch <- end
		}
	}
	eoj.signalMutex.Unlock()
}

// NewEOJ returns a new eoj channel and increments the usage count.
func (eoj *EOJing) NewEOJ() (ch chan struct{}) {
	eoj.countMutex.Lock()
	eoj.count++
	ch = eoj.ch
	eoj.countMutex.Unlock()
	return
}

// Release decrements the usage count.
// Call this at the end of your lpc handler func.
func (eoj *EOJing) Release() {
	eoj.countMutex.Lock()
	if eoj.count > 0 {
		eoj.count--
	}
	eoj.countMutex.Unlock()
}
`

// LPCLockedGo is /mainprocess/lpc/locked.go
const LPCLockedGo = `package lpc

import (
	"sync"
	"time"
)

// these are vars that need to be locked to avoid data races.
var (
	connectionCount = -1

	lastDisconnectMutex = &sync.Mutex{}
	// when the last web socket connection was closed.
	lastDisconnect = time.Now()
)

func (server *Server) incConnectionCount() {
	server.connectionCountMutex.Lock()
	if connectionCount == -1 {
		connectionCount = 1
	} else {
		connectionCount++
	}
	server.connectionCountMutex.Unlock()
}

func (server *Server) decConnectionCount() {
	server.connectionCountMutex.Lock()
	if connectionCount > 0 {
		connectionCount--
	}
	server.connectionCountMutex.Unlock()
}

// GetConnectionCount returns the connection count.
func (server *Server) GetConnectionCount() int {
	server.connectionCountMutex.Lock()
	cc := connectionCount
	server.connectionCountMutex.Unlock()
	return cc
}

func (server *Server) setLastDisconnect(t time.Time) {
	lastDisconnectMutex.Lock()
	lastDisconnect = t
	lastDisconnectMutex.Unlock()
}

// GetLastDisconnect returns the time of the last disconnect.
func (server *Server) GetLastDisconnect() time.Time {
	lastDisconnectMutex.Lock()
	t := lastDisconnect
	lastDisconnectMutex.Unlock()
	return t
}
`

// LPCServerGo is the /mainprocess/tranpsports/callserver/callserver.go file.
const LPCServerGo = `package lpc

import (
	"log"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const pongWait = 60 * time.Second

// Server is a main process local client call.
type Server struct {
	listener      net.Listener
	DisconnectMax time.Duration

	connectionCountMutex *sync.Mutex
	// the number of web socket connections
	connectionCount int

	lastDisconnectMutex *sync.Mutex
	// when the last web socket connection was closed.
	lastDisconnect time.Time

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod time.Duration

	upgrader websocket.Upgrader

	SendChan    Sending
	ReceiveChan Receiving
	QuitChan    chan struct{}
}

// NewServer constructs a new Server.
func NewServer(listener net.Listener, quitCh chan struct{}, receiving Receiving, sending Sending) *Server {
	return &Server{
		listener:      listener,
		DisconnectMax: time.Millisecond * 500,

		connectionCountMutex: &sync.Mutex{},
		// the number of web socket connections
		connectionCount: -1,

		lastDisconnectMutex: &sync.Mutex{},
		// when the last web socket connection was closed.
		lastDisconnect: time.Now(),

		// Send pings to peer with this period. Must be less than pongWait.
		pingPeriod: (pongWait * 9) / 10,

		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				values, found := r.Header["Origin"]
				if !found {
					log.Println("required oringin header not found")
					return false
				}
				appHost := listener.Addr().String()
				for _, value := range values {
					loc, err := url.Parse(value)
					if err == nil {
						if loc.Host == appHost {
							return true
						}
					}
				}
				return false
			},
		},

		SendChan:    sending,
		ReceiveChan: receiving,
		QuitChan:    quitCh,
	}
}
`

// LPCRunGo is /mainprocess/lpc/run.go
const LPCRunGo = `package lpc

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"time"
)

var (
	myServer *http.Server
)

// Run runs the application until the main process terminates.
// The main process detects when the browser window closes and then stops.
// The browser window also closes if the main process stops.
//
// Package callserver handles all local http requests where either of the following 2 conditions are true.
//  strings.HasPrefix(r.URL.Path, "/ws")
//  r.URL.Path == "/callserver.js"
// All other requests are passed to the message handlerFunc.
// Param: handlerFunc http.HandlerFunc
//        You will want your handlerFunc to do things like load your javascript, css and any other files.
func (server *Server) Run(handlerFunc http.HandlerFunc) error {
	appurl := "http://" + server.listener.Addr().String()
	log.Println("listen and serve: ", appurl)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		server.serve(w, r, handlerFunc)
	})
	// start the server
	go func() {
		if waitServer(appurl) && startBrowser(appurl) {
			log.Printf("A browser window should open. If not, please visit %s", appurl)
		} else {
			log.Printf("Please open your web browser and visit %s", appurl)
		}
	}()
	// start the still connected loop.
	// 2 func ending possibilities:
	//  1.a user closed browser window causing a connection error.
	//  1.b stillConnectedLoop detects the error, stops the server and itself.
	//  1.c myServer.Serve(server.listener) ends because the server is closed.
	//
	//  2.a the terminal user types ^c or ^\ and stopRunLoopCh gets the signal.
	//  2.b stillConnectedLoop stops the server and itself.
	//  2.c myServer.Serve(server.listener) ends because the server is closed.
	stopRunLoopCh := make(chan os.Signal, 1)
	signal.Notify(stopRunLoopCh, os.Interrupt)
	go server.stillConnectedLoop(stopRunLoopCh)
	// start the server
	myServer = &http.Server{}
	return myServer.Serve(server.listener)
}

// startBrowser tries to open the URL in a browser, and returns
// whether it succeed.
// This code is from google's godoc application.
func startBrowser(url string) bool {
	// try to start the browser
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	args2 := append(args[1:], url)
	log.Println(args[0], strings.Join(args2, ", "))
	cmd := exec.Command(args[0], args2...)
	return cmd.Start() == nil
}

// waitServer waits some time for the http Server to start
// serving url. The return value reports whether it starts.
// This code is from google's godoc application.
func waitServer(url string) bool {
	tries := 20
	for tries > 0 {
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
			return true
		}
		time.Sleep(100 * time.Millisecond)
		tries--
	}
	return false
}

// serve
//  * serves the web socket requests to func serveWebSocket.
//  * passes others to the handlerFunc func.
func (server *Server) serve(w http.ResponseWriter, r *http.Request, handlerFunc http.HandlerFunc) {
	// Only method GET allowed:
	// * GET + path "/ws" initializes the web sockets.
	// * Other GETs are not allowed.
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/ws") {
		// This is the initial web socket call from the renderer.
		// It will start the main process web socket server, opening the web socket connection.
		// The renderer and main process will send messages back and forth through this single connection.
		// See domain/renderer/client.go func NewClient.
		server.serveWebSocket(w, r)
		return
	}
	// Other requests.
	// See mainprocess/server.go
	handlerFunc(w, r)
}

// stillConnectedLoop periodically checks the number of web socket connections.
//  which are set in the func serveWebSocket.
// If there are no connections and the server has been disconnected for too long
//  * it closes myServer and returns
//  * the closed server causes myServer.ListenAndServe() in func Run to end.
//  * which causes func Run to end.
//  * An empty struct is sent on the quit chan signaling func main to stop the listening lpc go funcs.
// If there is an operationg system interrupt signal
//  * it closes myServer and returns
//  * the closed server causes myServer.ListenAndServe() in func Run to end.
//  * which causes func Run to end.
//  * An empty struct is sent on the quit chan signaling func main to stop the listening lpc go funcs.
func (server *Server) stillConnectedLoop(stopRunLoopCh chan os.Signal) {
	ticker := time.NewTicker(server.DisconnectMax / 2)
	defer ticker.Stop()
	for {
		select {
		case now := <-ticker.C:
			if server.GetConnectionCount() == 0 {
				if now.Sub(server.GetLastDisconnect()) > server.DisconnectMax {
					myServer.Close()
					log.Println("Ending server. Renderer disconnected.")
					server.QuitChan <- struct{}{}
					return
				}
			}
		case <-stopRunLoopCh:
			myServer.Close()
			log.Println("Ending server. Main process canceled.")
			server.QuitChan <- struct{}{}
			return
		}
	}
}
`

//LPCWebsocketGo is /mainprocess/lpc/websocket.go
const LPCWebsocketGo = `package lpc

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// Serve does
//  * serves the web socket requests to func serveWebSocket.
//  * passes others to the handlerFunc func.
func (server *Server) Serve(w http.ResponseWriter, r *http.Request, handlerFunc http.HandlerFunc) {
	if r.Method == "GET" {
		if strings.HasPrefix(r.URL.Path, "/ws") {
			server.serveWebSocket(w, r)
			return
		}
	}
	handlerFunc(w, r)
}

// serveWebSocket handles a new web socket connection.
// it keeps the web socket connection open until the ping loop or read loop tell it to stop.
// it has a write loop which tells the ping and read loops to stop if there is a write error.
func (server *Server) serveWebSocket(w http.ResponseWriter, r *http.Request) {
	var err error
	var ws *websocket.Conn
	if ws, err = server.upgrader.Upgrade(w, r, nil); err != nil {
		return
	}
	server.incConnectionCount()
	defer func() {
		log.Println("CLOSED WS CONNECTION")
		ws.Close()
		server.decConnectionCount()
		server.setLastDisconnect(time.Now())
	}()
	log.Println("OPENED WS CONNECTION")
	closeWSConnectionCh := make(chan struct{})
	stopReadLoopCh := make(chan struct{}, 1)
	stopPingLoopCh := make(chan struct{}, 1)
	var t time.Time
	ws.SetReadDeadline(t)
	ws.SetWriteDeadline(t)
	go server.pingLoop(ws, stopPingLoopCh, closeWSConnectionCh)
	go server.readLoop(ws, stopReadLoopCh, closeWSConnectionCh)
	// this is the write loop
	// it also checks for stop
	for {
		select {
		case cargo := <-server.SendChan:
			var payload []byte
			if payload, err = server.SendChan.Payload(cargo); err != nil {
				log.Println("serveWebSocket: server.SendChan.Payload(cargo) error is ", err.Error())
				continue
			}
			if err = ws.WriteMessage(websocket.TextMessage, payload); err != nil {
				log.Println("serveWebSocket: ws.WriteMessage(websocket.TextMessage, payload) error is ", err.Error())
			}
		case <-closeWSConnectionCh:
			// the ping loop says that the connection is has problems or is closed.
			// stop the read and ping loops.
			stopReadLoopCh <- struct{}{}
			stopPingLoopCh <- struct{}{}
			return
		}
	}
}

// readLoop periodically checks for incoming messages.
// it sends messages to the message dispatcher.
// if there is an error it sends via closeWSConnectionCh and returns.
// if it receives via stopReadLoopCh it returns.
func (server *Server) readLoop(ws *websocket.Conn,
	stopReadLoopCh chan struct{},
	closeWSConnectionCh chan struct{}) {
	ticker := time.NewTicker(time.Millisecond * 250)
	defer ticker.Stop()
	firstRun := true
	for {
		select {
		case <-ticker.C:
			_, payload, err := ws.ReadMessage()
			if err != nil && err != io.EOF {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
					log.Printf("readloop: unexpected close error is %s", err.Error())
				} else {
					log.Printf("readloop: error is %s", err.Error())
				}
				closeWSConnectionCh <- struct{}{}
				return
			}
			if err != io.EOF {
				if len(payload) == 0 {
					log.Println("readLoop: len(payload) == 0", err)
				} else {
					// send this payload to the receive channel.
					var cargo interface{}
					if cargo, err = server.ReceiveChan.Cargo(payload); err != nil {
						log.Println("readLoop error: ", err)
						return
					}
					server.ReceiveChan <- cargo
				}
			}
			if firstRun {
				firstRun = false
				ticker = time.NewTicker(time.Millisecond * 250)
			}
		case <-stopReadLoopCh:
			// the read loop needs to stop.
			log.Println("read loop: <-stopReadLoopCh")
			return
		}
	}
}

// pingLoop periodically sends pings.
// if there is an error it sends via closeWSConnectionCh and returns.
// if it receives via stopPingLoopCh it returns.
func (server *Server) pingLoop(ws *websocket.Conn, stopPingLoopCh, closeWSConnectionCh chan struct{}) {
	ticker := time.NewTicker(server.pingPeriod)
	defer ticker.Stop()
	for {
		select {
		case lastPing := <-ticker.C:
			// send a ping
			if err := ws.WriteControl(websocket.PingMessage, []byte{}, lastPing.Add(server.pingPeriod)); err != nil {
				// tell the server to stop and return
				closeWSConnectionCh <- struct{}{}
				return
			}
		case <-stopPingLoopCh:
			return
		}
	}
}
`
