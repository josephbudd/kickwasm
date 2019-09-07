package templates

// ViewTools is the renderer/viewtools/viewtools.go template.
const ViewTools = `package viewtools

import (
	"syscall/js"

	"{{.ApplicationGitPath}}{{.ImportRendererNotJS}}"
)

/*

WARNING:

DO NOT EDIT THIS FILE.

*/

// Visibility class names.
const (
	spawnIDReplacePattern  = "{{.SpawnIDReplacePattern}}"
	TabClassName           = "{{.Classes.Tab}}"
	SelectedTabClassName   = "{{.Classes.SelectedTab}}"
	UnSelectedTabClassName = "{{.Classes.UnSelectedTab}}"
	TabPanelClassName      = "{{.Classes.TabPanel}}"

	TabBarClassName      = "{{.Classes.TabBar}}"
	UnderTabBarClassName = "{{.Classes.UnderTabBar}}"

	PanelClassName                   = "{{.Classes.Panel}}"
	PanelWithHeadingClassName        = "{{.Classes.PanelWithHeading}}"
	PanelWithTabBarClassName         = "{{.Classes.PanelWithTabBar}}"
	PanelHeadingClassName            = "{{.Classes.PanelHeading}}"
	PanelHeadingLevelPrefixClassName = "{{.Classes.PanelHeadingLevelPrefix}}"
	TabPanelGroupClassName           = "{{.Classes.TabPanelGroup}}"
	UserContentClassName             = "{{.Classes.UserContent}}"
	ResizeMeWidthClassName           = "{{.Classes.ResizeMeWidth}}"

	SliderClassName                  = "{{.Classes.Slider}}"
	SliderPanelClassName             = "{{.Classes.SliderPanel}}"
	SliderPanelInnerClassName        = "{{.Classes.SliderPanelPad}}"
	SliderPanelInnerSiblingClassName = "{{.Classes.SliderPanelInnerSibling}}"
	SliderButtonPadClassName         = "{{.Classes.SliderButtonPad}}"

	SeenClassName       = "{{.Classes.Seen}}"
	UnSeenClassName     = "{{.Classes.UnSeen}}"
	ToBeSeenClassName   = "{{.Classes.ToBeSeen}}"
	ToBeUnSeenClassName = "{{.Classes.ToBeUnSeen}}"

	CookieCrumbClassName            = "{{.Classes.CookieCrumb}}"
	CookieCrumbLevelPrefixClassName = "{{.Classes.CookieCrumbLevelPrefix}}"

	VScrollClassName  = "{{.Classes.VScroll}}"
	HVScrollClassName = "{{.Classes.HVScroll}}"

	MasterID           = "{{.IDs.Master}}"
	HomeID             = "{{.IDs.Home}}"
	HomePadID          = "{{.IDs.HomePad}}"
	SliderID           = "{{.IDs.Slider}}"
	SliderBackID       = "{{.IDs.SliderBack}}"
	SliderCollectionID = "{{.IDs.SliderCollection}}"

	BackIDAttribute         = "{{.Attributes.BackID}}"
	BackColorLevelAttribute = "{{.Attributes.BackColorLevel}}"
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

	// markup panels

	countMarkupPanels        int
	countSpawnedMarkupPanels int
	countWidgetsWaiting      int

	// spawned widgets

	spawnedWidgets map[uint64]spawnedWidgetInfo
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

		panelNameHVScroll: {{.PanelNameHVScroll}},

		countMarkupPanels: {{.NumberOfMarkupPanels}},

		spawnedWidgets: make(map[uint64]spawnedWidgetInfo, 100),
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
`
