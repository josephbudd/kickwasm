package abouttabpanel

import (
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: AboutTabPanel

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

	aboutAuthor  js.Value
	aboutVersion js.Value
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

	// Define the author paragraph.
	if panelPresenter.aboutAuthor = notJS.GetElementByID("aboutAuthor"); panelPresenter.aboutAuthor == null {
		err = errors.New("unable to find #aboutAuthor")
		return
	}

	// Define the version paragraph.
	if panelPresenter.aboutVersion = notJS.GetElementByID("aboutVersion"); panelPresenter.aboutVersion == null {
		err = errors.New("unable to find #aboutVersion")
		return
	}

	return
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your Presenter functions.

*/

func (panelPresenter *Presenter) displayAbout(author string, version []string) {
	notJS := panelPresenter.notJS
	// author
	tn := notJS.CreateTextNode(author)
	notJS.AppendChild(panelPresenter.aboutAuthor, tn)
	// version
	l := len(version) - 1
	for i, s := range version {
		tn = notJS.CreateTextNode(s)
		notJS.AppendChild(panelPresenter.aboutVersion, tn)
		if i < l {
			br := notJS.CreateElementBR()
			notJS.AppendChild(panelPresenter.aboutVersion, br)
		}
	}

}
