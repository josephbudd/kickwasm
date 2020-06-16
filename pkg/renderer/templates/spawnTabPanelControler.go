package templates

// SpawnTabPanelController is the genereric renderer panel controller template.
const SpawnTabPanelController = `// +build js, wasm

package {{call .PackageNameCase .PanelName}}

import (
	"context"
	"fmt"

	"{{.ApplicationGitPath}}{{.ImportRendererAPIDOM}}"
)

/*

	Panel name: {{.PanelName}}

*/

// panelController controls user input.
type panelController struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	uniqueID  uint64
	document  *dom.DOM
	panel     *spawnedPanel
	group     *panelGroup
	presenter *panelPresenter
	messenger *panelMessenger

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your panelController members.

	// example:

	// my spawn template has a name input field and a submit button.
	// <label for="addCustomerName{{.SpawnID}}">Name</label><input type="text" id="addCustomerName{{.SpawnID}}">
	// <button id="addCustomerSubmit{{.SpawnID}}">Close</button>

	import "syscall/js"
	import "{{.ApplicationGitPath}}{{.ImportRendererAPIMarkup}}"

	addCustomerName   *markup.Element
	addCustomerSubmit *markup.Element

	*/
}

// defineControlsHandlers defines the GUI's controllers and their event handlers.
// Returns the error.
func (controller *panelController) defineControlsHandlers() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("(controller *panelController) defineControlsHandlers(): %w", err)
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define each controller in the GUI by it's html element.
	// Handle each controller's events.

	// example:

	import "{{.ApplicationGitPath}}{{.ImportRendererAPIDisplay}}"

	var id string

	// Define the customer name input field.
	id = display.SpawnID("addCustomerName{{.SpawnID}}", controller.uniqueID)
	if controller.addCustomerName = contoller.document.ElementByID(id); controller.addCustomerName == nil {
		err = fmt.Errorf("unable to find #" + id)
		return
	}

	// Define the submit button.
	id = display.SpawnID("addCustomerSubmit{{.SpawnID}}", controller.uniqueID)
	if controller.addCustomerSubmit = contoller.document.ElementByID(id); controller.addCustomerSubmit == nil {
		err = fmt.Errorf("unable to find #" + id)
		return
	}
	// Handle the submit button's onclick event.
	controller.addCustomerSubmit.SetEventHandler(controller.handleSubmit, "click", false)

	*/

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

// example:

import "{{.ApplicationGitPath}}{{.ImportDomainStoreRecord}}"
import "{{.ApplicationGitPath}}{{.ImportRendererAPIDisplay}}"
import "{{.ApplicationGitPath}}{{.ImportRendererAPIEvent}}"

func (controller *panelController) handleSubmit(e event.Event) (nilReturn interface{}) {
	// See rendererprocess/api/event/event.go.
	// The event.Event funcs.
	//   e.PreventDefaultBehavior()
	//   e.StopCurrentPhasePropagation()
	//   e.StopAllPhasePropagation()
	//   target := e.JSTarget
	//   event := e.JSEvent
	// You must use the javascript event e.JSEvent, as a js.Value.
	// However, you can use the target as a *markup.Element
	//   target := controller.document.NewElementFromJSValue(e.JSTarget)

	name := strings.TrimSpace(controller.addCustomerName.Value())
	if len(name) == 0 {
		display.Error("Customer Name is required.")
		return
	}
	r := &record.CustomerRecord{
		Name: name,
	}
	controller.messenger.AddCustomer(r)
	return
}

*/

// initialCalls runs the first code that the controller needs to run.
func (controller *panelController) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	// example:

	controller.customerSelectWidget.start()

	*/
}
`
