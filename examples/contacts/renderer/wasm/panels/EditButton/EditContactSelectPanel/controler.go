package EditContactSelectPanel

import (
	//"syscall/js"

	"syscall/js"

	"github.com/josephbudd/kicknotjs"
	"github.com/josephbudd/kickwasmwidgets"

	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/data/records"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/wasm/viewtools"
)

/*

	Panel name: EditContactSelectPanel
	Panel id:   tabsMasterView-home-pad-EditButton-EditContactSelectPanel

*/

// Controler is a HelloPanel Controler.
type Controler struct {
	panel     *Panel
	presenter *Presenter
	caller    *Caller
	quitCh    chan struct{}    // send an empty struct to start the quit process.
	tools     *viewtools.Tools // see /renderer/wasm/viewtools
	notjs     *kicknotjs.NotJS

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your Controler members.

	*/

	contactEditSelect   *kickwasmwidgets.VList
	contactEditSelectID uint64
}

// defineControlsSetHandlers defines controler members and sets their handlers.
func (controler *Controler) defineControlsSetHandlers() {

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define the Controler members by their html elements.
	// Set handlers.
	// example:

	// Define controler members.
	notjs := controler.notjs
	controler.addCustomerName := notjs.GetElementByID("addCustomerName")
	controler.addCustomerSubmit := notjs.GetElementByID("addCustomerSubmit")

	// Set handlers.
	cb := notjs.RegisterCallBack(controler.handleSubmit)
	notjs.SetOnClick(controler.addCustomerSubmit, cb)

	*/

	notjs := controler.notjs
	stater := kickwasmwidgets.NewVListState()
	controler.contactEditSelectID = stater.GetNextState()
	controler.contactEditSelect = kickwasmwidgets.NewVList(
		// div
		notjs.GetElementByID("contactEditSelect"),
		// vlist ID from kickwasmwidgets.VListState.GetNextState()
		controler.contactEditSelectID,
		// max buttons
		10,
		// onSizeFunc
		// Called when there are records in the db.
		func() {
			controler.panel.showEditContactSelectPanel(false)
		},
		// onNoSizeFunc
		// Called when there are no records in the db.
		func() {
			controler.panel.showEditContactNotReadyPanel(false)
		},
		// needToInitializeFunc
		// Called when the list needs to initialize itself with all new buttons.
		func(pageSize, state uint64) {
			controler.caller.getContactsPage(0, pageSize, state)
		},
		// needToPrependFunc
		// Called when the user has scrolled to the top of the list.
		// The list needs more to append more buttons to the top of the list.
		func(button js.Value, pageSize, state uint64) {
			sortedIndex := controler.notjs.GetAttributeUint64(button, "vlist_sorted_index")
			if sortedIndex == 0 {
				return
			}
			if sortedIndex < pageSize {
				pageSize = sortedIndex
				sortedIndex = 0
			} else {
				sortedIndex -= pageSize
			}
			controler.caller.getContactsPage(sortedIndex, pageSize, state)
		},
		// needToAppendFunc
		// Called when the user has scrolled to the bottom of the list.
		// The list needs more to append more buttons to the bottom of the list.
		func(button js.Value, pageSize, state uint64) {
			sortedIndex := controler.notjs.GetAttributeUint64(button, "vlist_sorted_index")
			sortedIndex++
			controler.caller.getContactsPage(sortedIndex, pageSize, state)
		},
		// hideFunc
		controler.tools.ElementHide,
		// showFunc
		controler.tools.ElementShow,
		// isShownFunc
		controler.tools.ElementIsShown,
		// notjs
		controler.notjs,
	)
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

*/

func (controler *Controler) handleContactsPage(contacts []*records.ContactRecord, sortedIndex, recordCount, state uint64) {
	notjs := controler.notjs
	l := len(contacts)
	buttons := make([]js.Value, l, l)
	for i, contact := range contacts {
		button := notjs.CreateElement("button")
		notjs.SetAttributeUint64(button, "vlist_sorted_index", sortedIndex)
		sortedIndex++
		notjs.SetAttributeUint64(button, "vlist_record_id", contact.ID)
		b := notjs.CreateElement("b")
		tn := notjs.CreateTextNode(contact.Name)
		notjs.AppendChild(b, tn)
		notjs.AppendChild(button, b)
		cb := notjs.RegisterCallBack(func(args []js.Value) {
			// the user has clicked on a button
			event := args[0]
			target := notjs.GetEventTarget(event)
			for {
				if notjs.TagName(target) == "BUTTON" {
					break
				}
				target = notjs.ParentNode(target)
			}
			id := notjs.GetAttributeUint64(target, "vlist_record_id")
			controler.caller.getContact(id)
			// subpanelIndex := notjs.GetAttributeUint64(target, "vlist_subpanel_index")
			// stateValue := notjs.GetAttribute("vlist_state_value");
			// cityValue := notjs.GetAttribute("vlist_city_value");
			// controler.contactEditSelect.SetSubPanelTitle(subpanelindex+1, "Select a City in ".concat(stateValue, "."));
			// controler.caller.GetContactsPageCategoryCityCB(0, max, stateValue, (subpanelindex+1)|idState);
		})
		notjs.SetOnClick(button, cb)
		buttons[i] = button
	}
	controler.contactEditSelect.Build(buttons, state, recordCount)
}

// initialCalls runs the first code that the controler needs to run.
func (controler *Controler) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.

	*/

	controler.contactEditSelect.Start()

}
