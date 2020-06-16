// +build js, wasm

package lpc

import (
	"container/list"
	"context"
	"fmt"
	"log"
	"syscall/js"

	"github.com/josephbudd/kickwasm/proofs/tp2b/tp2btest/domain/lpc/message"
	"github.com/josephbudd/kickwasm/proofs/tp2b/tp2btest/rendererprocess/framework/callback"
	"github.com/josephbudd/kickwasm/proofs/tp2b/tp2btest/rendererprocess/framework/viewtools"
)

/*
	DO NOT EDIT THIS FILE.

	USE THE TOOL kicklpc TO ADD OR REMOVE LPC Messages.

	kicklpc will edit this file for you.

*/

// Queue is a queue.
type Queue struct {
	queue *list.List
}

func newQueue() (queue *Queue) {
	queue = &Queue{
		queue: list.New(),
	}
	return
}

func (q *Queue) length() (l int) {
	l = q.queue.Len()
	return
}

func (q *Queue) push(value []byte) {
	q.queue.PushBack(value)
}

func (q *Queue) pop() (value []byte, found bool) {
	if q.queue.Len() == 0 {
		return
	}
	found = true
	e := q.queue.Front()
	value = e.Value.([]byte)
	q.queue.Remove(e)
	return
}

type Client struct {
	host           string
	port           uint64
	location       string
	connection     js.Value
	connected      bool
	dispatching    bool
	queue          *Queue
	OnOpenCallBack func()
	SendChan    Sending
	ReceiveChan Receiving
	ctx         context.Context
	ctxCancel   context.CancelFunc
	lpcing bool
}

// NewClient costructs a new Client.
func NewClient(ctx context.Context, ctxCancel context.CancelFunc, host string, port uint64, receiving Receiving, sending Sending) *Client {
	v := &Client{
		host:        host,
		port:        port,
		location:    fmt.Sprintf("ws://%s:%d/ws", host, port),
		queue:       newQueue(),
		SendChan:    sending,
		ReceiveChan: receiving,
		ctx:         ctx,
		ctxCancel:   ctxCancel,
	}
	return v
}

// Connect connects to the server.
func (client *Client) Connect(callBack func()) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("client.Connect: %w", err)
		}
	}()

	client.OnOpenCallBack = callBack
	if client.connected {
		log.Println("client is connected")
	}
	// setup the web socket
	ws := global.Get("WebSocket")
	client.connection = ws.New(client.location)
	if client.connection.IsUndefined() {
		err = fmt.Errorf("connection is undefined")
		return
	}
	rs := client.connection.Get("readyState")
	if rs.String() == "undefined" {
		err = fmt.Errorf("readystate is undefined")
		return
	}
	client.connection.Set("onopen", callback.RegisterCallBack(client.onOpen))
	client.connection.Set("onclose", callback.RegisterCallBack(client.onClose))
	client.connection.Set("onmessage", callback.RegisterCallBack(client.onMessage))
	return
}

func (client *Client) dispatch() {
	if client.dispatching {
		return
	}
	var payload []byte
	payload, client.dispatching = client.queue.pop()
	for client.dispatching {
		var cargo interface{}
		var err error
		if cargo, err = client.ReceiveChan.Cargo(payload); err != nil {
			alert.Invoke(err.Error())
			return
		}
		// Trap fatals from the main process and stop.
		switch msg := cargo.(type) {
		case *message.LogMainProcessToRenderer:
			if msg.Fatal {
				viewtools.GoModal(msg.ErrorMessage, "Fatal Error", client.onFatal)
				return
			}
		case *message.InitMainProcessToRenderer:
			if msg.Fatal {
				viewtools.GoModal(msg.ErrorMessage, "Fatal Error", client.onFatal)
				return
			}
		}
		// No fatals so go on.
		panelsCount := viewtools.CountMarkupPanels()
		for i := 0; i < panelsCount; i++ {
			client.ReceiveChan <- cargo
		}
		payload, client.dispatching = client.queue.pop()
	}
}

// Handlers.

func (client *Client) onFatal() {
	client.ctxCancel()
}

func (client *Client) onOpen(this js.Value, args []js.Value) (nilReturn interface{}) {
	client.connected = true
	log.Println("LPC has connected.")
	if client.lpcing {
		return
	}
	client.lpcing = true
	log.Println("starting LPC go routine.")
	go func() {
		for {
			select {
			case <-client.ctx.Done():
				client.lpcing = false
				return
			case cargo := <-client.SendChan:
				log.Println("will send lpc cargo to main process")
				var payload []byte
				var err error
				if payload, err = client.SendChan.Payload(cargo); err != nil {
					log.Printf("sending.Payload(cargo) error is %s", err.Error())
				} else {
					log.Println("payload is " + string(payload))
					client.connection.Call("send", string(payload))
				}
			}
		}
	}()
	client.OnOpenCallBack()
	return
}

func (client *Client) onClose(this js.Value, args []js.Value) (nilReturn interface{}) {
	client.connected = false
	log.Println("LPC has disconnected.")
	client.ctxCancel()
	return
}

func (client *Client) onMessage(this js.Value, args []js.Value) (nilReturn interface{}) {
	e := args[0]
	data := e.Get("data").String()
	client.queue.push([]byte(data))
	client.dispatch()
	return
}