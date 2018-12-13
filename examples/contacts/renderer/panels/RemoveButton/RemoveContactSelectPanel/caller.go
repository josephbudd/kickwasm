package RemoveContactSelectPanel

import (
	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/callids"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/interfaces/caller"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	Panel name: RemoveContactSelectPanel

*/

// Caller communicates with the main process via an asynchrounous connection.
type Caller struct {
	panelGroup *PanelGroup
	presenter  *Presenter
	controler  *Controler
	quitCh     chan struct{} // send an empty struct to start the quit process.
	connection map[types.CallID]caller.Renderer
	tools      *viewtools.Tools // see /renderer/viewtools
	notJS      *notjs.NotJS

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// 1: Declare your Caller members.


	*/

	// my added members
	state uint64
	// callers
	getContactCaller                           caller.Renderer
	getContactsPageStatesCaller                caller.Renderer
	getContactsPageCitiesMatchStateCaller      caller.Renderer
	getContactsPageRecordsMatchStateCityCaller caller.Renderer
}

// addMainProcessCallBacks tells the main process what funcs to call back to.
func (panelCaller *Caller) addMainProcessCallBacks() (err error) {
	defer func() {
		// close and check for the error
		if err != nil {
			err = errors.WithMessage(err, "(panelCaller *Caller) addMainProcessCallBacks()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// 2.1: Define each one of your added Caller members.
	// 2.2: Tell the main processs to add a call back to each of your call back funcs.

	*/

	var found bool
	var cllr caller.Renderer

	if panelCaller.getContactCaller, found = panelCaller.connection[callids.GetContactCallID]; !found {
		err = errors.New(`unable to find panelCaller.connection[callids.GetContactCallID]`)
		return
	}
	panelCaller.getContactCaller.AddCallBack(panelCaller.getContactCB)

	if panelCaller.getContactsPageStatesCaller, found = panelCaller.connection[callids.GetContactsPageStatesCallID]; !found {
		err = errors.New(`unable to find panelCaller.connection[callids.GetContactsPageStatesCallID]`)
		return
	}
	panelCaller.getContactsPageStatesCaller.AddCallBack(panelCaller.GetContactsPageStatesCB)

	if panelCaller.getContactsPageCitiesMatchStateCaller, found = panelCaller.connection[callids.GetContactsPageCitiesMatchStateCallID]; !found {
		err = errors.New(`unable to find panelCaller.connection[callids.GetContactsPageCitiesMatchStateCallID]`)
		return
	}
	panelCaller.getContactsPageCitiesMatchStateCaller.AddCallBack(panelCaller.GetContactsPageCitiesMatchStateCB)

	if panelCaller.getContactsPageRecordsMatchStateCityCaller, found = panelCaller.connection[callids.GetContactsPageRecordsMatchStateCityCallID]; !found {
		err = errors.New(`unable to find panelCaller.connection[callids.GetContactsPageRecordsMatchStateCityCallID]`)
		return
	}
	panelCaller.getContactsPageRecordsMatchStateCityCaller.AddCallBack(panelCaller.GetContactsPageRecordsMatchStateCityCB)

	if cllr, found = panelCaller.connection[callids.UpdateContactCallID]; !found {
		err = errors.New(`unable to find panelCaller.connection[callids.UpdateContactCallID]`)
		return
	}
	cllr.AddCallBack(panelCaller.updateContactCB)

	if cllr, found = panelCaller.connection[callids.RemoveContactCallID]; !found {
		err = errors.New(`unable to find panelCaller.connection[callids.RemoveContactCallID]`)
		return
	}
	cllr.AddCallBack(panelCaller.removeContactCB)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// 3.1: Define your funcs which call to the main process.
// 3.2: Define your funcs which the main process calls back to.

*/

// UpdateContact

func (panelCaller *Caller) updateContactCB(params interface{}) {
	// the contacts store has been modified to restart the contact selector.
	panelCaller.controler.contactRemoveSelect.Start()
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
			// no error so let the confirm panel handle the call back.
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
		panelCaller.controler.contactRemoveSelect.Start()
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
			panelCaller.controler.contactRemoveSelect.Build(params.Records, params.SortedIndex, params.RecordCount, params.State)
		}
	}
}

// GetContactsPageCitiesMatchState

// GetContactsPageCitiesMatchState gets a page of records with unique cities that match stateMatch.
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
			panelCaller.controler.contactRemoveSelect.Build(params.Records, params.SortedIndex, params.RecordCount, params.State)
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
			panelCaller.controler.contactRemoveSelect.Build(params.Records, params.SortedIndex, params.RecordCount, params.State)
		}
	}
}

// initialCalls makes the first calls to the main process.
func (panelCaller *Caller) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	//4: Make any initial calls to the main process that must be made when the app starts.

	*/

}
