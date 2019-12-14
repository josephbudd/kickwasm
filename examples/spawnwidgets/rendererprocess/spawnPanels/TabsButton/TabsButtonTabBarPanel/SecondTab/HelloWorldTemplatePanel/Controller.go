// +build js, wasm

package helloworldtemplatepanel

import (
	"errors"
	"fmt"

	"github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/api/display"
	"github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/api/dom"
	"github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/api/event"
	"github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/api/markup"
	"github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/widgets"
)

/*

	Panel name: HelloWorldTemplatePanel

*/

// panelController controls user input.
type panelController struct {
	uniqueID  uint64
	document  *dom.DOM
	panel     *spawnedPanel
	group     *panelGroup
	presenter *panelPresenter
	messenger *panelMessenger
	unspawn   func() error

	/* NOTE TO DEVELOPER. Step 1 of 5.

	// Declare your panelController members.

	// example:

	// my spawn template has a name input field and a submit button.
	// <label for="addCustomerName{{.SpawnID}}">Name</label><input type="text" id="addCustomerName{{.SpawnID}}">
	// <button id="addCustomerSubmit{{.SpawnID}}">Close</button>

	import "syscall/js"
	import "github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/api/markup"

	addCustomerName   *markup.Element
	addCustomerSubmit *markup.Element

	*/

	widget *widgets.Button
}

// defineControlsHandlers defines the GUI's controllers and their event handlers.
// Returns the error.
func (controller *panelController) defineControlsHandlers() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("(controller *panelController) defineControlsHandlers(): %w", err)
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 5.

	// Define each controller in the GUI by it's html element.
	// Handle each controller's events.

	// example:

	import "github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/api/display"

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

	// The button widget handles it's own events not this controller.
	id := display.SpawnID("widgetWrapper{{.SpawnID}}", controller.uniqueID)
	var widgetWrapper *markup.Element
	if widgetWrapper = controller.document.ElementByID(id); widgetWrapper == nil {
		err = errors.New("unable to find #" + id)
		return
	}
	controller.widget = widgets.SpawnButton(controller.document, widgetWrapper, "Close", controller.handleClick)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 5.

// Handlers and other functions.

// example:

import "github.com/josephbudd/kickwasm/examples/spawnwidgets/domain/store/record"
import "github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/api/event"
import "github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/api/display"

func (controller *panelController) handleSubmit(e event.Event) (nilReturn interface{}) {
	// See renderer/event/event.go.
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

// handleClick unspawns this tab.
func (controller *panelController) handleClick(e event.Event) (nilInterface interface{}) {
	// Unspawn this panel's tab and all of it's panels.
	if err := controller.unspawn(); err != nil {
		display.Error(err.Error())
	}
	return
}

func (controller *panelController) UnSpawning() {

	/* NOTE TO DEVELOPER. Step 4 of 5.

	// This func is called when this tab and it's panels are in the process of unspawning.
	// So if you have some cleaning up to do then do it now.
	//
	// For example if you have a widget that needs to be unspawned
	//   because maybe it has a go routine running that needs to be stopped
	//   then do it here.

	// example:

	controller.myWidget.UnSpawn()

	*/

	// Unspawn the widget.
	controller.widget.UnSpawn()
}

// initialCalls runs the first code that the controller needs to run.
func (controller *panelController) initialCalls() {

	/* NOTE TO DEVELOPER. Step 5 of 5.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	// example:

	controller.customerSelectWidget.start()

	*/
}
