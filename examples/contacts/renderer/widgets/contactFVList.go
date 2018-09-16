package widgets

import (
	"fmt"
	"syscall/js"

	"github.com/josephbudd/kicknotjs"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/kickwasmwidgets"
)

// button attributes
const (
	sortedIndexAttributeName = "vlist_sorted_index"
	recordIDAttributeName    = "vlist_record_id"
	stateValueAttributeName  = "state_value"
	cityValueAttributeName   = "city_value"
)

// ContactGetter calls for pages of contact records.
type ContactGetter interface {
	GetContactsPageStates(sortedIndex, pageSize, state uint64)
	GetContactsPageCitiesMatchState(sortedIndex, pageSize, state uint64, stateMatch string)
	GetContactsPageRecordsMatchStateCity(sortedIndex, pageSize, state uint64, stateMatch, cityMatch string)
	GetContact(id uint64)
}

// ContactFVList is an FVList
type ContactFVList struct {
	*kickwasmwidgets.FVList
	getter     ContactGetter
	stateMatch string
	cityMatch  string
}

// NewContactFVList constructs a new FVList for contacts.
func NewContactFVList(div js.Value,
	onSizeFunc func(),
	onNoSizeFunc func(),
	hideFunc func(js.Value),
	showFunc func(js.Value),
	isShownFunc func(js.Value) bool,
	notjs *kicknotjs.NotJS,
	getter ContactGetter,
) *ContactFVList {
	stater := kickwasmwidgets.NewVListState()
	fvlist := kickwasmwidgets.NewFVList(div,
		stater.GetNextState(),
		onSizeFunc,
		onNoSizeFunc,
		hideFunc,
		showFunc,
		isShownFunc,
		notjs,
	)
	cfvlist := &ContactFVList{
		FVList: fvlist,
		getter: getter,
	}
	// add the carousel panel and vlist for the first filter - state
	cfvlist.FVList.AddFirstFilter("State", 10,
		// needToInitializeFunc
		func(count, state uint64) {
			getter.GetContactsPageStates(0, count, state)
		},
		// needToPrependFunc
		func(button js.Value, count, state uint64) {
			sortedIndex := notjs.GetAttributeUint64(button, sortedIndexAttributeName)
			if count > sortedIndex {
				return
			}
			getter.GetContactsPageStates(sortedIndex-count, count, state)
		},
		// needToAppendFunc
		func(button js.Value, count, state uint64) {
			sortedIndex := notjs.GetAttributeUint64(button, sortedIndexAttributeName)
			getter.GetContactsPageStates(sortedIndex+1, count, state)
		},
	)
	// add the carousel panel and vlist for the second filter - city
	cfvlist.FVList.AddAnotherFilter(10,
		// needToInitializeFunc
		func(count, state uint64) {
			getter.GetContactsPageCitiesMatchState(0, count, state, cfvlist.stateMatch)
		},
		// needToPrependFunc
		func(button js.Value, count, state uint64) {
			sortedIndex := notjs.GetAttributeUint64(button, sortedIndexAttributeName)
			getter.GetContactsPageCitiesMatchState(sortedIndex-count, count, state, cfvlist.stateMatch)
		},
		// needToAppendFunc
		func(button js.Value, count, state uint64) {
			sortedIndex := notjs.GetAttributeUint64(button, sortedIndexAttributeName)
			getter.GetContactsPageCitiesMatchState(sortedIndex+1, count, state, cfvlist.stateMatch)
		},
	)
	// add the carousel panel and vlist for the record list
	cfvlist.FVList.AddRecordList(10,
		// needToInitializeFunc
		func(count, state uint64) {
			getter.GetContactsPageRecordsMatchStateCity(0, count, state, cfvlist.stateMatch, cfvlist.cityMatch)
		},
		// needToPrependFunc
		func(button js.Value, count, state uint64) {
			sortedIndex := notjs.GetAttributeUint64(button, sortedIndexAttributeName)
			getter.GetContactsPageRecordsMatchStateCity(sortedIndex-count, count, state, cfvlist.stateMatch, cfvlist.cityMatch)
		},
		// needToAppendFunc
		func(button js.Value, count, state uint64) {
			sortedIndex := notjs.GetAttributeUint64(button, sortedIndexAttributeName)
			getter.GetContactsPageRecordsMatchStateCity(sortedIndex+1, count, state, cfvlist.stateMatch, cfvlist.cityMatch)
		},
	)
	return cfvlist
}

// Start starts the carousel.
func (cfvlist *ContactFVList) Start() {
	cfvlist.FVList.Start()
}

// Build rebuilds a list in a panel in the carousel.
func (cfvlist *ContactFVList) Build(contacts []*types.ContactRecord, sortedIndex, recordCount, state uint64) {
	panelIndex := kickwasmwidgets.StateToSubPanelIndex(state)
	switch panelIndex {
	case 0:
		notjs := cfvlist.FVList.NotJS
		l := len(contacts)
		buttons := make([]js.Value, l, l)
		for i, contact := range contacts {
			button := notjs.CreateElementBUTTON()
			// sorted index
			notjs.SetAttributeUint64(button, sortedIndexAttributeName, sortedIndex)
			sortedIndex++
			// record id
			notjs.SetAttributeUint64(button, recordIDAttributeName, contact.ID)
			// state value
			notjs.SetAttribute(button, stateValueAttributeName, contact.State)
			// button text
			h4 := notjs.CreateElementH4()
			tn := notjs.CreateTextNode(contact.State)
			notjs.AppendChild(h4, tn)
			notjs.AppendChild(button, h4)
			// button onclick
			cb := notjs.RegisterEventCallBack(false, false, false, func(event js.Value) {
				// the user has clicked on a button
				target := notjs.GetEventTarget(event)
				for {
					if notjs.TagName(target) == "BUTTON" {
						break
					}
					target = notjs.ParentNode(target)
				}
				nextPanel := cfvlist.GetSubPanel(1)
				stateMatch := notjs.GetAttribute(target, stateValueAttributeName)
				cfvlist.stateMatch = stateMatch
				// set the next panel's list title.
				notjs.SetInnerText(nextPanel.Title, fmt.Sprintf("Cities in %s", stateMatch))
				cfvlist.getter.GetContactsPageCitiesMatchState(
					0,
					nextPanel.VList.GetPageSize(),
					nextPanel.VList.GetIDState(),
					stateMatch)
			})
			notjs.SetOnClick(button, cb)
			buttons[i] = button
		}
		cfvlist.FVList.Build(buttons, state, recordCount)
	case 1:
		notjs := cfvlist.FVList.NotJS
		l := len(contacts)
		buttons := make([]js.Value, l, l)
		for i, contact := range contacts {
			button := notjs.CreateElementBUTTON()
			// sorted index
			notjs.SetAttributeUint64(button, sortedIndexAttributeName, sortedIndex)
			sortedIndex++
			// record id
			notjs.SetAttributeUint64(button, recordIDAttributeName, contact.ID)
			// city value
			notjs.SetAttribute(button, cityValueAttributeName, contact.City)
			// button text
			h4 := notjs.CreateElementH4()
			tn := notjs.CreateTextNode(contact.City)
			notjs.AppendChild(h4, tn)
			notjs.AppendChild(button, h4)
			// button onclick
			cb := notjs.RegisterEventCallBack(false, false, false, func(event js.Value) {
				// the user has clicked on a button
				target := notjs.GetEventTarget(event)
				for {
					if notjs.TagName(target) == "BUTTON" {
						break
					}
					target = notjs.ParentNode(target)
				}
				nextPanel := cfvlist.GetSubPanel(2)
				cityMatch := notjs.GetAttribute(target, cityValueAttributeName)
				cfvlist.cityMatch = cityMatch
				// set the next panel's list title.
				notjs.SetInnerText(nextPanel.Title, fmt.Sprintf("Contacts in %s, %s", cityMatch, cfvlist.stateMatch))
				cfvlist.getter.GetContactsPageRecordsMatchStateCity(
					0,
					nextPanel.VList.GetPageSize(),
					nextPanel.VList.GetIDState(),
					cfvlist.stateMatch,
					cityMatch)
			})
			notjs.SetOnClick(button, cb)
			buttons[i] = button
		}
		cfvlist.FVList.Build(buttons, state, recordCount)
	default:
		notjs := cfvlist.FVList.NotJS
		l := len(contacts)
		buttons := make([]js.Value, l, l)
		for i, contact := range contacts {
			button := notjs.CreateElementBUTTON()
			// sorted index
			notjs.SetAttributeUint64(button, sortedIndexAttributeName, sortedIndex)
			sortedIndex++
			// record id
			notjs.SetAttributeUint64(button, recordIDAttributeName, contact.ID)
			// button text
			// name
			h4 := notjs.CreateElementH4()
			tn := notjs.CreateTextNode(contact.Name)
			notjs.AppendChild(h4, tn)
			notjs.AppendChild(button, h4)
			// address
			p := notjs.CreateElementP()
			tn = notjs.CreateTextNode(contact.Address1)
			notjs.AppendChild(p, tn)
			br := notjs.CreateElementBR()
			if len(contact.Address2) > 0 {
				notjs.AppendChild(p, br)
				tn = notjs.CreateTextNode(contact.Address2)
				notjs.AppendChild(p, tn)
				br = notjs.CreateElementBR()
				notjs.AppendChild(p, br)
			}
			tn = notjs.CreateTextNode(fmt.Sprintf("%s, %s %s", contact.City, contact.State, contact.Zip))
			notjs.AppendChild(p, tn)
			notjs.AppendChild(button, p)
			// contact
			p = notjs.CreateElementP()
			br = notjs.CreateElementBR()
			notjs.AppendChild(p, br)
			tn = notjs.CreateTextNode(contact.Email)
			notjs.AppendChild(p, tn)
			br = notjs.CreateElementBR()
			notjs.AppendChild(p, br)
			tn = notjs.CreateTextNode(contact.Phone)
			notjs.AppendChild(p, tn)
			br = notjs.CreateElementBR()
			notjs.AppendChild(p, br)
			tn = notjs.CreateTextNode(contact.Social)
			notjs.AppendChild(p, tn)
			notjs.AppendChild(button, p)
			// button onclick
			cb := notjs.RegisterEventCallBack(false, false, false, func(event js.Value) {
				// the user has clicked on a button
				target := notjs.GetEventTarget(event)
				for {
					if notjs.TagName(target) == "BUTTON" {
						break
					}
					target = notjs.ParentNode(target)
				}
				id := notjs.GetAttributeUint64(target, recordIDAttributeName)
				cfvlist.getter.GetContact(id)
			})
			notjs.SetOnClick(button, cb)
			buttons[i] = button
		}
		cfvlist.FVList.Build(buttons, state, recordCount)
	}
}
