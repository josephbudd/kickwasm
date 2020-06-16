package project

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Colors are information defining css colors.
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
	UnderTab                string
	PanelWithHeading        string
	PanelWithTabBar         string
	PanelHeading            string
	PanelHeadingLevelPrefix string
	TabPanelGroup           string
	UserContent             string
	ModalUserContent        string
	ResizeMeWidth           string
	ResizeMeHeight          string
	DoNotPrint              string
	VScroll                 string
	HVScroll                string
	UserMarkup              string

	Slider                  string
	SliderBack              string
	SliderPanel             string
	SliderPanelInnerSibling string
	SliderPanelPad          string
	SliderButtonPad         string

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
	Name               string
	Title              string
	ImportPath         string
	Homes              []*Button
	panel              *Panel
	sliderPanelIndent  uint
	Classes            *Classes
	Attributes         *Attributes
	IDs                *IDs
	Colors             *Colors
	markedUp           bool
	SitePackImportPath string
	SitePackPackage    string
	MarkupPanelCount   uint64
}

// NewBuilder constructs a new builder.
func NewBuilder() *Builder {
	return &Builder{
		panel: newPanel(),
		Classes: &Classes{
			Tab:         classTab,
			SelectedTab: classSelected,
			UnderTab:    classUnderTab,
			TabPanel:    classTabPanel,

			TabBar:      classTabBar,
			UnderTabBar: classUnderTabBar,

			Panel:                   classPanel,
			PanelWithHeading:        classPanelWithHeading,
			PanelWithTabBar:         classPanelWithTabBar,
			PanelHeading:            classPanelHeading,
			PanelHeadingLevelPrefix: classPanelHeadingLevelPrefix,
			TabPanelGroup:           classTabPanelGroup,
			UserContent:             classUserContent,
			ModalUserContent:        classModalUserContent,
			ResizeMeWidth:           classResizeMeWidthClassName,
			ResizeMeHeight:          classResizeMeHeightClassName,
			DoNotPrint:              classDoNotPrintClassName,
			VScroll:                 classVScroll,
			HVScroll:                classHVScroll,
			UserMarkup:              classUserMarkup,

			Slider:                  classSlider,
			SliderPanel:             classSliderPanel,
			SliderPanelInnerSibling: classSliderPanelInnerSibling,
			SliderPanelPad:          classSliderPanelPad,
			SliderButtonPad:         classSliderButtonPad,

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

// ToHTML converts homes to html
func (builder *Builder) ToHTML(masterid string, addLocations bool) (markup string, err error) {
	node := builder.ToHTMLNode(masterid, addLocations)
	bb := &bytes.Buffer{}
	if err = html.Render(bb, node); err != nil {
		return
	}
	markup = bb.String()
	return
}

// ToHTMLNode converts the builder to an html node.
func (builder *Builder) ToHTMLNode(masterid string, addLocations bool) (master *html.Node) {
	if builder.markedUp {
		panic("builder already marked up")
	}
	builder.markedUp = true
	// convert "{home}" to home.Name
	for _, homeButton := range builder.Homes {
		cleanHomeMarkup(homeButton)
	}
	builder.panel.ID = "home" + suffixPanel
	builder.panel.HTMLID = fmt.Sprintf("%s-%s", masterid, builder.panel.ID)
	builder.IDs.Master = masterid
	// the master home panel
	master = &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr:     []html.Attribute{{Key: "id", Val: masterid}},
	}
	// add the h1 to the master home panel
	heading := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.H1,
		Data:     "h1",
		Attr:     []html.Attribute{{Key: "class", Val: classPanelHeading}},
	}
	textNode := &html.Node{
		Type: html.TextNode,
		Data: builder.Title,
	}
	heading.AppendChild(textNode)
	master.AppendChild(heading)
	// add the home div to the master.
	home := builder.toHomeHTML(addLocations)
	master.AppendChild(home)
	// add the slider and its collection to the home div.
	slider := builder.toSliderCollectionHTML(addLocations)
	master.AppendChild(slider)
	return
}

func (builder *Builder) toHomeHTML(addLocations bool) (home *html.Node) {
	// the master home
	builder.IDs.Home = builder.IDs.Master + "-home"
	home = &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr:     []html.Attribute{{Key: "id", Val: builder.IDs.Home}},
	}
	// add the home button pad to the master home pad
	// buttons
	builder.IDs.HomePad = builder.IDs.Home + "-pad"
	homePad := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			{Key: "id", Val: builder.IDs.HomePad},
			{Key: "class", Val: fmt.Sprintf("%s0", classPadColorLevelPrefix)},
		},
	}
	home.AppendChild(homePad)
	// add each home's button to the home button pad.
	for _, homeButton := range builder.Homes {
		button := homeButton.toButtonHTML(builder.IDs.HomePad, builder.IDs.Home, homeButton.ID)
		homePad.AppendChild(button)
	}
	return
}

func (builder *Builder) toSliderCollectionHTML(addLocations bool) (slider *html.Node) {
	builder.IDs.Slider = builder.IDs.Home + "-slider"
	builder.IDs.SliderBack = builder.IDs.Slider + "-back"
	builder.IDs.SliderCollection = builder.IDs.Slider + "-collection"
	// open slider
	slider = &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			{Key: "id", Val: builder.IDs.Slider},
			{Key: "class", Val: classUnSeen},
		},
	}
	// add the back button to the slider
	backButton := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Button,
		Data:     "button",
		Attr: []html.Attribute{
			{Key: "id", Val: builder.IDs.SliderBack},
			{Key: "class", Val: fmt.Sprintf("%s%d", classBackColorLevelPrefix, 0)},
		},
	}
	textNode := &html.Node{
		Type: html.TextNode,
		// Data: html.UnescapeString("&#11176;"),
		Data: "↩",
	}
	backButton.AppendChild(textNode)
	slider.AppendChild(backButton)
	// add the slider panel collection to the slider
	sliderCollection := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			{Key: "id", Val: builder.IDs.SliderCollection},
		},
	}
	slider.AppendChild(sliderCollection)
	// add each home's slider panels to the collection.
	// slider panels for every button
	locations := make([]string, 0, 10)
	for _, homeButton := range builder.Homes {
		sliderButtonPanels := builder.toButtonSliderPanelsHTML(homeButton.ID, homeButton, locations, builder.IDs.Home, true, addLocations)
		for _, p := range sliderButtonPanels {
			sliderCollection.AppendChild(p)
		}
	}
	// close slider panel collection
	// close slider
	return
}

func (builder *Builder) toSliderButtonPadPanelHTML(homeName string, panel *Panel, locations []string, heading string, seen bool, addLocations bool) (sliderPanel *html.Node) {
	colorLevelUint := uint(len(locations))
	var backLevel string
	if colorLevelUint == 1 {
		backLevel = "0"
	} else {
		backLevel = homeName
	}
	//backLevel := homeName - 1 // color level of the previous button pad
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
	sliderPanel = &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			{Key: "id", Val: panel.HTMLID},
			{Key: "class", Val: fmt.Sprintf("%s %s", classSliderPanel, visibility)},
			{Key: attributeBackColorLevel, Val: fmt.Sprintf("%s%s", classBackColorLevelPrefix, backLevel)},
		},
	}
	// add the cookie crumbs inside the slider
	if addLocations {
		if cc := builder.cookieCrumbs(homeName, locations); cc != nil {
			sliderPanel.AppendChild(cc)
		}
	}
	// add the heading inside the slider
	h2 := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.H2,
		Data:     "h2",
		Attr: []html.Attribute{
			{Key: "class", Val: fmt.Sprintf("%s %s%s", classPanelHeading, classPanelHeadingLevelPrefix, homeName)},
		},
	}
	textNode := &html.Node{
		Type: html.TextNode,
		Data: heading,
	}
	h2.AppendChild(textNode)
	sliderPanel.AppendChild(h2)
	// add the panel pad inside the slider
	innerPanelID := panel.HTMLID + DashInnerString
	innerPanel := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			{Key: "id", Val: innerPanelID},
			{Key: "class", Val: fmt.Sprintf("%s %s%s", classSliderPanelPad, classPadColorLevelPrefix, homeName)},
		},
	}
	sliderPanel.AppendChild(innerPanel)
	// add the buttons inside the pad
	buttonPadID := innerPanelID + dashButtonPadString
	buttonPad := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			{Key: "id", Val: buttonPadID},
			{Key: "class", Val: fmt.Sprintf("%s %s%s", classSliderButtonPad, classPadColorLevelPrefix, homeName)},
		},
	}
	innerPanel.AppendChild(buttonPad)
	// add the buttons inside the button pad
	for _, b := range panel.Buttons {
		button := b.toButtonHTML(panel.HTMLID, panel.HTMLID, homeName)
		buttonPad.AppendChild(button)
	}
	// close button pad
	// close inner
	// close slider panel
	return
}

func (builder *Builder) toSliderMarkupPanelHTML(homeName string, panel *Panel, button *Button, locations []string, seen bool, addLocations bool) (sliderMarkupPanel *html.Node, innerPanelID string) {
	colorLevelUint := uint(len(locations))
	var backLevel string
	if colorLevelUint == 1 {
		backLevel = "0"
	} else {
		backLevel = homeName
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
	sliderMarkupPanel = &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			{Key: "id", Val: panel.HTMLID},
			{Key: "class", Val: fmt.Sprintf("%s %s", classSliderPanel, visibility)},
			{Key: attributeBackColorLevel, Val: fmt.Sprintf("%s%s", classBackColorLevelPrefix, backLevel)},
		},
	}
	// add the cookie crumbs inside the slider markup panel.
	if addLocations {
		if cc := builder.cookieCrumbs(homeName, locations); cc != nil {
			sliderMarkupPanel.AppendChild(cc)
		}
	}
	// add the h2 inside the slider panel.
	h2 := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.H2,
		Data:     "h2",
		Attr: []html.Attribute{
			{Key: "class", Val: fmt.Sprintf("%s %s%s", classPanelHeading, classPanelHeadingLevelPrefix, homeName)},
		},
	}
	textNode := &html.Node{
		Type: html.TextNode,
		Data: button.Heading,
	}
	h2.AppendChild(textNode)
	sliderMarkupPanel.AppendChild(h2)
	// add the pad inside the slider panel.
	innerPanelID = panel.HTMLID + DashInnerString
	innerPanel := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			{Key: "id", Val: innerPanelID},
			{Key: "class", Val: fmt.Sprintf("%s %s%s", classSliderPanelPad, classPadColorLevelPrefix, homeName)},
		},
	}
	sliderMarkupPanel.AppendChild(innerPanel)
	// the user content panel.
	var scroll string
	if panel.HVScroll {
		scroll = classHVScroll
	} else {
		scroll = classVScroll
	}
	contentID := innerPanelID + dashContentString
	userContentPanel := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			{Key: "id", Val: contentID},
			{Key: "class", Val: fmt.Sprintf("%s %s", classUserContent, scroll)},
		},
	}
	innerPanel.AppendChild(userContentPanel)
	attributes := make([]html.Attribute, 0, 1)
	if !panel.HVScroll {
		// this panel will not horizontally scroll so size the width.
		attributes = append(attributes, html.Attribute{Key: "class", Val: classResizeMeWidthClassName})
	}
	markupPanel := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr:     attributes,
	}
	userContentPanel.AppendChild(markupPanel)
	l := len(button.Panels)
	var forwhat string
	if l == 1 {
		forwhat = fmt.Sprintf("This panel is displayed when the %q button is clicked.", button.Label)
	} else {
		forwhat = fmt.Sprintf("This is one of a group of %d panels displayed when the %q button is clicked.", l, button.Label)
	}
	// The markup for this panel is defined in the panel's yaml file.
	// The panel's markup will be in a template file.
	panel.Template = panel.markItUp(forwhat, button.Panels)
	// Put template code linking to the template file in this markup panel.
	templateLink := &html.Node{
		Type: html.TextNode,
		Data: html.UnescapeString(fmt.Sprintf(`{{template "%s.tmpl"}}`, panel.Name)),
	}
	markupPanel.AppendChild(templateLink)
	// close inner
	// close slider panel
	return
}

func (builder *Builder) toSliderTabBarPanelHTML(homeName string, panel *Panel, locations []string, heading string, seen bool, addLocations bool) (sliderPanel *html.Node) {
	colorLevelUint := uint(len(locations))
	var backLevel string
	if colorLevelUint == 1 {
		backLevel = "0"
	} else {
		backLevel = homeName
	}
	// this panel is a slider panel
	var visibility string
	if seen {
		visibility = classUnSeen + spaceString + classToBeSeen
	} else {
		visibility = classUnSeen + spaceString + classToBeUnSeen
	}
	// this panel is a slider panel
	sliderPanel = &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			{Key: "id", Val: panel.HTMLID},
			{Key: "class", Val: fmt.Sprintf("%s %s", classSliderPanel, visibility)},
			{Key: attributeBackColorLevel, Val: fmt.Sprintf("%s%s", classBackColorLevelPrefix, backLevel)},
		},
	}
	// add the cookie crumbs inside the slider panel
	if addLocations && len(locations) > 0 {
		if cc := builder.cookieCrumbs(homeName, locations); cc != nil {
			sliderPanel.AppendChild(cc)
		}
	}
	// add the heading inside the slider panel
	h2 := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.H2,
		Data:     "h2",
		Attr: []html.Attribute{
			{Key: "class", Val: fmt.Sprintf("%s %s%s", classPanelHeading, classPanelHeadingLevelPrefix, homeName)},
		},
	}
	textNode := &html.Node{
		Type: html.TextNode,
		Data: heading,
	}
	h2.AppendChild(textNode)
	sliderPanel.AppendChild(h2)
	// add the inner panel inside the slider panel
	innerPanelID := panel.HTMLID + DashInnerString
	innerPanel := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			{Key: "id", Val: innerPanelID},
			{Key: "class", Val: fmt.Sprintf("%s %s%s", classSliderPanelPad, classPadColorLevelPrefix, homeName)},
		},
	}
	sliderPanel.AppendChild(innerPanel)
	// add the tab bar inside the slider panel
	tabBar, underTabBar := builder.toTabBarHTML(panel, true)
	innerPanel.AppendChild(tabBar)
	innerPanel.AppendChild(underTabBar)
	// close inner
	// close slider panel
	return
}

func (builder *Builder) toTabBarHTML(panel *Panel, seen bool) (tabBarPanel, underTabBarPanel *html.Node) {
	var t *Tab
	// how many real buttons are there in this tab bar?
	buttonCount := len(panel.Tabs)
	for _, t := range panel.Tabs {
		if t.Spawn {
			buttonCount--
		}
	}
	panel.HasRealTabs = buttonCount > 0
	if !panel.HasRealTabs {
		// no buttons so this tab bar is not seen.
		seen = false
	}
	var visibility string
	if seen {
		visibility = classSeen + spaceString + classToBeSeen
	} else {
		visibility = classUnSeen + spaceString + classToBeUnSeen
	}
	panel.TabBarHTMLID = strings.Replace(panel.HTMLID, dashString, underscoreString, -1) + underscoreTabBar
	// this panel is a tab bar
	tabBarPanel = &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			{Key: "id", Val: panel.TabBarHTMLID},
			{Key: "class", Val: fmt.Sprintf("%s %s", classTabBar, visibility)},
		},
	}
	// insert the real buttons inside the tab bar panel
	position := 0
	for _, t = range panel.Tabs {
		if !t.Spawn {
			button := t.toButtonHTML(panel.TabBarHTMLID, (position == 0))
			position++
			tabBarPanel.AppendChild(button)
		} else {
			// Spawned tab so don't create a tab.
			// Only set t.HTMLID.
			_ = t.toButtonHTML(panel.TabBarHTMLID, false)
		}
	}
	// under tab bar panel
	panel.UnderTabBarHTMLID = panel.TabBarHTMLID + DashUnderTabBar
	underTabBarPanel = &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			{Key: "id", Val: panel.UnderTabBarHTMLID},
			{Key: "class", Val: classUnderTabBar},
		},
	}
	if panel.HasRealTabs {
		// add the markup panels to the under tab bar panel.
		position := 0
		for _, t = range panel.Tabs {
			if !t.Spawn {
				p := builder.toTabPanelHTML(t, (position == 0))
				underTabBarPanel.AppendChild(p)
				position++
			} else {
				// Spawned tab so don't create a tab.
				// Only set t.HTMLID.
				_ = builder.toTabPanelHTML(t, false)
			}
		}
	}
	return
}

func (builder *Builder) cookieCrumbs(homeName string, locations []string) (cc *html.Node) {
	if len(locations) < 1 {
		return nil
	}
	cc = &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			{Key: "class", Val: classPanelHeading},
		},
	}
	l := len(locations)
	for i := 0; i < l; i++ {
		h2 := &html.Node{
			Type:     html.ElementNode,
			DataAtom: atom.H2,
			Data:     "h2",
			Attr: []html.Attribute{
				{Key: "class", Val: fmt.Sprintf("%s %s%s", classCookieCrumb, classCookieCrumbLevelPrefix, homeName)},
			},
		}
		textNode := &html.Node{
			Type: html.TextNode,
			Data: locations[i],
		}
		h2.AppendChild(textNode)
		cc.AppendChild(h2)
	}
	return
}
