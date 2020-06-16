package project

import (
	"strings"

	"github.com/josephbudd/kickwasm/pkg/cases"
)

// exported constants
const (
	SpawnIDReplacePattern = "{{.SpawnID}}"
	CommentLine           = "//"
	CommentStart          = "/*"
	CommentEnd            = "*/"
	EmptyString           = ""
	MaxLevels             = 5
	DashUnderTabBar       = "-under-tab-bar"
	DashInnerString       = "-inner"
	DashPanelHeading      = "-panel-heading"
	DashNoTabButtons      = "-no-tab-buttons"

	brString = "<br/>"

	suffixPanel  = "Panel"
	suffixButton = "Button"
	suffixTab    = "Tab"

	newline             = "\n"
	spaceString         = " "
	emptyString         = ""
	dashString          = "-"
	dashH3String        = "-H3"
	dashSubPanelString  = "-subpanel"
	dashButtonPadString = "-button-pad"
	dashTabBar          = "-tab-bar"
	underscoreTabBar    = "_tab_bar"
	dashContentString   = "-user-content"
	todo                = "<!-- TODO : Replace kickwasm's instructional markup with your own. -->"
	beginning           = "<!-- Beginning of kickwasm's instructional markup that you want to replace with your own. -->"
	end                 = "<!-- End of kickwasm's instructional markup that you want to replace with your own. -->"
	add                 = " You want to add your markup to this div."

	underscoreRune   = '_'
	underscoreString = "_"

	classTabBar      = "tab-bar"
	classTab         = "tab"
	classTabPanel    = "panel-bound-to-tab"
	classUnderTabBar = "under-tab-bar"
	classPanel       = "panel"
	classSeen        = "seen"
	classUnSeen      = "unseen"
	classToBeSeen    = "tobe-seen"
	classToBeUnSeen  = "tobe-unseen"

	classSelected                = "selected-tab"
	classUnderTab                = "under-tab"
	classUnSelected              = "unselected-tab"
	classPanelWithHeading        = "panel-with-heading"
	classPanelWithTabBar         = "panel-with-tab-bar"
	classTabPanelGroup           = "inner-panel"
	classPanelHeading            = "heading-of-panel"
	classPanelHeadingLevelPrefix = "heading-of-panel-level-"

	classSlider     = "slider"
	classSliderBack = "slider-back"

	classSliderPanel     = "slider-panel"
	classSliderPanelPad  = "slider-panel-pad"
	classSliderButtonPad = "slider-button-pad"
	classSliderTabBar    = "slider-tab-bar"

	classSliderPanelWithButtonPad = "slider-panel-with-button-pad"
	classSliderPanelWithTabBar    = "slider-panel-with-tab-bar"

	classSliderPanelInnerSibling = "slider-panel-inner-sibling"

	classBackColorLevelPrefix      = "back-color-level-"
	classPadColorLevelPrefix       = "pad-color-level-"
	classPadButtonColorLevelPrefix = "pad-button-color-level-"

	classCookieCrumb            = "cookie-crumb"
	classCookieCrumbLevelPrefix = "cookie-crumb-level-"

	classUserContent             = "user-content"
	classModalUserContent        = "modal-user-content"
	classResizeMeWidthClassName  = "resize-me-width"
	classResizeMeHeightClassName = "resize-me-height"
	classDoNotPrintClassName     = "do-not-print"

	classVScroll    = "vscroll"
	classHVScroll   = "hvscroll"
	classUserMarkup = "user-markup"

	attributeSpawnable = "spawnable"

	attributeBackID         = "backid"
	attributeBackColorLevel = "backColorLevel"
)

var (
	emptyBB = []byte("")
)

func panelIDToName(id string) string {
	return cases.CamelCase(strings.Replace(strings.Replace(id, underscoreString, spaceString, -1), dashString, spaceString, -1))
}

func splitBackTicked(src string) []string {
	return strings.Split(src, "\n")
}
