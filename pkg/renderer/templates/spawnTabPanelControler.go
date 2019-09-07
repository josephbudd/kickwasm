package templates

// SpawnTabPanelController is the genereric renderer panel controller template.
const SpawnTabPanelController = `package {{call .PackageNameCase .PanelName}}

import (
	"syscall/js"

	"github.com/pkg/errors"

	"{{.ApplicationGitPath}}{{.ImportRendererViewTools}}"
)

/*

	Panel name: {{.PanelName}}

*/

// panelController controls user input.
type panelController struct {
	uniqueID  uint64
	panel     *spawnedPanel
	group     *panelGroup
	presenter *panelPresenter
	caller    *panelCaller
	eventCh   chan viewtools.Event
	unSpawningCh chan struct{}
	unspawn   func() error

	/* NOTE TO DEVELOPER. Step 1 of 5.

	// Declare your panelController members.
	// example:

	// my spawn template has a name input field and a submit button.
	// <label for="addCustomerName{{.SpawnID}}">Name</label><input type="text" id="addCustomerName{{.SpawnID}}">
	// <button id="addCustomerSubmit{{.SpawnID}}">Close</button>

	addCustomerName   js.Value
	addCustomerSubmit js.Value

	*/
}

// defineControlsReceiveEvents defines controller members and starts receiving their events.
// Returns the error.
func (controller *panelController) defineControlsReceiveEvents() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(controller *panelController) defineControlsReceiveEvents()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 5.

	// Define the controller members by their html elements.
	// Receive their events.
	// example:

	// import "fmt"

	var id string

	// Define the customer name input field.
	id = tools.FixSpawnID("addCustomerName{{.SpawnID}}", controller.uniqueID)
	if controller.addCustomerName = notJS.GetElementByID(id); controller.addCustomerName == null {
		err = errors.New("unable to find #" + id)
		return
	}

	// Define the submit button.
	id = tools.FixSpawnID("addCustomerSubmit{{.SpawnID}}", controller.uniqueID)
	if controller.addCustomerSubmit = notJS.GetElementByID(id); controller.addCustomerSubmit == null {
		err = errors.New("unable to find #" + id)
		return
	}
	// Receive the submit button's onclick event.
	controller.receiveEvent(controller.handleSubmit, "onclick", false, false, false)

	*/

	return
}

/* NOTE TO DEVELOPER. Step 3 of 5.

// Handlers and other functions.
// example:

// import "{{.ApplicationGitPath}}{{.ImportDomainStoreRecord}}"

func (controller *panelController) handleSubmit(event js.Value) {
	name := strings.TrimSpace(notJS.GetValue(controller.addCustomerName))
	if len(name) == 0 {
		tools.Error("Customer Name is required.")
		return
	}
	r := &record.CustomerRecord{
		Name: name,
	}
	controller.caller.AddCustomer(r)
}

*/

// dispatchEvents dispatches events from the controls.
// It stops when it receives on the eoj channel.
func (controller *panelController) dispatchEvents() {
	go func() {
		var event viewtools.Event
		for {
			select {
			case <-eojCh:
				return
			case <-controller.unSpawningCh:
				return
			case event = <-controller.eventCh:
				// An event that this controller is receiving from one of its members.
				switch event.Target {

				/* NOTE TO DEVELOPER. Step 4 of 5.

				// 4.1.a: Add a case for each controller member
				//          that you are receiving events for.
				// 4.1.b: In that case statement, pass the event to your event handler.

				// example:

				case controller.addCustomerSubmit:
					if event.On == "onclick" {
						controller.handleSubmit(event.Event)
					}

				*/

				}
			}
		}
	}()

	return
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

func (controller *panelController) receiveEvent(element js.Value, event string, preventDefault, stopPropagation, stopImmediatePropagation bool) {
	tools.SendSpawnEvent(controller.eventCh, element, event, preventDefault, stopPropagation, stopImmediatePropagation, controller.uniqueID)
}
`
