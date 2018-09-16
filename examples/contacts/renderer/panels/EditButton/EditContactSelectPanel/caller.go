package EditContactSelectPanel

import (
	"github.com/josephbudd/kicknotjs"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/implementations/calling"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/states"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: EditContactSelectPanel
	Panel id:   tabsMasterView-home-pad-EditButton-EditContactSelectPanel

*/

// Caller communicates with the main process via an asynchrounous connection.
type Caller struct {
	panel      *Panel
	presenter  *Presenter
	controler  *Controler
	quitCh     chan struct{} // send an empty struct to start the quit process.
	connection types.RendererCallMap
	tools      *viewtools.Tools // see /renderer/viewtools
	notjs      *kicknotjs.NotJS
	// my added members
	serviceStates *states.States
	// my calls
	getContactCaller                           caller.Renderer
	getContactsPageStatesCaller                caller.Renderer
	getContactsPageCitiesMatchStateCaller      caller.Renderer
	getContactsPageRecordsMatchStateCityCaller caller.Renderer
}

// setMainProcessCallBacks tells the main process what funcs to call back to.
func (panelCaller *Caller) addMainProcessCallBacks() {

	/* NOTE TO DEVELOPER. Step 1 of 3.

	// Tell the main processs to call back to your funcs.

	*/

	panelCaller.getContactCaller = panelCaller.connection[calling.GetContactCallID]
	panelCaller.getContactCaller.AddCallBack(panelCaller.getContactCB)

	panelCaller.getContactsPageStatesCaller = panelCaller.connection[calling.GetContactsPageStatesCallID]
	panelCaller.getContactsPageStatesCaller.AddCallBack(panelCaller.GetContactsPageStatesCB)

	panelCaller.getContactsPageCitiesMatchStateCaller = panelCaller.connection[calling.GetContactsPageCitiesMatchStateCallID]
	panelCaller.getContactsPageCitiesMatchStateCaller.AddCallBack(panelCaller.GetContactsPageCitiesMatchStateCB)

	panelCaller.getContactsPageRecordsMatchStateCityCaller = panelCaller.connection[calling.GetContactsPageRecordsMatchStateCityCallID]
	panelCaller.getContactsPageRecordsMatchStateCityCaller.AddCallBack(panelCaller.GetContactsPageRecordsMatchStateCityCB)

	panelCaller.connection[calling.UpdateContactCallID].AddCallBack(panelCaller.updateContactCB)
	panelCaller.connection[calling.RemoveContactCallID].AddCallBack(panelCaller.removeContactCB)

}

/* NOTE TO DEVELOPER. Step 2 of 3.

// Define calls to the main process and their and call backs.

*/

// UpdateContact

func (panelCaller *Caller) updateContactCB(params interface{}) {
	// the contacts repo has been changed to restart the contact selector.
	panelCaller.controler.contactEditSelect.Start()
}

// GetContact

// GetContact gets a single contact record.
func (panelCaller *Caller) GetContact(id uint64) {
	params := &calling.RendererToMainProcessGetContactParams{
		ID:    id,
		State: panelCaller.serviceStates.Edit,
	}
	panelCaller.getContactCaller.CallMainProcess(params)
}

func (panelCaller *Caller) getContactCB(params interface{}) {
	switch params := params.(type) {
	case *calling.MainProcessToRendererGetContactParams:
		if params.State&panelCaller.serviceStates.Edit == panelCaller.serviceStates.Edit {
			if params.Error {
				panelCaller.tools.Error(params.ErrorMessage)
			}
			// no error so let the edit panel handle the call back.
		}
	}
}

// Remove Contact

func (panelCaller *Caller) removeContactCB(params interface{}) {
	switch params := params.(type) {
	case *calling.MainProcessToRendererRemoveContactParams:
		if params.Error {
			return
		}
		// the select panelCaller needs to handle this.
		panelCaller.controler.contactEditSelect.Start()
	}
}

// GetContactsPageStates

// GetContactsPageStates gets a page records with unique States.
func (panelCaller *Caller) GetContactsPageStates(sortedIndex, pageSize, state uint64) {
	params := &calling.RendererToMainProcessGetContactsPageStatesParams{
		SortedIndex: sortedIndex,
		PageSize:    pageSize,
		State:       state | panelCaller.serviceStates.Edit,
	}
	panelCaller.getContactsPageStatesCaller.CallMainProcess(params)
}

// GetContactsPageStatesCB handles the main process call back.
func (panelCaller *Caller) GetContactsPageStatesCB(params interface{}) {
	switch params := params.(type) {
	case *calling.MainProcessToRendererGetContactsPageStatesParams:
		if params.State&panelCaller.serviceStates.Edit == panelCaller.serviceStates.Edit {
			if params.Error {
				panelCaller.tools.Error(params.ErrorMessage)
				return
			}
			// ok
			panelCaller.controler.contactEditSelect.Build(params.Records, params.SortedIndex, params.RecordCount, params.State)
		}
	}
}

// GetContactsPageCitiesMatchState

// GetContactsPageCitiesMatchState gets a page of records that match stateMatch.
func (panelCaller *Caller) GetContactsPageCitiesMatchState(sortedIndex, pageSize, state uint64, stateMatch string) {
	params := &calling.RendererToMainProcessGetContactsPageCitiesMatchStateParams{
		SortedIndex: sortedIndex,
		PageSize:    pageSize,
		State:       state | panelCaller.serviceStates.Edit,
		StateMatch:  stateMatch,
	}
	panelCaller.getContactsPageCitiesMatchStateCaller.CallMainProcess(params)
}

// GetContactsPageCitiesMatchStateCB handles the main process call back.
func (panelCaller *Caller) GetContactsPageCitiesMatchStateCB(params interface{}) {
	switch params := params.(type) {
	case *calling.MainProcessToRendererGetContactsPageCitiesMatchStateParams:
		if params.State&panelCaller.serviceStates.Edit == panelCaller.serviceStates.Edit {
			if params.Error {
				panelCaller.tools.Error(params.ErrorMessage)
				return
			}
			// ok
			panelCaller.controler.contactEditSelect.Build(params.Records, params.SortedIndex, params.RecordCount, params.State)
		}
	}
}

// GetContactsPageRecordsMatchStateCity

// GetContactsPageRecordsMatchStateCity gets records with matching cities and states.
func (panelCaller *Caller) GetContactsPageRecordsMatchStateCity(sortedIndex, pageSize, state uint64, stateMatch, cityMatch string) {
	params := &calling.RendererToMainProcessGetContactsPageRecordsMatchStateCityParams{
		SortedIndex: sortedIndex,
		PageSize:    pageSize,
		State:       state | panelCaller.serviceStates.Edit,
		StateMatch:  stateMatch,
		CityMatch:   cityMatch,
	}
	panelCaller.getContactsPageRecordsMatchStateCityCaller.CallMainProcess(params)
}

// GetContactsPageRecordsMatchStateCityCB handles the main process call back.
func (panelCaller *Caller) GetContactsPageRecordsMatchStateCityCB(params interface{}) {
	switch params := params.(type) {
	case *calling.MainProcessToRendererGetContactsPageRecordsMatchStateCityParams:
		if params.State&panelCaller.serviceStates.Edit == panelCaller.serviceStates.Edit {
			if params.Error {
				panelCaller.tools.Error(params.ErrorMessage)
				return
			}
			// ok
			panelCaller.controler.contactEditSelect.Build(params.Records, params.SortedIndex, params.RecordCount, params.State)
		}
	}
}

// initialCalls makes the first calls to the main process.
func (panelCaller *Caller) initialCalls() {

	/* NOTE TO DEVELOPER. Step 3 of 3.

	// Make any initial calls to the main process that must be made when the app starts.

	*/

}
