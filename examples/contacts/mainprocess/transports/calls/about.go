package calls

import (
	"encoding/json"
	"log"

	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/services/about"
)

// MainProcessToRendererGetAboutParams are the GetAbout function parameters that the main process sends to the renderer.
type MainProcessToRendererGetAboutParams struct {
	Error        bool
	ErrorMessage string
	Version      string
	Releases     map[string]map[string][]string
	Contributors map[string]string
	Credits      []about.Credit
	Licenses     []about.License
}

func newGetAboutLPC(rendererSendPayload func(payload []byte) error) *LPC {
	return newLPC(
		func(params []byte, call func([]byte)) {
			mainProcessReceiveGetAbout(params, call)
		},
		rendererReceiveAndDispatchGetAbout,
		rendererSendPayload,
	)
}

func rendererCallMainProcessGetAbout(params interface{}, call func([]byte)) error {
	call([]byte{})
	return nil
}

func mainProcessReceiveGetAbout(params []byte, callBack func(params []byte)) {
	log.Println("mainProcessReceiveGetAbout")
	txparams := &MainProcessToRendererGetAboutParams{
		Version:      about.String(),
		Releases:     about.GetReleases(),
		Contributors: about.GetContributors(),
		Credits:      about.GetCredits(),
		Licenses:     about.GetLicenses(),
	}
	txparamsbb, _ := json.Marshal(txparams)
	callBack(txparamsbb)
}

func rendererReceiveAndDispatchGetAbout(params []byte, dispatch func(interface{})) {
	rxparams := &MainProcessToRendererGetAboutParams{}
	if err := json.Unmarshal(params, rxparams); err != nil {
		rxparams = &MainProcessToRendererGetAboutParams{
			Error:        true,
			ErrorMessage: err.Error(),
		}
	}
	dispatch(rxparams)
}

/*

	Here is what a call to the renderer call should look like.

	import 	"github.com/josephbudd/wasmattempt/mainprocess/calls"

	func (caller *Caller) setCallBacks() {
		caller.lpc.GetAbout.AddCallBack(caller.GetAboutCB)
	}

	// GetAbout a message
	func (caller *Caller) GetAboutFatal(message string) {
		params := &calls.RendererToMainProcessGetAboutParams{
			Type: calls.GetAboutTypeFatal,
			Message: message,
		}
		if err := caller.lpc.GetAbout.Call(params); err != nil {
			caller.panelGroup.Error(err.Error())
		}
	}

	// GetAboutCB GetAbout call back from the main process.
	func (caller *Caller) GetAboutCB(params interface{}) {
		switch params.(type) {
		case *calls.MainProcessToRendererGetAboutParams:
			if params.Error {
				caller.panelGroup.Error(params.ErrorMessage)
				return
			}
			caller.panelGroup.Success(params.S)
		}
	}

*/
