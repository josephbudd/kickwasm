package EditContactSelectPanel

import (
	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
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
	connection map[types.CallID]caller.Renderer
	tools      *viewtools.Tools // see /renderer/viewtools
	notJS      *notjs.NotJS

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your Caller members.

	*/

	// my added members
	state uint64
	// my calls
	getContactCaller                           caller.Renderer
	getContactsPageStatesCaller                caller.Renderer
	getContactsPageCitiesMatchStateCaller      caller.Renderer
	getContactsPageRecordsMatchStateCityCaller caller.Renderer
}

// addMainProcessCallBacks tells the main process what funcs to call back to.
func (panelCaller *Caller) addMainProcessCallBacks() {

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define your added Caller members.
	// Tell the main processs to call back to your funcs.

	*/

	panelCaller.getContactCaller = panelCaller.connection[callids.GetContactCallID]
	panelCaller.getContactCaller.AddCallBack(panelCaller.getContactCB)

	panelCaller.getContactsPageStatesCaller = panelCaller.connection[callids.GetContactsPageStatesCallID]
	panelCaller.getContactsPageStatesCaller.AddCallBack(panelCaller.GetContactsPageStatesCB)

	panelCaller.getContactsPageCitiesMatchStateCaller = panelCaller.connection[callids.GetContactsPageCitiesMatchStateCallID]
	panelCaller.getContactsPageCitiesMatchStateCaller.AddCallBack(panelCaller.GetContactsPageCitiesMatchStateCB)

	panelCaller.getContactsPageRecordsMatchStateCityCaller = panelCaller.connection[callids.GetContactsPageRecordsMatchStateCityCallID]
	panelCaller.getContactsPageRecordsMatchStateCityCaller.AddCallBack(panelCaller.GetContactsPageRecordsMatchStateCityCB)

	panelCaller.connection[callids.UpdateContactCallID].AddCallBack(panelCaller.updateContactCB)
	panelCaller.connection[callids.RemoveContactCallID].AddCallBack(panelCaller.removeContactCB)

}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Define calls to the main process and their and call backs.

*/

// UpdateContact

func (panelCaller *Caller) updateContactCB(params interface{}) {
	// the contacts store has been modified to restart the contact selector.
	panelCaller.controler.contactEditSelect.Start()
}

// GetContact

// GetContact gets a single contact record.
func (panelCaller *Caller) GetContact(id uint64) {
	params := &types.RendererToMainProcessGetContactParams{
		ID:    id,
		State: panelCaller.state,
	}
	panelCaller.getContactCaller.CallMainProcess(params)
}

func (panelCaller *Caller) getContactCB(params interface{}) {
	switch params := params.(type) {
	case *types.MainProcessToRendererGetContactParams:
		if params.State&panelCaller.state == panelCaller.state {
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
	case *types.MainProcessToRendererRemoveContactParams:
		if params.Error {
			return
		}
		// the contacts store has been modified to restart the contact selector.
		panelCaller.controler.contactEditSelect.Start()
	}
}

// GetContactsPageStates

// GetContactsPageStates gets a page records with unique States.
func (panelCaller *Caller) GetContactsPageStates(sortedIndex, pageSize, state uint64) {
	params := &types.RendererToMainProcessGetContactsPageStatesParams{
		SortedIndex: sortedIndex,
		PageSize:    pageSize,
		State:       state | panelCaller.state,
	}
	panelCaller.getContactsPageStatesCaller.CallMainProcess(params)
}

// GetContactsPageStatesCB handles the main process call back.
func (panelCaller *Caller) GetContactsPageStatesCB(params interface{}) {
	switch params := params.(type) {
	case *types.MainProcessToRendererGetContactsPageStatesParams:
		if params.State&panelCaller.state == panelCaller.state {
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
	params := &types.RendererToMainProcessGetContactsPageCitiesMatchStateParams{
		SortedIndex: sortedIndex,
		PageSize:    pageSize,
		State:       state | panelCaller.state,
		StateMatch:  stateMatch,
	}
	panelCaller.getContactsPageCitiesMatchStateCaller.CallMainProcess(params)
}

// GetContactsPageCitiesMatchStateCB handles the main process call back.
func (panelCaller *Caller) GetContactsPageCitiesMatchStateCB(params interface{}) {
	switch params := params.(type) {
	case *types.MainProcessToRendererGetContactsPageCitiesMatchStateParams:
		if params.State&panelCaller.state == panelCaller.state {
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
	params := &types.RendererToMainProcessGetContactsPageRecordsMatchStateCityParams{
		SortedIndex: sortedIndex,
		PageSize:    pageSize,
		State:       state | panelCaller.state,
		StateMatch:  stateMatch,
		CityMatch:   cityMatch,
	}
	panelCaller.getContactsPageRecordsMatchStateCityCaller.CallMainProcess(params)
}

// GetContactsPageRecordsMatchStateCityCB handles the main process call back.
func (panelCaller *Caller) GetContactsPageRecordsMatchStateCityCB(params interface{}) {
	switch params := params.(type) {
	case *types.MainProcessToRendererGetContactsPageRecordsMatchStateCityParams:
		if params.State&panelCaller.state == panelCaller.state {
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

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make any initial calls to the main process that must be made when the app starts.

	*/

}
