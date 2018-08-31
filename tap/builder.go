package tap

import (
	"fmt"
	"strings"
)

// Colors are information about css colors.
type Colors struct {
	FirstColorLevel                uint
	LastColorLevel                 uint
	ClassBackColorLevelPrefix      string
	ClassPadColorLevelPrefix       string
	ClassPadButtonColorLevelPrefix string
}

// IDs is ids
type IDs struct {
	Master           string
	Home             string
	HomePad          string
	Slider           string
	SliderBack       string
	SliderCollection string
	Panels           []string
}

// Classes is the class names
type Classes struct {
	TabBar                  string
	Tab                     string
	TabPanel                string
	UnderTabBar             string
	Panel                   string
	Seen                    string
	UnSeen                  string
	ToBeSeen                string
	ToBeUnSeen              string
	SelectedTab             string
	UnSelectedTab           string
	PanelWithHeading        string
	PanelWithTabBar         string
	PanelHeading            string
	PanelHeadingLevelPrefix string
	InnerPanel              string
	UserContent             string

	Slider                  string
	SliderBack              string
	SliderPanel             string
	SliderPanelInner        string
	SliderButtonPad         string
	SliderPanelInnerSibling string

	CookieCrumb            string
	CookieCrumbLevelPrefix string
}

// Attributes are attributes other than classes.
type Attributes struct {
	BackID         string
	BackColorLevel string
}

// Builder builds html.
type Builder struct {
	Name              string
	Title             string
	ImportPath        string
	Repos             []string
	Services          []*Service
	panel             *Panel
	indentAmount      uint
	sliderPanelIndent uint
	Classes           *Classes
	Attributes        *Attributes
	IDs               *IDs
	Colors            *Colors
	aboutMapPointers  *AboutMapPointers // not sure i need to keep these
	AboutIDs          *AboutIDs
	markedUp          bool
}

// NewBuilder constructs a new builder.
func NewBuilder() *Builder {
	return &Builder{
		panel:        newPanel(),
		indentAmount: 2,
		Classes: &Classes{
			Tab:           classTab,
			SelectedTab:   classSelected,
			UnSelectedTab: classUnSelected,
			TabPanel:      classTabPanel,

			TabBar:      classTabBar,
			UnderTabBar: classUnderTabBar,

			Panel:                   classPanel,
			PanelWithHeading:        classPanelWithHeading,
			PanelWithTabBar:         classPanelWithTabBar,
			PanelHeading:            classPanelHeading,
			PanelHeadingLevelPrefix: classPanelHeadingLevelPrefix,
			InnerPanel:              classInnerPanel,
			UserContent:             classUserContent,

			Slider:                  classSlider,
			SliderPanel:             classSliderPanel,
			SliderPanelInner:        classSliderPanelInner,
			SliderButtonPad:         classSliderButtonPad,
			SliderPanelInnerSibling: classSliderPanelInnerSibling,

			Seen:       classSeen,
			UnSeen:     classUnSeen,
			ToBeSeen:   classToBeSeen,
			ToBeUnSeen: classToBeUnSeen,

			CookieCrumb:            classCookieCrumb,
			CookieCrumbLevelPrefix: classCookieCrumbLevelPrefix,
		},
		Attributes: &Attributes{
			BackID:         attributeBackID,
			BackColorLevel: attributeBackColorLevel,
		},
		Colors: &Colors{
			ClassBackColorLevelPrefix:      classBackColorLevelPrefix,
			ClassPadColorLevelPrefix:       classPadColorLevelPrefix,
			ClassPadButtonColorLevelPrefix: classPadButtonColorLevelPrefix,
		},
		IDs: &IDs{
			Panels: make([]string, 0, 5),
		},
	}
}

func (builder *Builder) incIndent(i uint) uint {
	return i + builder.indentAmount
}

func (builder *Builder) decIndent(i uint) uint {
	if i > builder.indentAmount {
		return i - builder.indentAmount
	}
	return i
}

// ToHTML converts services to html
func (builder *Builder) ToHTML(masterid string, indent uint, addLocations bool) string {
	if builder.markedUp {
		panic("builder already marked up")
	}
	builder.markedUp = true
	// convert "{service}" to service.Name
	for _, service := range builder.Services {
		service.cleanMarkup()
	}
	builder.panel.ID = "home" + suffixPanel
	builder.panel.HTMLID = fmt.Sprintf("%s-%s", masterid, builder.panel.ID)
	builder.IDs.Master = masterid
	indent2 := builder.incIndent(indent)
	lines := make([]string, 0, 5)
	// open master
	lines = append(lines, indentationString[:indent]+fmt.Sprintf(`<div id="%s">`, masterid))
	// heading
	lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`<h1 class="%s">%s</h1>`, classPanelHeading, builder.Title))
	// home
	lines = append(lines, builder.toHomeHTML(indent2, addLocations))
	// slider
	lines = append(lines, builder.toSliderCollectionHTML(indent2, addLocations))
	// close master
	lines = append(lines, indentationString[:indent]+fmt.Sprintf(`</div> <!-- end of #%s -->`, masterid))
	html := strings.Join(lines, newline)
	if builder.aboutMapPointers != nil {
		builder.AboutIDs = &AboutIDs{
			ButtonName:                  builder.aboutMapPointers.AboutButton.ID,
			ButtonID:                    builder.aboutMapPointers.AboutButton.HTMLID,
			ButtonPanelID:               builder.aboutMapPointers.AboutButton.PanelHTMLID,
			DefaultPanelID:              builder.aboutMapPointers.DefaultPanel.HTMLID,
			ReleasesTabPanelInnerID:     builder.aboutMapPointers.ReleasesTab.Panels[0].HTMLID,
			ContributorsTabPanelInnerID: builder.aboutMapPointers.ContributorsTab.Panels[0].HTMLID,
			CreditsTabPanelInnerID:      builder.aboutMapPointers.CreditsTab.Panels[0].HTMLID,
			LicensesTabPanelInnerID:     builder.aboutMapPointers.LicensesTab.Panels[0].HTMLID,
		}
	}
	return html
}

func (builder *Builder) toHomeHTML(indent uint, addLocations bool) string {
	lines := make([]string, 0, 5)
	// open home
	builder.IDs.Home = builder.IDs.Master + "-home"
	lines = append(lines, indentationString[:indent]+fmt.Sprintf(`<div id="%s">`, builder.IDs.Home))
	// open button pad
	builder.IDs.HomePad = builder.IDs.Home + "-pad"
	indent2 := builder.incIndent(indent)
	lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`<div id="%s" class="%s0">`, builder.IDs.HomePad, classPadColorLevelPrefix))
	// buttons
	indent3 := builder.incIndent(indent2)
	for _, s := range builder.Services {
		lines = append(lines, s.Button.toButtonHTML(builder.IDs.HomePad, indent3, builder.IDs.Home, s.Name))
	}
	// close button pad
	lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`</div> <!-- end of #%s -->`, builder.IDs.HomePad))
	// close home
	lines = append(lines, indentationString[:indent]+fmt.Sprintf(`</div> <!-- end of #%s -->`, builder.IDs.Home))
	return strings.Join(lines, newline)
}

func (builder *Builder) toSliderCollectionHTML(indent uint, addLocations bool) string {
	lines := make([]string, 0, 5)
	builder.IDs.Slider = builder.IDs.Home + "-slider"
	builder.IDs.SliderBack = builder.IDs.Slider + "-back"
	builder.IDs.SliderCollection = builder.IDs.Slider + "-collection"
	// open slider
	lines = append(lines, indentationString[:indent]+fmt.Sprintf(`<div id="%s" class="%s">`, builder.IDs.Slider, classUnSeen))
	// back button
	indent2 := builder.incIndent(indent)
	lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`<button id="%s" class="%s%d">&#11176;</button>`, builder.IDs.SliderBack, classBackColorLevelPrefix, 0))
	// open slider panel collection
	lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`<div id="%s">`, builder.IDs.SliderCollection))
	builder.sliderPanelIndent = builder.incIndent(indent2)
	// slider panels for every button
	locations := make([]string, 0, 10)
	indent3 := builder.incIndent(indent2)
	for _, s := range builder.Services {
		//locations[0] = s.Button.Location
		lines = append(lines, builder.toButtonSliderPanelsHTML(s.Name, s.Button, locations, builder.IDs.Home, indent3, true, addLocations))
	}
	// close slider panel collection
	lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`</div> <!-- end of #%s -->`, builder.IDs.SliderCollection))
	// close slider
	lines = append(lines, indentationString[:indent]+fmt.Sprintf(`</div> <!-- end of #%s -->`, builder.IDs.Slider))
	return strings.Join(lines, newline)
}

func (builder *Builder) toSliderButtonPadPanelHTML(serviceName string, panel *Panel, locations []string, heading string, seen bool, addLocations bool) string {
	colorLevelUint := uint(len(locations))
	colorLevel := serviceName
	var backLevel string
	if colorLevelUint == 1 {
		backLevel = "0"
	} else {
		backLevel = colorLevel
	}
	//backLevel := colorLevel - 1 // color level of the previous button pad
	// keep track how far the color levels go for css.
	if builder.Colors.LastColorLevel < colorLevelUint {
		builder.Colors.LastColorLevel = colorLevelUint
	}
	var visibility string
	if seen {
		visibility = classUnSeen + spaceString + classToBeSeen
	} else {
		visibility = classUnSeen + spaceString + classToBeUnSeen
	}
	// this panel is a slider panel.
	lines := make([]string, 0, 5)
	// open slider panel
	lines = append(lines, indentationString[:builder.sliderPanelIndent]+fmt.Sprintf(`<div id="%s" class="%s %s" %s="%s%s">`, panel.HTMLID, classSliderPanel, visibility, attributeBackColorLevel, classBackColorLevelPrefix, backLevel))
	indent2 := builder.incIndent(builder.sliderPanelIndent)
	// cookie crumbs
	if addLocations {
		lines = append(lines, builder.cookieCrumbs(locations, indent2))
	}
	// under the cookie crumbs is the heading
	lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`<h2 class="%s %s%s">%s</h2>`, classPanelHeading, classPanelHeadingLevelPrefix, colorLevel, heading))
	// under the header is the inner panel
	// inner panels can scroll if needed.
	innerID := panel.HTMLID + dashInnerString
	// open inner
	lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`<div id="%s" class="%s  %s%s">`, innerID, classSliderPanelInner, classPadColorLevelPrefix, colorLevel))
	indent3 := builder.incIndent(indent2)
	// make the button pad
	// open button pad
	buttonPadID := innerID + dashButtonPadString
	lines = append(lines, indentationString[:indent3]+fmt.Sprintf(`<div id="%s" class="%s %s%s">`, buttonPadID, classSliderButtonPad, classPadColorLevelPrefix, colorLevel))
	// buttons
	indent4 := builder.incIndent(indent3)
	for _, b := range panel.Buttons {
		lines = append(lines, b.toButtonHTML(panel.HTMLID, indent4, panel.HTMLID, colorLevel))
	}
	// close button pad
	lines = append(lines, indentationString[:indent3]+fmt.Sprintf(`</div> <!-- end of #%s -->`, buttonPadID))
	// close inner
	lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`</div> <!-- end of #%s -->`, innerID))
	// close slider panel
	lines = append(lines, indentationString[:builder.sliderPanelIndent]+fmt.Sprintf(`</div> <!-- end of #%s -->`, panel.HTMLID))
	return strings.Join(lines, newline)
}

func (builder *Builder) toSliderMarkupPanelHTML(serviceName string, panel *Panel, button *Button, locations []string, seen bool, addLocations bool) (string, string) {
	colorLevelUint := uint(len(locations))
	colorLevel := serviceName
	var backLevel string
	if colorLevelUint == 1 {
		backLevel = "0"
	} else {
		backLevel = colorLevel
	}
	// keep track how far the color levels go for css.
	if builder.Colors.LastColorLevel < colorLevelUint {
		builder.Colors.LastColorLevel = colorLevelUint
	}
	var visibility string
	if seen {
		visibility = classUnSeen + spaceString + classToBeSeen
	} else {
		visibility = classUnSeen + spaceString + classToBeUnSeen
	}
	// this panel is a slider panel.
	lines := make([]string, 0, 5)
	// open slider panel
	lines = append(lines, indentationString[:builder.sliderPanelIndent]+fmt.Sprintf(`<div id="%s" class="%s %s" %s="%s%s">`, panel.HTMLID, classSliderPanel, visibility, attributeBackColorLevel, classBackColorLevelPrefix, backLevel))
	indent2 := builder.incIndent(builder.sliderPanelIndent)
	// cookie crumbs
	if addLocations {
		lines = append(lines, builder.cookieCrumbs(locations, indent2))
	}
	// under the cookie crumbs is the heading
	lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`<h2 class="%s %s%s">%s</h2>`, classPanelHeading, classPanelHeadingLevelPrefix, colorLevel, button.Heading))
	// under the header is the inner panel
	indent3 := builder.incIndent(indent2)
	// inner panels have the rounded shape
	innerID := panel.HTMLID + dashInnerString
	// open inner
	lines = append(lines, indentationString[:indent3]+fmt.Sprintf(`<div id="%s" class="%s %s%s">`, innerID, classSliderPanelInner, classPadColorLevelPrefix, colorLevel))
	// under the inner panel is the user content panel
	indent4 := builder.incIndent(indent3)
	contentID := innerID + dashContentString
	// open user content and add the designer's comments, close user content
	lines = append(lines, indentationString[:indent4]+fmt.Sprintf(`<div id="%s" class="%s">`, contentID, classUserContent))
	l := len(button.Panels)
	var forwhat string
	if l == 1 {
		forwhat = fmt.Sprintf("This panel is displayed when the %q button is clicked.", button.Label)
	} else {
		forwhat = fmt.Sprintf("This is one of a group of %d panels displayed when the %q button is clicked.", l, button.Label)
	}
	if !panel.Generated {
		// make a template for the markup
		panel.Template = panel.markItUp(forwhat, button.Panels)
		lines = append(lines, fmt.Sprintf(`{{template "%s.tmpl"}}`, panel.Name))
	} else {
		// put the markup in the html
		lines = append(lines, panel.markItUp(forwhat, button.Panels))
	}
	lines = append(lines, indentationString[:indent4]+fmt.Sprintf(`</div> <!-- end of #%s -->`, contentID))
	// close inner
	lines = append(lines, indentationString[:indent3]+fmt.Sprintf(`</div> <!-- end of #%s -->`, innerID))
	// close slider panel
	lines = append(lines, indentationString[:builder.sliderPanelIndent]+fmt.Sprintf(`</div> <!-- end of #%s -->`, panel.HTMLID))
	return strings.Join(lines, newline), innerID
}

func (builder *Builder) toSliderTabBarPanelHTML(serviceName string, panel *Panel, locations []string, heading string, seen bool, addLocations bool) string {
	colorLevelUint := uint(len(locations))
	colorLevel := serviceName
	var backLevel string
	if colorLevelUint == 1 {
		backLevel = "0"
	} else {
		backLevel = colorLevel
	}

	// this panel is a slider panel
	lines := make([]string, 0, 5)
	var visibility string
	if seen {
		visibility = classUnSeen + spaceString + classToBeSeen
	} else {
		visibility = classUnSeen + spaceString + classToBeUnSeen
	}
	// open slider panel
	lines = append(lines, indentationString[:builder.sliderPanelIndent]+fmt.Sprintf(`<div id="%s" class="%s %s" %s="%s%s">`, panel.HTMLID, classSliderPanel, visibility, attributeBackColorLevel, classBackColorLevelPrefix, backLevel))
	indent2 := builder.incIndent(builder.sliderPanelIndent)
	// cookie crumbs
	if addLocations && len(locations) > 0 {
		lines = append(lines, builder.cookieCrumbs(locations, indent2))
	}
	// under the cookie crumbs is the heading
	lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`<h2 class="%s %s%s">%s</h2>`, classPanelHeading, classPanelHeadingLevelPrefix, colorLevel, heading))
	// under the header is the inner panel
	// inner panels can scroll if needed.
	innerID := panel.HTMLID + dashInnerString
	// open inner
	lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`<div id="%s" class="%s %s%s">`, innerID, classSliderPanelInner, classPadColorLevelPrefix, colorLevel))
	indent3 := builder.incIndent(indent2)
	// make the tab bar
	lines = append(lines, builder.toTabBarHTML(panel, indent3, true))
	// close inner
	lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`</div> <!-- end of #%s -->`, innerID))
	// close slider panel
	lines = append(lines, indentationString[:builder.sliderPanelIndent]+fmt.Sprintf(`</div> <!-- end of #%s -->`, panel.HTMLID))
	return strings.Join(lines, newline)
}

func (builder *Builder) toTabBarHTML(panel *Panel, indent uint, seen bool) string {
	var visibility string
	if seen {
		visibility = classSeen + spaceString + classToBeSeen
	} else {
		visibility = classUnSeen + spaceString + classToBeUnSeen
	}
	panel.TabBarHTMLID = strings.Replace(panel.HTMLID, dashString, underscoreString, -1) + underscoreTabBar
	lines := make([]string, 0, 5)
	// open tab bar
	lines = append(lines, indentationString[:indent]+fmt.Sprintf(`<div id="%s" class="%s %s">`, panel.TabBarHTMLID, classTabBar, visibility))
	// buttons
	indent2 := builder.incIndent(indent)
	for i, t := range panel.Tabs {
		lines = append(lines, t.toButtonHTML(panel.TabBarHTMLID, indent2, (i == 0)))
	}
	// close tab bar
	lines = append(lines, indentationString[:indent]+fmt.Sprintf(`</div> <!-- end of #%s -->`, panel.TabBarHTMLID))
	// open under tab bar
	underID := panel.TabBarHTMLID + dashUnderTabBar
	lines = append(lines, indentationString[:indent]+fmt.Sprintf(`<div id="%s" class="%s">`, underID, classUnderTabBar))
	for i, t := range panel.Tabs {
		lines = append(lines, builder.toTabPanelHTML(t, indent2, (i == 0)))
	}
	// close under tab bar
	lines = append(lines, indentationString[:indent]+fmt.Sprintf(`</div> <!-- end of #%s -->`, underID))
	return strings.Join(lines, newline)
}

func (builder *Builder) cookieCrumbs(locations []string, indent uint) string {
	if len(locations) < 1 {
		return emptyString
	}
	lines := make([]string, 0, 5)
	lines = append(lines, indentationString[:indent]+fmt.Sprintf(`<div class="%s"> <!-- cookie crumbs -->`, classPanelHeading))
	indent2 := builder.incIndent(indent)
	innerLines := make([]string, 1, 5)
	innerLines[0] = indentationString[:indent2-1]
	l := len(locations)
	for i := 0; i < l; i++ {
		lines = append(lines, fmt.Sprintf(`<h2 class="%s %s%d">%s</h2>`, classCookieCrumb, classCookieCrumbLevelPrefix, (i+1), locations[i]))
	}
	lines = append(lines, strings.Join(innerLines, spaceString))
	lines = append(lines, indentationString[:indent]+`</div> <!-- end of cookie crumbs -->`)
	return strings.Join(lines, newline)
}
