package lpc

import (
	"encoding/json"
	"fmt"

	"github.com/josephbudd/kickwasm/examples/spawnwidgets/domain/lpc"
	"github.com/josephbudd/kickwasm/examples/spawnwidgets/domain/lpc/message"
)

// Sending is a chanel that sends to the renderer.
type Sending chan interface{}

// Receiving is a channel that receives from the renderer.
type Receiving chan interface{}

var (
	send    Sending
	receive Receiving
)

func init() {
	send = make(chan interface{}, 1024)
	receive = make(chan interface{})
}

// Channels returns the renderer connection channels.
func Channels() (sendChan Sending, receiveChan Receiving) {
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
	case *message.LogMainProcessToRenderer:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 0
	case *message.InitMainProcessToRenderer:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 1
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
		cargo = msg
	default:
		err = fmt.Errorf("no case found for payload id %d", payload.ID)
	}
	return
}
