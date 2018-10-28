package callserver

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"

	"github.com/josephbudd/kickwasm/examples/colors/domain/types"
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
					// get the client and start it to process the message.
					if client, ok := callServer.callMap[payload.Procedure]; ok {
						callBack := func(paramsbb []byte) {
							payload := &types.Payload{
								Params:    string(paramsbb),
								Procedure: payload.Procedure,
							}
							payloadbb, _ := json.Marshal(payload)
							callCh <- payloadbb
						}
						go client.Process([]byte(payload.Params), callBack)
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

