package calls

import (
	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/behavior/repoi"
)

/*

	TODO:

	1. Complete the definition of the Calls struct.
	   example:
		   AddCustomer *LPC

	2. In func newCalls, complete the inline construction of &Calls{...}
	   example:
	       AddCustomer: newAddCustomerLPC(customerRepo, rendererSendPayload),

    3. In func newCallsMap, complete the construction of callsMap.
	   example:
       callsMap[callsStruct.AddCustomer.ID] = callsStruct.AddCustomer

*/

// Calls is the calls between the main process and the renderer.
// TODO: you need to add your procedure names to this struct.
// example: GetCustomer *LPC
type Calls struct {
	Log             *LPC
	GetAbout        *LPC
	UpdateContact   *LPC
	GetContact      *LPC
	GetContactsPage *LPC
}

// newCalls constructs a new Calls
// TODO: You need to complete the inline construction of &Calls
// example:
//   GetCustomer: newGetCustomerLPC(rendererSendPayload),
func newCalls(contactRepo repoi.ContactRepoI, rendererSendPayload func(params []byte) error) *Calls {
	lpcID = 0
	return &Calls{
		Log:             newLogLPC(rendererSendPayload),
		GetAbout:        newGetAboutLPC(rendererSendPayload),
		UpdateContact:   newUpdateContactLPC(contactRepo, rendererSendPayload),
		GetContact:      newGetContactLPC(contactRepo, rendererSendPayload),
		GetContactsPage: newGetContactsPageLPC(contactRepo, rendererSendPayload),
	}
}

// newCallsMap constructs new Calls for the renderer.
// TODO: You need to complete the construction of callsMap.
// example:
//   callsMap[callsStruct.GetCustomer.ID] = callsStruct.GetCustomer
func newCallsMap(callsStruct *Calls) map[int]*LPC {
	callsMap := make(map[int]*LPC)
	// build the map from Calls
	callsMap[callsStruct.Log.ID] = callsStruct.Log
	callsMap[callsStruct.GetAbout.ID] = callsStruct.GetAbout
	callsMap[callsStruct.UpdateContact.ID] = callsStruct.UpdateContact
	callsMap[callsStruct.GetContactsPage.ID] = callsStruct.GetContactsPage
	return callsMap
}

// NewCallsAndMap constructs a new Calls and a new map of LPC ids matched to their LPC
func NewCallsAndMap(contactRepo repoi.ContactRepoI, rendererSendPayload func(params []byte) error) (callsStruct *Calls, callsMap map[int]*LPC) {
	// make the calls
	callsStruct = newCalls(contactRepo, rendererSendPayload)
	// make the map
	callsMap = newCallsMap(callsStruct)
	// return both
	return
}
