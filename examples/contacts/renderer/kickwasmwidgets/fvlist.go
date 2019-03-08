package kickwasmwidgets

import (
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/contacts/renderer/notjs"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

const (
	fvlistClassName                   = "fvlist"
	arrowClassName                    = "arrow"
	listClassName                     = "vlist"
	titleClassName                    = "title"
	filterWrapper1ClassName           = "filter-wrapper-1"
	filterWrapperClassName            = "filter-wrapper"
	recordWrapperClassName            = "record-wrapper"
	listWrapperClassName              = "list-wrapper"
	arrowTargetSubpanelIndexAttribute = "target-subpanel-index"
)

// FVListPanel is information about a carousel panel.
type FVListPanel struct {
	Panel       js.Value
	ListWrapper js.Value
	Arrow       js.Value
	Title       js.Value
	VList       *VList
}

// FVList is a carousel of VLists.
// The lists are a series of filters followed by a final filtered list of articles.
type FVList struct {
	div     js.Value
	IDState uint64
	panels  []*FVListPanel

	hideFunc     func(js.Value)
	showFunc     func(js.Value)
	isShownFunc  func(js.Value) bool
	onSizeFunc   func()
	onNoSizeFunc func()

	NotJS            *notjs.NotJS
	Tools            *viewtools.Tools
	addedRecordVList bool

	StateMatch string
	CityMatch  string

	openPanelIndex int
}

// NewFVList constructs a new FVList.
// Param div is the div containing the list.
//  If the div is not empty it will be emptied so don't bother putting anything in it.
//  div must have.
//   * a styled height ( See "github.com/josephbudd/kickwasmwidgets/css/vlist.css" )
// Param IDState
//  is vlist vlist's unique id.
//  it must be generated by VListState.GetNextState().
//  it will be passed as param state being ored with StateInitialize in calls to needToInitializeFunc.
//  it will be passed as param state being ored with StatePrepend in calls to needToPrependFunc.
//  it will be passed as param state being ored with StateAppend in calls to needToAppendFunc.
// Param onSizeFunc
//  is called when there will be a size to the list.
// Param onNoSizeFunc
//  is called when there will be no size to the list.
// Param hideFunc is a func that will hide this div.
//  it will take one param, any js.Value.
// Param showFunc is a func that will show this div.
//  it will take one param, any js.Value.
// Param isShownFunc is a func that will return if this div is shown.VList
//  it takes one param, any js.Value and returns a bool.
// Param notJS is a pointer to notjs.NotJS
func NewFVList(div js.Value,
	IDState uint64,
	onSizeFunc func(),
	onNoSizeFunc func(),
	hideFunc func(js.Value),
	showFunc func(js.Value),
	isShownFunc func(js.Value) bool,
	notJS *notjs.NotJS,
	tools *viewtools.Tools,
) *FVList {
	// setup div
	notJS.RemoveChildNodes(div)
	return &FVList{
		div:            div,
		IDState:        IDState,
		panels:         make([]*FVListPanel, 0, 5),
		NotJS:          notJS,
		Tools:          tools,
		hideFunc:       hideFunc,
		showFunc:       showFunc,
		isShownFunc:    isShownFunc,
		onSizeFunc:     onSizeFunc,
		onNoSizeFunc:   onNoSizeFunc,
		openPanelIndex: -1,
	}
}

// BuildFVList builds a new FVList.
// Param div is the div containing the list.
//  If the div is not empty it will be emptied so don't bother putting anything in it.
//  div must have.
//   * a styled height ( See "github.com/josephbudd/kickwasmwidgets/css/vlist.css" )
// Param IDState
//  is vlist vlist's unique id.
//  it must be generated by VListState.GetNextState().
//  it will be passed as param state being ored with StateInitialize in calls to needToInitializeFunc.
//  it will be passed as param state being ored with StatePrepend in calls to needToPrependFunc.
//  it will be passed as param state being ored with StateAppend in calls to needToAppendFunc.
// Param onSizeFunc
//  is called when there will be a size to the list.
// Param onNoSizeFunc
//  is called when there will be no size to the list.
// Param hideFunc is a func that will hide this div.
//  it will take one param, any js.Value.
// Param showFunc is a func that will show this div.
//  it will take one param, any js.Value.
// Param isShownFunc is a func that will return if this div is shown.VList
//  it takes one param, any js.Value and returns a bool.
// Param notJS is a pointer to notjs.NotJS
func BuildFVList(div js.Value,
	IDState uint64,
	onSizeFunc func(),
	onNoSizeFunc func(),
	hideFunc func(js.Value),
	showFunc func(js.Value),
	isShownFunc func(js.Value) bool,
	notJS *notjs.NotJS,
	tools *viewtools.Tools,
) FVList {
	// setup div
	notJS.RemoveChildNodes(div)
	return FVList{
		div:            div,
		IDState:        IDState,
		panels:         make([]*FVListPanel, 0, 5),
		NotJS:          notJS,
		Tools:          tools,
		hideFunc:       hideFunc,
		showFunc:       showFunc,
		isShownFunc:    isShownFunc,
		onSizeFunc:     onSizeFunc,
		onNoSizeFunc:   onNoSizeFunc,
		openPanelIndex: -1,
	}
}

// AddFirstFilter adds the very first filter list.
func (fvlist *FVList) AddFirstFilter(
	title string,
	max uint64,
	needToInitializeFunc func(count, state uint64),
	needToPrependFunc func(button js.Value, count, state uint64),
	needToAppendFunc func(button js.Value, count, state uint64),
) {
	if len(fvlist.panels) > 0 {
		panic("you must call FVList.AddFilter after calling FVList.AddFirstFilter")
	}
	fvlist.addList(max, needToInitializeFunc, needToPrependFunc, needToAppendFunc, filterWrapper1ClassName)
	panel := fvlist.panels[0]
	fvlist.NotJS.SetInnerText(panel.Title, title)
}

// AddAnotherFilter adds an additional filter list.
func (fvlist *FVList) AddAnotherFilter(
	max uint64,
	needToInitializeFunc func(count, state uint64),
	needToPrependFunc func(button js.Value, count, state uint64),
	needToAppendFunc func(button js.Value, count, state uint64),
) {
	if len(fvlist.panels) == 0 {
		panic("you must call FVList.AddFirstFilter before calling FVList.AddFilter")
	}
	if fvlist.addedRecordVList {
		panic("you must not call FVList.AddFilter after calling FVList.AddRecordList")
	}
	fvlist.addList(max, needToInitializeFunc, needToPrependFunc, needToAppendFunc, filterWrapperClassName)
}

// AddRecordList adds the final list, the record list.
func (fvlist *FVList) AddRecordList(
	max uint64,
	needToInitializeFunc func(count, state uint64),
	needToPrependFunc func(button js.Value, count, state uint64),
	needToAppendFunc func(button js.Value, count, state uint64),
) {
	if len(fvlist.panels) == 0 {
		panic("if you don't have any filters use VList not FVlist")
	}
	if fvlist.addedRecordVList {
		panic("you must not call FVList.AddRecordList after calling FVList.AddRecordList")
	}
	fvlist.addList(max, needToInitializeFunc, needToPrependFunc, needToAppendFunc, recordWrapperClassName)
	fvlist.OpenSubPanel(0)
}

// Start starts the list initializing it with the first records.
func (fvlist *FVList) Start() {
	panel := fvlist.panels[0]
	panel.VList.Start()
}

// Build rebuilds the fvlist in some way.
func (fvlist *FVList) Build(buttons []js.Value, state, recordCount uint64) {
	panelIndex := StateToSubPanelIndex(state)
	panel := fvlist.panels[panelIndex]
	panel.VList.Build(buttons, state, recordCount)
	fvlist.OpenSubPanel(int(panelIndex))
}

// Hide hides the fvlist.
func (fvlist *FVList) Hide() {
	fvlist.hideFunc(fvlist.div)
}

// Show unshides the fvlist.
func (fvlist *FVList) Show() {
	fvlist.showFunc(fvlist.div)
}

// Toggle toggles the fvlist visibility.
func (fvlist *FVList) Toggle() {
	if fvlist.isShownFunc(fvlist.div) {
		fvlist.hideFunc(fvlist.div)
	} else {
		fvlist.showFunc(fvlist.div)
	}
}

// OpenSubPanel opens a sub panel.
func (fvlist *FVList) OpenSubPanel(subPanelIndex int) {
	if fvlist.openPanelIndex == subPanelIndex {
		return
	}
	if subPanelIndex < 0 || subPanelIndex >= len(fvlist.panels) {
		return
	}
	fvlist.openPanelIndex = subPanelIndex
	for i, panel := range fvlist.panels {
		if i != subPanelIndex {
			//fvlist.hideFunc(panel.VList.div)
			fvlist.hideFunc(panel.Panel)
		} else {
			//fvlist.showFunc(panel.VList.div)
			fvlist.showFunc(panel.Panel)
		}
	}
}

// SetSubPanelTitle sets the text of a sub panel's list title.
func (fvlist *FVList) SetSubPanelTitle(subPanelIndex int, text string) {
	panel := fvlist.panels[subPanelIndex]
	fvlist.NotJS.SetInnerText(panel.Title, text)
}

// GetSubPanel returns one of the fvlist panel structs.
func (fvlist *FVList) GetSubPanel(index int) *FVListPanel {
	if index < 0 {
		return nil
	}
	if index >= len(fvlist.panels) {
		return nil
	}
	return fvlist.panels[index]
}

func (fvlist *FVList) addList(
	max uint64,
	needToInitializeFunc func(count, state uint64),
	needToPrependFunc func(button js.Value, count, state uint64),
	needToAppendFunc func(button js.Value, count, state uint64),
	wrapperClassName string,
) {
	// filter wrapper
	i := len(fvlist.panels)
	notJS := fvlist.NotJS
	tools := fvlist.Tools
	wrapper := notJS.CreateElementDIV()
	//notJS.SetID(wrapper, fmt.Sprintf(subDivIndexAttributeNameFormatter, i))
	// left arrow
	var arrow js.Value
	notJS.ClassListAddClass(wrapper, wrapperClassName)
	if i > 0 {
		arrow = notJS.CreateElementBUTTON()
		notJS.ClassListAddClass(arrow, arrowClassName)
		tn := notJS.CreateTextNode("↩")
		notJS.AppendChild(arrow, tn)
		notJS.SetAttributeInt(arrow, arrowTargetSubpanelIndexAttribute, i-1)
		cb := tools.RegisterEventCallBack(
			func(event js.Value) interface{} {
				index := notJS.GetAttributeInt(arrow, arrowTargetSubpanelIndexAttribute)
				fvlist.OpenSubPanel(index)
				return nil
			},
			false, false, false,
		)
		notJS.SetOnClick(arrow, cb)
		notJS.AppendChild(wrapper, arrow)
	}
	// add the list wrapper
	listwrapper := notJS.CreateElementDIV()
	notJS.ClassListAddClass(listwrapper, listWrapperClassName)
	notJS.AppendChild(wrapper, listwrapper)
	// title
	title := notJS.CreateElementH4() // span
	notJS.ClassListAddClass(title, titleClassName)
	notJS.AppendChild(listwrapper, title)
	// list
	list := notJS.CreateElementDIV()
	notJS.ClassListAddClass(list, listClassName)
	notJS.AppendChild(listwrapper, list)
	// add the whole filter wrapper
	notJS.AppendChild(fvlist.div, wrapper)
	// vlist
	vlist := NewVList(list,
		fvlist.IDState|uint64(i),
		max,
		fvlist.onSizeFunc,
		fvlist.onNoSizeFunc,
		needToInitializeFunc,
		needToPrependFunc,
		needToAppendFunc,
		fvlist.hideFunc,
		fvlist.showFunc,
		fvlist.isShownFunc,
		notJS,
		tools,
	)
	panel := &FVListPanel{
		Panel:       wrapper,
		ListWrapper: listwrapper,
		Arrow:       arrow,
		Title:       title,
		VList:       vlist,
	}
	fvlist.panels = append(fvlist.panels, panel)
}
