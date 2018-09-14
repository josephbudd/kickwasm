package templates

// CallServerLockedGo is /mainprocess/transports/callserver/locked.go
const CallServerLockedGo = `package callserver

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

func (callServer *Server) incConnectionCount() {
	callServer.connectionCountMutex.Lock()
	if connectionCount == -1 {
		connectionCount = 1
	} else {
		connectionCount++
	}
	callServer.connectionCountMutex.Unlock()
}

func (callServer *Server) decConnectionCount() {
	callServer.connectionCountMutex.Lock()
	if connectionCount > 0 {
		connectionCount--
	}
	callServer.connectionCountMutex.Unlock()
}

// GetConnectionCount returns the connection count.
func (callServer *Server) GetConnectionCount() int {
	callServer.connectionCountMutex.Lock()
	cc := connectionCount
	callServer.connectionCountMutex.Unlock()
	return cc
}

func (callServer *Server) setLastDisconnect(t time.Time) {
	lastDisconnectMutex.Lock()
	lastDisconnect = t
	lastDisconnectMutex.Unlock()
}

// GetLastDisconnect returns the time of the last disconnect.
func (callServer *Server) GetLastDisconnect() time.Time {
	lastDisconnectMutex.Lock()
	t := lastDisconnect
	lastDisconnectMutex.Unlock()
	return t
}

`

// CallServerGo is the /mainprocess/tranpsports/callserver/callserver.go file.
const CallServerGo = `package callserver

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"{{.ApplicationGitPath}}{{.ImportDomainTypes}}"
)

const pongWait = 60 * time.Second

// Server is a main process local procedure call.
type Server struct {
	host          string
	port          uint
	callMap       types.MainProcessCallsMap
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
}

// NewCallServer constructs a new Server.
func NewCallServer(host string, port uint, callMap types.MainProcessCallsMap) *Server {
	return &Server{
		host:          host,
		port:          port,
		callMap:       callMap,
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
				appHost := fmt.Sprintf("%s:%d", host, port)
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
	}
}

`

// CallServerRunGo is /mainprocess/transports/callserver/run.go
const CallServerRunGo = `package callserver

import (
	"fmt"
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
// Param port is the http port like 9090
//
// Package callserver handles all local http requests where either of the following 2 conditions are true.
//  strings.HasPrefix(r.URL.Path, "/ws")
//  r.URL.Path == "/callserver.js"
// All other requests are passed to the parameter handlerFunc.
// You will want your handlerFunc to do things like load your javascript, css and any other files.
// Example HTML:
//  <style> @import url(css/keyboard.css); </style>
func (callServer *Server) Run(handlerFunc http.HandlerFunc) error {
	appurl := fmt.Sprintf("http://%s:%d", callServer.host, callServer.port)
	log.Println("listen and serve: ", appurl)
	myServer = &http.Server{
		Addr: fmt.Sprintf(":%d", callServer.port),
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		callServer.serve(w, r, handlerFunc)
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
	//  1.c myServer.ListenAndServe() ends because the server is closed.
	//
	//  2.a the terminal user types ^c or ^\ and stopRunLoopCh gets the signal.
	//  2.b stillConnectedLoop stops the server and itself.
	//  2.c myServer.ListenAndServe() ends because the server is closed.
	stopRunLoopCh := make(chan os.Signal, 1)
	signal.Notify(stopRunLoopCh, os.Interrupt)
	go callServer.stillConnectedLoop(stopRunLoopCh)
	// start the server
	return myServer.ListenAndServe()
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
//  * serves the callserver javascript file request to func serveJS.
//  * passes others to the handlerFunc func.
func (callServer *Server) serve(w http.ResponseWriter, r *http.Request, handlerFunc http.HandlerFunc) {
	if r.Method == "GET" {
		if strings.HasPrefix(r.URL.Path, "/ws") {
			callServer.serveWebSocket(w, r)
			return
		}
	}
	handlerFunc(w, r)
}

// stillConnectedLoop periodically checks the number of web socket connections.
//  which are set in the func serveWebSocket.
// If there are no connections and the server has been disconnected for too long
//  * it closes myServer and returns
//  * the closed server causes myServer.ListenAndServe() in func Run to end.
//  * which causes func Run to end.
// If there is an operationg system interupt signal
//  * it closes myServer and returns
//  * the closed server causes myServer.ListenAndServe() in func Run to end.
//  * which causes func Run to end.
func (callServer *Server) stillConnectedLoop(stopRunLoopCh chan os.Signal) {
	ticker := time.NewTicker(callServer.DisconnectMax / 2)
	defer ticker.Stop()
	for {
		select {
		case now := <-ticker.C:
			if callServer.GetConnectionCount() == 0 {
				if now.Sub(callServer.GetLastDisconnect()) > callServer.DisconnectMax {
					myServer.Close()
					log.Println("ticker myServer.Close()")
					return
				}
			}
		case <-stopRunLoopCh:
			myServer.Close()
			log.Println("<-stopRunLoopCh myServer.Close()")
			return
		}
	}
}

`

//CallServerWebsocketGo is /mainprocess/transports/callserver/websocket.go
const CallServerWebsocketGo = `package callserver

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"

	"{{.ApplicationGitPath}}{{.ImportDomainTypes}}"
)

// Serve does
//  * serves the web socket requests to func serveWebSocket.
//  * passes others to the handlerFunc func.
func (callServer *Server) Serve(w http.ResponseWriter, r *http.Request, handlerFunc http.HandlerFunc) {
	if r.Method == "GET" {
		if strings.HasPrefix(r.URL.Path, "/ws") {
			callServer.serveWebSocket(w, r)
			return
		}
	}
	handlerFunc(w, r)
}

// serveWebSocket handles a new web socket connection.
// it keeps the web socket connection open until the ping loop or read loop tell it to stop.
// it has a write loop which tells the ping and read loops to stop if there is a write error.
func (callServer *Server) serveWebSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := callServer.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	callServer.incConnectionCount()
	defer func() {
		log.Println("CLOSED WS CONNECTION")
		ws.Close()
		callServer.decConnectionCount()
		callServer.setLastDisconnect(time.Now())
	}()
	log.Println("OPENED WS CONNECTION")
	messageFromSenderCh := make(chan []byte)
	closeWSConnectionCh := make(chan struct{})
	stopReadLoopCh := make(chan struct{}, 1)
	stopPingLoopCh := make(chan struct{}, 1)
	var t time.Time
	ws.SetReadDeadline(t)
	ws.SetWriteDeadline(t)
	go callServer.pingLoop(ws, stopPingLoopCh, closeWSConnectionCh)
	go callServer.readLoop(ws, messageFromSenderCh, stopReadLoopCh, closeWSConnectionCh)
	// this is the write loop
	// it also checks for stop
	for {
		select {
		case bb := <-messageFromSenderCh:
			err := ws.WriteMessage(websocket.TextMessage, bb)
			if err != nil {
				log.Println("serveWebSocket: ws.WriteMessage(websocket.TextMessage, bb) error is ", err.Error())
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
func (callServer *Server) readLoop(ws *websocket.Conn,
	callCh chan []byte,
	stopReadLoopCh chan struct{},
	closeWSConnectionCh chan struct{}) {
	ticker := time.NewTicker(time.Millisecond * 250)
	defer ticker.Stop()
	//var timeout = time.After(time.Millisecond * 250)
	for {
		select {
		case <-ticker.C:
			_, payloadbb, err := ws.ReadMessage()
			//messageType, payloadbb, err := ws.ReadMessage()
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
				if len(payloadbb) == 0 {
					log.Println("readLoop: len(payloadbb) == 0", err)
				} else {
					log.Println("main process incoming message")
					payload := &types.Payload{}
					err := json.Unmarshal(payloadbb, payload)
					if err != nil {
						log.Println("readLoop: json.Unmarshal(payloadbb)", err)
						// don't stop the server because the pingloop will.
						return
					}
					// dispatch the message.
					if lpc, ok := callServer.callMap[payload.Procedure]; ok {
						callBack := func(paramsbb []byte) {
							payload := &types.Payload{
								Params:    string(paramsbb),
								Procedure: payload.Procedure,
							}
							payloadbb, _ := json.Marshal(payload)
							callCh <- payloadbb
						}
						log.Println("calling go lpc.MainProcessReceive")
						go lpc.MainProcessReceive([]byte(payload.Params), callBack)
					} else {
						log.Printf("payload.Procedure %d not found", payload.Procedure)
					}
				}
			}
		case <-stopReadLoopCh:
			// the read loop needs to stop.
			return
		}
	}
}

// pingLoop periodically sends pings.
// if there is an error it sends via closeWSConnectionCh and returns.
// if it receives via stopPingLoopCh it returns.
func (callServer *Server) pingLoop(ws *websocket.Conn, stopPingLoopCh, closeWSConnectionCh chan struct{}) {
	ticker := time.NewTicker(callServer.pingPeriod)
	defer ticker.Stop()
	for {
		select {
		case lastPing := <-ticker.C:
			// send a ping
			if err := ws.WriteControl(websocket.PingMessage, []byte{}, lastPing.Add(callServer.pingPeriod)); err != nil {
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
