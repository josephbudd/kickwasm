package lpc

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
