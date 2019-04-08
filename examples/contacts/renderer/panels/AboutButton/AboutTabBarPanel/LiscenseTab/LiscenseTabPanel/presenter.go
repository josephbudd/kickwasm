package liscensetabpanel

import (
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: LiscenseTabPanel

*/

// Presenter writes to the panel
type Presenter struct {
	panelGroup *PanelGroup
	controler  *Controler
	caller     *Caller
	tools      *viewtools.Tools // see /renderer/viewtools
	notJS      *notjs.NotJS

	/* NOTE TO DEVELOPER: Step 1 of 3.

	// Declare your Presenter members here.

	*/

	aboutLicense js.Value
}

// defineMembers defines the Presenter members by their html elements.
// Returns the error.
func (panelPresenter *Presenter) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(panelPresenter *Presenter) defineMembers()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 3.

	// Define your Presenter members.

	*/

	notJS := panelPresenter.notJS
	null := js.Null()

	// Define the license div.
	if panelPresenter.aboutLicense = notJS.GetElementByID("aboutLicense"); panelPresenter.aboutLicense == null {
		err = errors.New("unable to find #aboutLicense")
		return
	}

	return
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your Presenter functions.

*/

// displayLicense displays the license in the panel.
func (panelPresenter *Presenter) displayLicense(license []string) {
	notJS := panelPresenter.notJS
	for _, s := range license {
		p := notJS.CreateElementP()
		tn := notJS.CreateTextNode(s)
		notJS.AppendChild(p, tn)
		notJS.AppendChild(panelPresenter.aboutLicense, p)
	}
}
