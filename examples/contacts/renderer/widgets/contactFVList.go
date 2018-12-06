package widgets

import (
	"fmt"
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/kickwasmwidgets"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
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
	notJS *notjs.NotJS,
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
		notJS,
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
			sortedIndex := notJS.GetAttributeUint64(button, sortedIndexAttributeName)
			if count > sortedIndex {
				return
			}
			getter.GetContactsPageStates(sortedIndex-count, count, state)
		},
		// needToAppendFunc
		func(button js.Value, count, state uint64) {
			sortedIndex := notJS.GetAttributeUint64(button, sortedIndexAttributeName)
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
			sortedIndex := notJS.GetAttributeUint64(button, sortedIndexAttributeName)
			getter.GetContactsPageCitiesMatchState(sortedIndex-count, count, state, cfvlist.stateMatch)
		},
		// needToAppendFunc
		func(button js.Value, count, state uint64) {
			sortedIndex := notJS.GetAttributeUint64(button, sortedIndexAttributeName)
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
			sortedIndex := notJS.GetAttributeUint64(button, sortedIndexAttributeName)
			getter.GetContactsPageRecordsMatchStateCity(sortedIndex-count, count, state, cfvlist.stateMatch, cfvlist.cityMatch)
		},
		// needToAppendFunc
		func(button js.Value, count, state uint64) {
			sortedIndex := notJS.GetAttributeUint64(button, sortedIndexAttributeName)
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
		notJS := cfvlist.FVList.NotJS
		l := len(contacts)
		buttons := make([]js.Value, l, l)
		for i, contact := range contacts {
			button := notJS.CreateElementBUTTON()
			// sorted index
			notJS.SetAttributeUint64(button, sortedIndexAttributeName, sortedIndex)
			sortedIndex++
			// record id
			notJS.SetAttributeUint64(button, recordIDAttributeName, contact.ID)
			// state value
			notJS.SetAttribute(button, stateValueAttributeName, contact.State)
			// button text
			h4 := notJS.CreateElementH4()
			tn := notJS.CreateTextNode(contact.State)
			notJS.AppendChild(h4, tn)
			notJS.AppendChild(button, h4)
			// button onclick
			cb := notJS.RegisterEventCallBack(false, false, false, func(event js.Value) {
				// the user has clicked on a button
				target := notJS.GetEventTarget(event)
				for {
					if notJS.TagName(target) == "BUTTON" {
						break
					}
					target = notJS.ParentNode(target)
				}
				nextPanel := cfvlist.GetSubPanel(1)
				stateMatch := notJS.GetAttribute(target, stateValueAttributeName)
				cfvlist.stateMatch = stateMatch
				// set the next panel's list title.
				notJS.SetInnerText(nextPanel.Title, fmt.Sprintf("Cities in %s", stateMatch))
				cfvlist.getter.GetContactsPageCitiesMatchState(
					0,
					nextPanel.VList.GetPageSize(),
					nextPanel.VList.GetIDState(),
					stateMatch)
			})
			notJS.SetOnClick(button, cb)
			buttons[i] = button
		}
		cfvlist.FVList.Build(buttons, state, recordCount)
	case 1:
		notJS := cfvlist.FVList.NotJS
		l := len(contacts)
		buttons := make([]js.Value, l, l)
		for i, contact := range contacts {
			button := notJS.CreateElementBUTTON()
			// sorted index
			notJS.SetAttributeUint64(button, sortedIndexAttributeName, sortedIndex)
			sortedIndex++
			// record id
			notJS.SetAttributeUint64(button, recordIDAttributeName, contact.ID)
			// city value
			notJS.SetAttribute(button, cityValueAttributeName, contact.City)
			// button text
			h4 := notJS.CreateElementH4()
			tn := notJS.CreateTextNode(contact.City)
			notJS.AppendChild(h4, tn)
			notJS.AppendChild(button, h4)
			// button onclick
			cb := notJS.RegisterEventCallBack(false, false, false, func(event js.Value) {
				// the user has clicked on a button
				target := notJS.GetEventTarget(event)
				for {
					if notJS.TagName(target) == "BUTTON" {
						break
					}
					target = notJS.ParentNode(target)
				}
				nextPanel := cfvlist.GetSubPanel(2)
				cityMatch := notJS.GetAttribute(target, cityValueAttributeName)
				cfvlist.cityMatch = cityMatch
				// set the next panel's list title.
				notJS.SetInnerText(nextPanel.Title, fmt.Sprintf("Contacts in %s, %s", cityMatch, cfvlist.stateMatch))
				cfvlist.getter.GetContactsPageRecordsMatchStateCity(
					0,
					nextPanel.VList.GetPageSize(),
					nextPanel.VList.GetIDState(),
					cfvlist.stateMatch,
					cityMatch)
			})
			notJS.SetOnClick(button, cb)
			buttons[i] = button
		}
		cfvlist.FVList.Build(buttons, state, recordCount)
	default:
		notJS := cfvlist.FVList.NotJS
		l := len(contacts)
		buttons := make([]js.Value, l, l)
		for i, contact := range contacts {
			button := notJS.CreateElementBUTTON()
			// sorted index
			notJS.SetAttributeUint64(button, sortedIndexAttributeName, sortedIndex)
			sortedIndex++
			// record id
			notJS.SetAttributeUint64(button, recordIDAttributeName, contact.ID)
			// button text
			// name
			h4 := notJS.CreateElementH4()
			tn := notJS.CreateTextNode(contact.Name)
			notJS.AppendChild(h4, tn)
			notJS.AppendChild(button, h4)
			// address 1 & 2
			p := notJS.CreateElementP()
			tn = notJS.CreateTextNode(contact.Address1)
			notJS.AppendChild(p, tn)
			br := notJS.CreateElementBR()
			if len(contact.Address2) > 0 {
				notJS.AppendChild(p, br)
				tn = notJS.CreateTextNode(contact.Address2)
				notJS.AppendChild(p, tn)
			}
			// city, state zip
			br = notJS.CreateElementBR()
			notJS.AppendChild(p, br)
			tn = notJS.CreateTextNode(fmt.Sprintf("%s, %s %s", contact.City, contact.State, contact.Zip))
			notJS.AppendChild(p, tn)
			notJS.AppendChild(button, p)
			// email, phone, social
			p = notJS.CreateElementP()
			br = notJS.CreateElementBR()
			notJS.AppendChild(p, br)
			tn = notJS.CreateTextNode(contact.Email)
			notJS.AppendChild(p, tn)
			br = notJS.CreateElementBR()
			notJS.AppendChild(p, br)
			tn = notJS.CreateTextNode(contact.Phone)
			notJS.AppendChild(p, tn)
			br = notJS.CreateElementBR()
			notJS.AppendChild(p, br)
			tn = notJS.CreateTextNode(contact.Social)
			notJS.AppendChild(p, tn)
			notJS.AppendChild(button, p)
			// button onclick
			cb := notJS.RegisterEventCallBack(false, false, false, func(event js.Value) {
				// the user has clicked on a button
				target := notJS.GetEventTarget(event)
				for {
					if notJS.TagName(target) == "BUTTON" {
						break
					}
					target = notJS.ParentNode(target)
				}
				id := notJS.GetAttributeUint64(target, recordIDAttributeName)
				cfvlist.getter.GetContact(id)
			})
			notJS.SetOnClick(button, cb)
			buttons[i] = button
		}
		cfvlist.FVList.Build(buttons, state, recordCount)
	}
}
