// +build js, wasm

package lpc

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/push/domain/lpc"
	"github.com/josephbudd/kickwasm/examples/push/domain/lpc/message"
)

/*
	DO NOT EDIT THIS FILE.

	USE THE TOOL kicklpc TO ADD OR REMOVE LPC Messages.

	kicklpc will edit this file for you.

*/

// Sending is a channel that sends to the main process.
type Sending chan interface{}

// Receiving is a channel that receives from the main process.
type Receiving chan interface{}

var (
	send    Sending
	receive Receiving
	global  js.Value
	alert   js.Value
)

func init() {
	send = make(chan interface{}, 1024)
	receive = make(chan interface{}, 1024)
	g := js.Global()
	global = g
	alert = g.Get("alert")
}

// Channels returns the renderer connection channels.
func Channels() (sendChan, receiveChan chan interface{}) {
	sendChan = send
	receiveChan = receive
	return
}

// Payload converts unmarshalled msg to the correct marshalled payload.
func (sending Sending) Payload(msg interface{}) (payload []byte, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("sending.Payload: %w", err)
		}
	}()

	var bb []byte
	var id uint64
	switch msg := msg.(type) {
	case *message.LogRendererToMainProcess:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 0
	case *message.InitRendererToMainProcess:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 1
	case *message.TimeRendererToMainProcess:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 2
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
			err = fmt.Errorf("receiving.Cargo: %w", err)
		}
	}()

	payload := lpc.Payload{}
	if err = json.Unmarshal(payloadbb, &payload); err != nil {
		return
	}
	switch payload.ID {
	case 0:
		msg := &message.LogMainProcessToRenderer{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 1:
		msg := &message.InitMainProcessToRenderer{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 2:
		msg := &message.TimeMainProcessToRenderer{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	default:
		errMsg := fmt.Sprintf("no case found for payload id %d", payload.ID)
		err = fmt.Errorf(errMsg)
	}
	return
}
