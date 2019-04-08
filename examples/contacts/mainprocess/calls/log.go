package calls

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/loglevels"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/implementations/calling"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
)

func newLogCall() *calling.MainProcess {
	return calling.NewMainProcess(
		callids.LogCallID,
		// main process receive
		func(params []byte, call func([]byte)) {
			processLog(params, call)
		},
	)
}

func processLog(params []byte, callBackToRenderer func(params []byte)) {
	rxparams := &types.RendererToMainProcessLogCallParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		log.Println("processLog error is ", err.Error())
		message := fmt.Sprintf("mainProcessLog: json.Unmarshal(params, rxparams): error is %s\n", err.Error())
		txparams := &types.MainProcessToRendererLogCallParams{
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
	}
	switch rxparams.Level {
	case loglevels.LogLevelInfo:
		log.Println("Renderer Log: Info: ", rxparams.Message)
	case loglevels.LogLevelWarning:
		log.Println("Renderer Log: Warning: ", rxparams.Message)
	case loglevels.LogLevelError:
		log.Println("Renderer Log: Error: ", rxparams.Message)
	case loglevels.LogLevelFatal:
		log.Println("Renderer Log: Fatal: ", rxparams.Message)
	default:
		message := "processLog: Error unknown rxparams.Level"
		log.Println(message)
		txparams := &types.MainProcessToRendererLogCallParams{
			Level:        rxparams.Level,
			Error:        true,
			ErrorMessage: message,
		}
		txparamsbb, _ := json.Marshal(txparams)
		callBackToRenderer(txparamsbb)
		return

	}
	txparams := &types.MainProcessToRendererLogCallParams{
		Level: rxparams.Level,
	}
	txparamsbb, _ := json.Marshal(txparams)
	callBackToRenderer(txparamsbb)
}

/*

	Here is an example of a panel's caller calling the mainprocess' "Log" procedure.

	import (
		"github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
		"github.com/josephbudd/kickwasm/examples/contacts/domain/implementations/calling"

	)

	func (panelCaller *Caller) setCallBacks() {
		logger := panelCaller.connections[callerids.LogCallID]
		logger.AddCallBack(panelCaller.LogCB)
	}

	// Log a fatal message.
	func (panelCaller *Caller) LogFatal(message string) {
		params := &types.RendererToMainProcessLogParams{
			Level:   loglevels.LogLevelFatal,
			Message: message,
		}
		logger := panelCaller.connections[callerids.LogCallID]
		logger.CallMainProcess(params)
	}

	// LogCB Log call back from the main process.
	func (panelCaller *Caller) LogCB(params interface{}) {
		switch params := params.(type) {
		case *calling.MainProcessToRendererLogCallParams:
			if params.Error {
				panelCaller.tools.Error(params.ErrorMessage)
				return
			}
		}
	}

*/
