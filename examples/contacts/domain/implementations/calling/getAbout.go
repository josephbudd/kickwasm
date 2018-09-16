package calling

import (
	"encoding/json"
	"log"

	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/services/about"
)

// GetAboutCallID is the GetAbout call.
var GetAboutCallID = nextCallID()

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

func newGetAboutCall(rendererSendPayload func(payload []byte) error) *Call {
	return newCall(
		GetAboutCallID,
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

	For renderer code see github.com/josephbudd/kickwasm/examples/contacts/home/user1/kick/output/contacts/renderer/panels/AboutButton/AboutPanel/caller.go

*/
