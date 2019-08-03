package lpc

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/spawntabs/domain/lpc"
	"github.com/josephbudd/kickwasm/examples/spawntabs/domain/lpc/message"
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
	default:
		errMsg := fmt.Sprintf("no case found for payload id %d", payload.ID)
		err = errors.New(errMsg)
	}
	return
}

// Signal sends on the eoj channel signaling lpc go funcs to quit.
func (eoj EOJing) Signal() {
	end := struct{}{}
	for i := 0; i < eoj.count; i++ {
		eoj.ch <- end
	}
}

// NewEOJ returns a new eoj channel and increments the usage count.
func (eoj EOJing) NewEOJ() (ch chan struct{}) {
	eoj.count++
	ch = eoj.ch
	return
}

// Release decrements the usage count.
// Call this at the end of your lpc handler func.
func (eoj EOJing) Release() {
	eoj.count--
}
