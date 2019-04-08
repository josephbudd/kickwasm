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
	UnSelectedTab           string
	PanelWithHeading        string
	PanelWithTabBar         string
	PanelHeading            string
	PanelHeadingLevelPrefix string
	InnerPanel              string
	UserContent             string
	ModalUserContent        string
	CloserUserContent       string
	ResizeMeWidth           string

	Slider                  string
	SliderBack              string
	SliderPanel             string
	SliderPanelInner        string
	SliderButtonPad         string
	SliderPanelInnerSibling string
	ResizeMeWidthClassName  string

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
	Stores             []string
	Services           []*Service
	panel              *Panel
	sliderPanelIndent  uint
	Classes            *Classes
	Attributes         *Attributes
	IDs                *IDs
	Colors             *Colors
	markedUp           bool
	SitePackImportPath string
	SitePackPackage    string
}

// NewBuilder constructs a new builder.
func NewBuilder() *Builder {
	return &Builder{
		panel: newPanel(),
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
			ModalUserContent:        classModalUserContent,
			CloserUserContent:       classCloserUserContent,
			ResizeMeWidth:           classResizeMeWidthClassName,

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

// ToHTML converts services to html
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
	// convert "{service}" to service.Name
	for _, service := range builder.Services {
		service.cleanMarkup()
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
	// add each service's button to the home button pad.
	for _, s := range builder.Services {
		button := s.Button.toButtonHTML(builder.IDs.HomePad, builder.IDs.Home, s.Name)
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
	// add each service's slider panels to the collection.
	// slider panels for every button
	locations := make([]string, 0, 10)
	for _, s := range builder.Services {
		//locations[0] = s.Button.Location
		sliderButtonPanels := builder.toButtonSliderPanelsHTML(s.Name, s.Button, locations, builder.IDs.Home, true, addLocations)
		for _, p := range sliderButtonPanels {
			sliderCollection.AppendChild(p)
		}
	}
	// close slider panel collection
	// close slider
	return
}

func (builder *Builder) toSliderButtonPadPanelHTML(serviceName string, panel *Panel, locations []string, heading string, seen bool, addLocations bool) (sliderPanel *html.Node) {
	colorLevelUint := uint(len(locations))
	var backLevel string
	if colorLevelUint == 1 {
		backLevel = "0"
	} else {
		backLevel = serviceName
	}
	//backLevel := serviceName - 1 // color level of the previous button pad
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
		if cc := builder.cookieCrumbs(serviceName, locations); cc != nil {
			sliderPanel.AppendChild(cc)
		}
	}
	// add the heading inside the slider
	h2 := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.H2,
		Data:     "h2",
		Attr: []html.Attribute{
			{Key: "class", Val: fmt.Sprintf("%s %s%s", classPanelHeading, classPanelHeadingLevelPrefix, serviceName)},
		},
	}
	textNode := &html.Node{
		Type: html.TextNode,
		Data: heading,
	}
	h2.AppendChild(textNode)
	sliderPanel.AppendChild(h2)
	// add the inner panel inside the slider
	innerPanelID := panel.HTMLID + dashInnerString
	innerPanel := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			{Key: "id", Val: innerPanelID},
			{Key: "class", Val: fmt.Sprintf("%s %s%s", classSliderPanelInner, classPadColorLevelPrefix, serviceName)},
		},
	}
	sliderPanel.AppendChild(innerPanel)
	// add the button pad inside the inner panel
	buttonPadID := innerPanelID + dashButtonPadString
	buttonPad := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			{Key: "id", Val: buttonPadID},
			{Key: "class", Val: fmt.Sprintf("%s %s%s", classSliderButtonPad, classPadColorLevelPrefix, serviceName)},
		},
	}
	innerPanel.AppendChild(buttonPad)
	// add the buttons inside the button pad
	for _, b := range panel.Buttons {
		button := b.toButtonHTML(panel.HTMLID, panel.HTMLID, serviceName)
		buttonPad.AppendChild(button)
	}
	// close button pad
	// close inner
	// close slider panel
	return
}

func (builder *Builder) toSliderMarkupPanelHTML(serviceName string, panel *Panel, button *Button, locations []string, seen bool, addLocations bool) (sliderMarkupPanel *html.Node, innerPanelID string) {
	colorLevelUint := uint(len(locations))
	var backLevel string
	if colorLevelUint == 1 {
		backLevel = "0"
	} else {
		backLevel = serviceName
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
	// this panel is a slider markup panel.
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
		if cc := builder.cookieCrumbs(serviceName, locations); cc != nil {
			sliderMarkupPanel.AppendChild(cc)
		}
	}
	// add the h2 inside the slider markup panel.
	h2 := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.H2,
		Data:     "h2",
		Attr: []html.Attribute{
			{Key: "class", Val: fmt.Sprintf("%s %s%s", classPanelHeading, classPanelHeadingLevelPrefix, serviceName)},
		},
	}
	textNode := &html.Node{
		Type: html.TextNode,
		Data: button.Heading,
	}
	h2.AppendChild(textNode)
	sliderMarkupPanel.AppendChild(h2)
	// add the inner panel inside the slider markup panel.
	innerPanelID = panel.HTMLID + dashInnerString
	innerPanel := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			{Key: "id", Val: innerPanelID},
			{Key: "class", Val: fmt.Sprintf("%s %s%s", classSliderPanelInner, classPadColorLevelPrefix, serviceName)},
		},
	}
	sliderMarkupPanel.AppendChild(innerPanel)
	// add the user markup panel inside the inner panel.
	contentID := innerPanelID + dashContentString
	markupPanel := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			{Key: "id", Val: contentID},
			{Key: "class", Val: classUserContent},
		},
	}
	innerPanel.AppendChild(markupPanel)
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

func (builder *Builder) toSliderTabBarPanelHTML(serviceName string, panel *Panel, locations []string, heading string, seen bool, addLocations bool) (sliderPanel *html.Node) {
	colorLevelUint := uint(len(locations))
	var backLevel string
	if colorLevelUint == 1 {
		backLevel = "0"
	} else {
		backLevel = serviceName
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
		if cc := builder.cookieCrumbs(serviceName, locations); cc != nil {
			sliderPanel.AppendChild(cc)
		}
	}
	// add the heading inside the slider panel
	h2 := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.H2,
		Data:     "h2",
		Attr: []html.Attribute{
			{Key: "class", Val: fmt.Sprintf("%s %s%s", classPanelHeading, classPanelHeadingLevelPrefix, serviceName)},
		},
	}
	textNode := &html.Node{
		Type: html.TextNode,
		Data: heading,
	}
	h2.AppendChild(textNode)
	sliderPanel.AppendChild(h2)
	// add the inner panel inside the slider panel
	innerPanelID := panel.HTMLID + dashInnerString
	innerPanel := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			{Key: "id", Val: innerPanelID},
			{Key: "class", Val: fmt.Sprintf("%s %s%s", classSliderPanelInner, classPadColorLevelPrefix, serviceName)},
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
	// insert the buttons inside the tab bar panel
	var i int
	var t *Tab
	for i, t = range panel.Tabs {
		button := t.toButtonHTML(panel.TabBarHTMLID, (i == 0))
		tabBarPanel.AppendChild(button)
	}
	// close tab bar
	// under tab bar panel
	underID := panel.TabBarHTMLID + dashUnderTabBar
	underTabBarPanel = &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			{Key: "id", Val: underID},
			{Key: "class", Val: classUnderTabBar},
		},
	}
	// add the markup panels to the under tab bar panel.
	for i, t = range panel.Tabs {
		panel := builder.toTabPanelHTML(t, (i == 0))
		underTabBarPanel.AppendChild(panel)
	}
	// close under tab bar
	return
}

func (builder *Builder) cookieCrumbs(serviceName string, locations []string) (cc *html.Node) {
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
				{Key: "class", Val: fmt.Sprintf("%s %s%s", classCookieCrumb, classCookieCrumbLevelPrefix, serviceName)},
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
