package calling

import (
	"encoding/json"
	"fmt"
	"log"
)

// The Log call id.
var LogCallID = nextCallID()

// Log types are the type message that is logged.
const (
	LogTypeNil int = iota
	LogTypeInfo
	LogTypeWarning
	LogTypeError
	LogTypeFatal
)

// The call Log must have 2 params.
// 1. RendererToMainProcessLogParams
// 2. MainProcessToRendererLogParams

// RendererToMainProcessLogParams are the Log function parameters that the renderer sends to the main process.
type RendererToMainProcessLogParams struct {
	Type    int
	Message string
}

// MainProcessToRendererLogParams are the Log function parameters that the main process sends to the renderer.
type MainProcessToRendererLogParams struct {
	Error        bool
	ErrorMessage string
	Type         int
}

// The Log call must have a constructor of type CallConstructor.

func newLogCall(rendererSendPayload func(payload []byte) error) *Call {
	return newCall(
		LogCallID,
		// main process receive
		func(params []byte, call func([]byte)) {
			mainProcessReceiveLog(params, call)
		},
		// renderer receive dispatch
		func(params []byte, dispatch func(interface{})) {
			rxparams := &MainProcessToRendererLogParams{}
			if err := json.Unmarshal(params, rxparams); err != nil {
				rxparams = &MainProcessToRendererLogParams{
					Error:        true,
					ErrorMessage: err.Error(),
				}
			}
			dispatch(rxparams)
		},
		rendererSendPayload,
	)
}

func mainProcessReceiveLog(params []byte, callBackToRenderer func(params []byte)) {
	log.Println("mainProcessReceiveLog")
	rxparams := &RendererToMainProcessLogParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		log.Println("mainProcessReceiveLog error is ", err.Error())
		message := fmt.Sprintf("mainProcessLog: json.Unmarshal(params, rxparams): error is %s\n", err.Error())
		txparams := &MainProcessToRendererLogParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
	}
	switch rxparams.Type {
	case LogTypeInfo:
		log.Println("Renderer Log: Info: ", rxparams.Message)
	case LogTypeWarning:
		log.Println("Renderer Log: Warning: ", rxparams.Message)
	case LogTypeError:
		log.Println("Renderer Log: Error: ", rxparams.Message)
	case LogTypeFatal:
		log.Println("Renderer Log: Fatal: ", rxparams.Message)
	default:
		message := "mainProcessReceiveLog: Error unknown rxparams.Type"
		log.Println(message)
		txparams := &MainProcessToRendererLogParams{
			Type:         rxparams.Type,
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return

	}
	txparams := &MainProcessToRendererLogParams{
		Type: rxparams.Type,
	}
	txparamsbb, _ := json.Marshal(txparams)
	callBackToRenderer(txparamsbb)
}

func rendererReceiveAndDispatchLog(params []byte, dispatch func(interface{})) {
	rxparams := &MainProcessToRendererLogParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		rxparams = &MainProcessToRendererLogParams{
			Error:        true,
			ErrorMessage: err.Error(),
		}
	}
	dispatch(rxparams)
}

/*

	Here is an example of a panel's caller calling the mainprocess' "Log" procedure.

	import 	"github.com/josephbudd/kickwasm/examples/contacts/domain/implementations/calling"

	func (caller *Caller) setCallBacks() {
		logger := caller.connections[calling.LogCallID]
		logger.AddCallBack(caller.LogCB)
	}

	// Log a fatal message.
	func (caller *Caller) LogFatal(message string) {
		params := &calling.RendererToMainProcessLogParams{
			Type: calling.LogTypeFatal,
			Message: message,
		}
		logger := caller.connections[calling.LogCallID]
		logger.CallMainProcess(params)
	}

	// LogCB Log call back from the main process.
	func (caller *Caller) LogCB(params interface{}) {
		switch params := params.(type) {
		case *calling.MainProcessToRendererLogParams:
			if params.Error {
				caller.tools.Error(params.ErrorMessage)
				return
			}
		}
	}

*/

