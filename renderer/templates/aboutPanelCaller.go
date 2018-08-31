package templates

// AboutPanelCaller is the main process about panel caller template.
const AboutPanelCaller = `package AboutPanel

import (
	"{{.ApplicationGitPath}}{{.ImportMainProcessTransportsCalls}}"
	"{{.ApplicationGitPath}}{{.ImportRendererWASMViewTools}}"
)

// Caller communicates with the main process.
type Caller struct {
	presenter *Presenter
	quitCh    chan struct{}
	callsStruct  *calls.Calls
	tools     *viewtools.Tools
}

func (caller *Caller) addCallBacks() {
	caller.callsStruct.GetAbout.AddCallBack(caller.getAboutCB)
}

func (caller *Caller) initialCalls() {
	caller.callsStruct.GetAbout.CallMainProcess(nil)
}

func (caller *Caller) getAboutCB(params interface{}) {
	switch params := params.(type) {
	case *calls.MainProcessToRendererGetAboutParams:
		if params.Error {
			caller.tools.Error(params.ErrorMessage)
			return
		}
		caller.presenter.displayReleases(params.Version, params.Releases)
		caller.presenter.displayContributors(params.Contributors)
		caller.presenter.displayCredits(params.Credits)
		caller.presenter.displayLicenses(params.Licenses)
	}
}
`
