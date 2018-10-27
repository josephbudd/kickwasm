package tap

import (
	"strings"

	"github.com/josephbudd/kickwasm/cases"
)

// exported constants
const (
	CommentLine  = "//"
	CommentStart = "/*"
	CommentEnd   = "*/"
	EmptyString  = ""
	MaxLevels    = 5

	brString = "<br/>"

	suffixPanel  = "Panel"
	suffixButton = "Button"
	suffixTab    = "Tab"

	newline             = "\n"
	spaceString         = " "
	emptyString         = ""
	dashString          = "-"
	dashSubPanelString  = "-subpanel"
	dashButtonPadString = "-button-pad"
	dashTabBar          = "-tab-bar"
	underscoreTabBar    = "_tab_bar"
	dashUnderTabBar     = "-under-tab-bar"
	dashInnerString     = "-inner"
	dashContentString   = "-user-content"
	todo                = "<!-- TODO : Replace kick's instructional markup with your own. -->"
	indentationString   = "                                                                                                    "
	beginning           = "<!-- Beginning of kick's instructional markup that you want to replace with your own. -->"
	end                 = "<!-- End of kick's instructional markup that you want to replace with your own. -->"
	add                 = " You want to add your markup to this div."

	underscoreRune   = '_'
	underscoreString = "_"

	indentAmount = 2

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
	classUnSelected              = "unselected-tab"
	classPanelWithHeading        = "panel-with-heading"
	classPanelWithTabBar         = "panel-with-tab-bar"
	classInnerPanel              = "inner-panel"
	classPanelHeading            = "heading-of-panel"
	classPanelHeadingLevelPrefix = "heading-of-panel-level-"

	classSlider     = "slider"
	classSliderBack = "slider-back"

	classSliderPanel             = "slider-panel"
	classSliderPanelInner        = "slider-panel-inner"
	classSliderPanelInnerSibling = "slider-panel-inner-sibling"
	classSliderButtonPad         = "slider-button-pad"
	classSliderTabBar            = "slider-tab-bar"

	classSliderPanelWithButtonPad = "slider-panel-with-button-pad"
	classSliderPanelWithTabBar    = "slider-panel-with-tab-bar"

	classBackColorLevelPrefix      = "back-color-level-"
	classPadColorLevelPrefix       = "pad-color-level-"
	classPadButtonColorLevelPrefix = "pad-button-color-level-"

	classCookieCrumb            = "cookie-crumb"
	classCookieCrumbLevelPrefix = "cookie-crumb-level-"

	classUserContent            = "user-content"
	classResizeMeWidthClassName = "resize-me-width"

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
