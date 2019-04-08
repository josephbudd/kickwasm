package kickwasmwidgets

import (
	"math"
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

// VList is vertical list of verbose buttons.
type VList struct {
	div                  js.Value
	max                  uint64
	onSizeFunc           func()
	onNoSizeFunc         func()
	needToInitializeFunc func(count, state uint64)
	needToPrependFunc    func(button js.Value, count, state uint64)
	needToAppendFunc     func(button js.Value, count, state uint64)
	idState              uint64
	notJS                *notjs.NotJS
	tools                *viewtools.Tools

	hideFunc    func(js.Value)
	showFunc    func(js.Value)
	isShownFunc func(js.Value) bool

	adjusting     bool
	srcollTop     int
	lastScrollTop int
}

// NewVList constructs a new VList.
// Param div is the div containing the list.
//  If the div is not empty it will be emptied so don't bother putting anything in it.
//  div must have.
//   * a styled height ( See "github.com/josephbudd/kickwasmwidgets/css/vlist.css" )
//   * its children will only be p's  and buttons.
//   * child p's are the top and botton padding.
//   * child buttons are the verbose options.
// Param idState
//  is vlist vlist's unique id.
//  it must be generated by VListState.GetNextState().
//  it will be passed as param state being ored with StateInitialize in calls to needToInitializeFunc.
//  it will be passed as param state being ored with StatePrepend in calls to needToPrependFunc.
//  it will be passed as param state being ored with StateAppend in calls to needToAppendFunc.
// Param max is the maximum amount of options.
//   The list will never have more than max options to control memory usage.
// Param onSizeFunc
//  is called when there will be a size to the list.
// Param onNoSizeFunc
//  is called when there will be no size to the list.
// Param needToInitializeFunc
//  is called when the list needs initialized.
//  is passed the params(count, state)
//  is asynchronous and returns nothing
//  it is expected to propagate a call back to vlist.Build
// Param needToPrependFunc
//  is called when the list needs to prepend more records.
//  is passed the params (button, count, state)
//  is asynchronous and returns nothing
//  it is expected to propagate a call back to vlist.Build
// Param needToAppendFunc
//  is called when the list needs to append more records.
//  is passed the params (button, count, state)
//  is asynchronous and returns nothing
//  it is expected to propagate a call back to vlist.Build
// Param notJS is a pointer to notjs.NotJS
// Param hideFunc is a func that will hide this div.
//  it will take one param, any js.Value.
// Param showFunc is a func that will show this div.
//  it will take one param, any js.Value.
// Param isShownFunc is a func that will return if this div is shown.VList
//  it takes one param, any js.Value and returns a bool.
func NewVList(div js.Value,
	idState uint64,
	max uint64,
	onSizeFunc func(),
	onNoSizeFunc func(),
	needToInitializeFunc func(count, state uint64),
	needToPrependFunc func(button js.Value, count, state uint64),
	needToAppendFunc func(button js.Value, count, state uint64),
	hideFunc func(js.Value),
	showFunc func(js.Value),
	isShownFunc func(js.Value) bool,
	notJS *notjs.NotJS,
	tools *viewtools.Tools,
) *VList {
	vlist := &VList{
		div:                  div,
		max:                  max,
		onSizeFunc:           onSizeFunc,
		onNoSizeFunc:         onNoSizeFunc,
		needToInitializeFunc: needToInitializeFunc,
		needToPrependFunc:    needToPrependFunc,
		needToAppendFunc:     needToAppendFunc,
		idState:              idState,
		notJS:                notJS,
		hideFunc:             hideFunc,
		showFunc:             showFunc,
		isShownFunc:          isShownFunc,
	}
	// setup div
	notJS.RemoveChildNodes(div)
	// fix max
	if vlist.max%2 != 0 {
		vlist.max++
	}
	vlist.max /= 2
	// setup scrolling
	cb := tools.RegisterEventCallBack(vlist.handleOnScroll, true, true, true)
	notJS.SetOnScroll(div, cb)

	return vlist
}

// Start starts the list initializing it with the first records.
func (vlist *VList) Start() {
	state := StateInitialize | vlist.idState
	vlist.needToInitializeFunc(vlist.max, state)
}

// Hide hides the vlist.
func (vlist *VList) Hide() {
	vlist.hideFunc(vlist.div)
}

// Show unshides the vlist.
func (vlist *VList) Show() {
	vlist.showFunc(vlist.div)
}

// Toggle toggles the vlist visibility.
func (vlist *VList) Toggle() {
	if vlist.isShownFunc(vlist.div) {
		vlist.hideFunc(vlist.div)
	} else {
		vlist.showFunc(vlist.div)
	}
}

// Build rebuilds the vlist in some way.
func (vlist *VList) Build(buttons []js.Value, state, recordCount uint64) {
	vlist.adjusting = true
	if recordCount == 0 {
		// the user has not added any records.
		vlist.clear()
		vlist.onNoSizeFunc()
		return
	}
	if len(buttons) == 0 {
		// nothing to add
		return
	}
	// there are buttons.
	vlist.onSizeFunc()
	// adjust the lists.
	if (state & StateAppend) == StateAppend {
		vlist.append(buttons)
		return
	}
	if (state & StatePrepend) == StatePrepend {
		vlist.prepend(buttons)
		return
	}
	// StateInitialize
	vlist.initialize(buttons)
}

// GetPageSize returns the list's page size.
func (vlist *VList) GetPageSize() uint64 {
	return vlist.max
}

// GetIDState returns the list's idState.
func (vlist *VList) GetIDState() uint64 {
	return vlist.idState
}

func (vlist *VList) initialize(buttons []js.Value) {
	vlist.clear()
	vlist.append(buttons)
}

func (vlist *VList) append(buttons []js.Value) {
	notJS := vlist.notJS
	l := uint64(len(buttons))
	if l > vlist.max {
		buttons = buttons[:vlist.max]
	}
	div := vlist.div
	children := notJS.ChildrenSlice(div)
	l = uint64(len(children) - 1)
	bottom := children[l]
	count := l - 1
	maxlen := vlist.max * 2
	rc := 1
	// append to the back end of the list.
	for _, button := range buttons {
		notJS.InsertChildBefore(div, button, bottom)
		if count == maxlen {
			// remove the top of the list.
			notJS.RemoveChild(div, children[rc])
			rc++
		} else {
			// total # of buttons increased.
			count++
		}
		//div.scrollTo(0, div.scrollTop + Sizer.height(button))
	}
}

func (vlist *VList) prepend(buttons []js.Value) {
	notJS := vlist.notJS
	lr := uint64(len(buttons))
	if lr > vlist.max {
		buttons = buttons[:vlist.max]
		lr = vlist.max
	}
	div := vlist.div
	children := notJS.ChildrenSlice(div)
	rc := uint64(len(children)) - 2
	count := rc
	maxlen := vlist.max * 2
	scrollTop := float64(notJS.GetScrollTop(div))
	// prepend to the front end of the list.
	top := children[1]
	for _, button := range buttons {
		notJS.InsertChildBefore(div, button, top)
		// remove the bottom of the list.
		if count == maxlen {
			notJS.RemoveChild(div, children[rc])
			rc--
		} else {
			// total # of buttons increased.
			count++
		}
	}
	buttonHt := float64(0)
	children = notJS.ChildrenSlice(div)
	if len(children) > 2 {
		buttonHt = notJS.OuterHeight(children[1])
	}
	scrollTo := scrollTop + (buttonHt * float64(lr))
	to := scrollTop + scrollTo
	vlist.notJS.ScrollTo(div, 0, int(math.Floor(to)))
}

func (vlist *VList) clear() {
	notJS := vlist.notJS
	div := vlist.div
	notJS.RemoveChildNodes(div)
	p := notJS.CreateElement("p")
	notJS.AppendChild(div, p)
	p = notJS.CreateElement("p")
	notJS.AppendChild(div, p)
}

func (vlist *VList) handleOnScroll(event js.Value) interface{} {
	if vlist.adjusting {
		vlist.adjusting = false
		return nil
	}
	notJS := vlist.notJS
	div := notJS.GetEventTarget(event)
	lastScrollTop := float64(vlist.lastScrollTop)
	vlist.lastScrollTop = notJS.GetScrollTop(div)
	scrollTop := float64(vlist.lastScrollTop)
	children := notJS.ChildrenSlice(div)
	l := len(children)
	bottom := float64(0)
	paddingHt := float64(0)
	if l >= 2 {
		// the first and last children are the p which are the padding.
		paddingHt = notJS.OuterHeight(children[0])
		bottom += paddingHt * 2.0
	}
	if l > 2 {
		// there is more then just padding
		bottom += notJS.OuterHeight(children[1]) * float64(l-2)
	}
	bottom -= notJS.InnerHeight(div)
	if scrollTop < lastScrollTop {
		// scrolling up
		if scrollTop > paddingHt {
			// not at the top.
			return nil
		}
		// at the top
		if lastScrollTop > paddingHt {
			// did not just do vlist the time before.
			vlist.needToPrependFunc(
				children[1],
				vlist.max,
				StatePrepend|vlist.idState,
			)
		}
	} else {
		// scrolling down
		if scrollTop < bottom {
			// not at the bottom
			return nil
		}
		// at the bottom
		// xcrollTop == bottom
		if lastScrollTop != bottom {
			// first time at the bottom
			vlist.needToAppendFunc(
				children[len(children)-2],
				vlist.max,
				StateAppend|vlist.idState,
			)
		}
	}
	return nil
}
