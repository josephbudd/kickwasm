package viewtools

import (
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/spawntabs/renderer/notjs"
)

/*

WARNING:

DO NOT EDIT THIS FILE.

*/

// Visibility class names.
const (
	spawnIDReplacePattern  = "{{.SpawnID}}"
	TabClassName           = "tab"
	SelectedTabClassName   = "selected-tab"
	UnSelectedTabClassName = "unselected-tab"
	TabPanelClassName      = "panel-bound-to-tab"

	TabBarClassName      = "tab-bar"
	UnderTabBarClassName = "under-tab-bar"

	PanelClassName                   = "panel"
	PanelWithHeadingClassName        = "panel-with-heading"
	PanelWithTabBarClassName         = "panel-with-tab-bar"
	PanelHeadingClassName            = "heading-of-panel"
	PanelHeadingLevelPrefixClassName = "heading-of-panel-level-"
	TabPanelGroupClassName           = "inner-panel"
	UserContentClassName             = "user-content"
	ResizeMeWidthClassName           = "resize-me-width"

	SliderClassName                  = "slider"
	SliderPanelClassName             = "slider-panel"
	SliderPanelInnerClassName        = "slider-panel-pad"
	SliderPanelInnerSiblingClassName = "slider-panel-inner-sibling"
	SliderButtonPadClassName         = "slider-button-pad"

	SeenClassName       = "seen"
	UnSeenClassName     = "unseen"
	ToBeSeenClassName   = "tobe-seen"
	ToBeUnSeenClassName = "tobe-unseen"

	CookieCrumbClassName            = "cookie-crumb"
	CookieCrumbLevelPrefixClassName = "cookie-crumb-level-"

	VScrollClassName  = "vscroll"
	HVScrollClassName = "hvscroll"

	MasterID           = "tabsMasterView"
	HomeID             = "tabsMasterView-home"
	HomePadID          = "tabsMasterView-home-pad"
	SliderID           = "tabsMasterView-home-slider"
	SliderBackID       = "tabsMasterView-home-slider-back"
	SliderCollectionID = "tabsMasterView-home-slider-collection"

	BackIDAttribute         = "backid"
	BackColorLevelAttribute = "backColorLevel"
)

// Tools are application view tools.
type Tools struct {
	body                               js.Value
	tabsMasterview                     js.Value
	tabsMasterviewHome                 js.Value
	tabsMasterviewHomeButtonPad        js.Value
	tabsMasterviewHomeSlider           js.Value
	tabsMasterviewHomeSliderBack       js.Value
	tabsMasterviewHomeSliderCollection js.Value

	Document js.Value
	Global   js.Value

	// closer
	closerMasterView js.Value
	lastMasterView   js.Value
	// modal
	modalMasterView        js.Value
	modalMasterViewCenter  js.Value
	modalMasterViewH1      js.Value
	modalMasterViewMessage js.Value
	modalMasterViewClose   js.Value
	modalQueue             []*modalViewData
	modalQueueLastIndex    int
	beingModal             bool
	modalCallBack          func()
	// misc
	alert   js.Value
	console js.Value
	// groups
	buttonPanelsMap map[string][]js.Value
	// slider
	here      js.Value
	backStack []js.Value
	// tabber
	tabberLastPanelID     string
	tabberTabBarLastPanel map[string]string
	// button locking
	buttonsLocked             bool
	buttonsLockedMessageTitle string
	buttonsLockedMessageText  string

	// call backs
	jsCallBacks map[uint64][]js.Func

	NotJS *notjs.NotJS

	// spawns

	spawnID               uint64
	SpawnIDReplacePattern string

	// user content

	panelNameHVScroll map[string]bool
}

// NewTools constructs a new Tools
func NewTools(notJS *notjs.NotJS) *Tools {
	g := js.Global()
	v := &Tools{
		Document:              g.Get("document"),
		Global:                g,
		NotJS:                 notJS,
		SpawnIDReplacePattern: spawnIDReplacePattern,

		buttonPanelsMap: make(map[string][]js.Value, 100),
		here:            js.Undefined(),
		alert:           g.Get("alert"),
		console:         g.Get("console"),
		jsCallBacks:     make(map[uint64][]js.Func, 100),

		panelNameHVScroll: map[string]bool{"CreatePanel": false, "HelloWorldTemplatePanel": false, "TabsButtonTabBarPanel": false},
	}
	bodies := notJS.GetElementsByTagName("body")
	v.body = bodies[0]
	v.tabsMasterview = notJS.GetElementByID(MasterID)
	v.tabsMasterviewHome = notJS.GetElementByID(HomeID)
	v.tabsMasterviewHomeButtonPad = notJS.GetElementByID(HomePadID)
	v.tabsMasterviewHomeSlider = notJS.GetElementByID(SliderID)
	v.tabsMasterviewHomeSliderBack = notJS.GetElementByID(SliderBackID)
	v.tabsMasterviewHomeSliderCollection = notJS.GetElementByID(SliderCollectionID)
	// closer
	v.closerMasterView = notJS.GetElementByID("closerMasterView")
	// modal
	v.modalMasterView = notJS.GetElementByID("modalInformationMasterView")
	v.modalMasterViewCenter = notJS.GetElementByID("modalInformationMasterView-center")
	v.modalMasterViewH1 = notJS.GetElementByID("modalInformationMasterView-h1")
	v.modalMasterViewMessage = notJS.GetElementByID("modalInformationMasterView-message")
	v.modalMasterViewClose = notJS.GetElementByID("modalInformationMasterView-close")
	v.modalQueue = make([]*modalViewData, 5, 5)
	v.modalQueueLastIndex = -1
	cb := v.RegisterEventCallBack(v.handleModalMasterViewClose, true, true, true)
	notJS.SetOnClick(v.modalMasterViewClose, cb)

	// misc
	v.initializeGroups()
	v.initializeSlider()
	v.initializeResize()
	v.initializeCloser()
	v.initializeTabBar()

	return v
}
