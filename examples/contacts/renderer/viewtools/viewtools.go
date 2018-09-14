package viewtools

import (
	"syscall/js"

	"github.com/josephbudd/kicknotjs"
)

/*

WARNING:

DO NOT EDIT THIS FILE.

*/

// Visibility class names.
const (
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
	InnerPanelClassName              = "inner-panel"
	UserContentClassName             = "user-content"

	SliderClassName                  = "slider"
	SliderPanelClassName             = "slider-panel"
	SliderPanelInnerClassName        = "slider-panel-inner"
	SliderButtonPadClassName         = "slider-button-pad"
	SliderPanelInnerSiblingClassName = "slider-panel-inner-sibling"

	SeenClassName       = "seen"
	UnSeenClassName     = "unseen"
	ToBeSeenClassName   = "tobe-seen"
	ToBeUnSeenClassName = "tobe-unseen"

	CookieCrumbClassName            = "cookie-crumb"
	CookieCrumbLevelPrefixClassName = "cookie-crumb-level-"

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
	tabberLastPanelLevels map[string]string

	notjs *kicknotjs.NotJS
}

// NewTools constructs a new Tools
func NewTools(notjs *kicknotjs.NotJS) *Tools {
	g := js.Global()
	v := &Tools{
		Document:        g.Get("document"),
		Global:          g,
		buttonPanelsMap: make(map[string][]js.Value),
		here:            js.Undefined(),
		alert:           g.Get("alert"),
		console:         g.Get("console"),
	}
	v.notjs = notjs
	bodies := v.GetElementsByTagName("body")
	v.body = bodies[0]
	v.tabsMasterview = v.GetElementByID(MasterID)
	v.tabsMasterviewHome = v.GetElementByID(HomeID)
	v.tabsMasterviewHomeButtonPad = v.GetElementByID(HomePadID)
	v.tabsMasterviewHomeSlider = v.GetElementByID(SliderID)
	v.tabsMasterviewHomeSliderBack = v.GetElementByID(SliderBackID)
	v.tabsMasterviewHomeSliderCollection = v.GetElementByID(SliderCollectionID)
	// closer
	v.closerMasterView = v.GetElementByID("closerMasterView")
	// modal
	v.modalMasterView = v.GetElementByID("modalInformationMasterView")
	v.modalMasterViewCenter = v.GetElementByID("modalInformationMasterView-center")
	v.modalMasterViewH1 = v.GetElementByID("modalInformationMasterView-h1")
	v.modalMasterViewMessage = v.GetElementByID("modalInformationMasterView-message")
	v.modalMasterViewClose = v.GetElementByID("modalInformationMasterView-close")
	v.modalQueue = make([]*modalViewData, 5, 5)
	v.modalQueueLastIndex = -1
	cb := notjs.RegisterCallBack(v.handleModalMasterViewClose)
	notjs.SetOnClick(v.modalMasterViewClose, cb)

	// misc
	v.initializeGroups()
	v.initializeSlider()
	v.initializeResize()
	v.initializeCloser()
	v.initializeTabBar()

	return v
}
