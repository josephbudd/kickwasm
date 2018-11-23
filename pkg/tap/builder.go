package tap

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
	Name              string
	Title             string
	ImportPath        string
	Stores            []string
	Services          []*Service
	panel             *Panel
	indentAmount      uint
	sliderPanelIndent uint
	Classes           *Classes
	Attributes        *Attributes
	IDs               *IDs
	Colors            *Colors
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
func (builder *Builder) ToHTML(masterid string, addLocations bool) (string, error) {
	node := builder.ToHTMLNode(masterid, addLocations)
	bb := &bytes.Buffer{}
	err := html.Render(bb, node)
	return bb.String(), err
}

// ToHTMLNode converts the builder to an html node.
func (builder *Builder) ToHTMLNode(masterid string, addLocations bool) *html.Node {
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
	//lines = append(lines, indentationString[:indent]+fmt.Sprintf(`<div id="%s">`, masterid))
	master := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr:     []html.Attribute{html.Attribute{Key: "id", Val: masterid}},
	}
	// add the h1 to the master home panel
	// lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`<h1 class="%s">%s</h1>`, classPanelHeading, builder.Title))
	heading := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.H1,
		Data:     "h1",
		Attr:     []html.Attribute{html.Attribute{Key: "class", Val: classPanelHeading}},
	}
	textNode := &html.Node{
		Type: html.TextNode,
		Data: builder.Title,
	}
	heading.AppendChild(textNode)
	master.AppendChild(heading)
	// add the home div to the master.
	// lines = append(lines, builder.toHomeHTML(indent2, addLocations))
	home := builder.toHomeHTML(addLocations)
	master.AppendChild(home)
	// add the slider and its collection to the home div.
	// lines = append(lines, builder.toSliderCollectionHTML(indent2, addLocations))
	slider := builder.toSliderCollectionHTML(addLocations)
	master.AppendChild(slider)
	return master
}

func (builder *Builder) toHomeHTML(addLocations bool) *html.Node {
	// the master home
	// lines = append(lines, indentationString[:indent]+fmt.Sprintf(`<div id="%s">`, builder.IDs.Home))
	builder.IDs.Home = builder.IDs.Master + "-home"
	home := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr:     []html.Attribute{html.Attribute{Key: "id", Val: builder.IDs.Home}},
	}
	// add the home button pad to the master home pad
	//lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`<div id="%s" class="%s0">`, builder.IDs.HomePad, classPadColorLevelPrefix))
	// buttons
	builder.IDs.HomePad = builder.IDs.Home + "-pad"
	homePad := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			html.Attribute{Key: "id", Val: builder.IDs.HomePad},
			html.Attribute{Key: "class", Val: fmt.Sprintf("%s0", classPadColorLevelPrefix)},
		},
	}
	home.AppendChild(homePad)
	// add each service's button to the home button pad.
	for _, s := range builder.Services {
		//lines = append(lines, s.Button.toButtonHTML(builder.IDs.HomePad, indent3, builder.IDs.Home, s.Name))
		button := s.Button.toButtonHTML(builder.IDs.HomePad, builder.IDs.Home, s.Name)
		homePad.AppendChild(button)
	}
	return home
}

func (builder *Builder) toSliderCollectionHTML(addLocations bool) *html.Node {
	builder.IDs.Slider = builder.IDs.Home + "-slider"
	builder.IDs.SliderBack = builder.IDs.Slider + "-back"
	builder.IDs.SliderCollection = builder.IDs.Slider + "-collection"
	// open slider
	// lines = append(lines, indentationString[:indent]+fmt.Sprintf(`<div id="%s" class="%s">`, builder.IDs.Slider, classUnSeen))
	slider := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			html.Attribute{Key: "id", Val: builder.IDs.Slider},
			html.Attribute{Key: "class", Val: classUnSeen},
		},
	}
	// add the back button to the slider
	//lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`<button id="%s" class="%s%d">&#11176;</button>`, builder.IDs.SliderBack, classBackColorLevelPrefix, 0))
	backButton := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Button,
		Data:     "button",
		Attr: []html.Attribute{
			html.Attribute{Key: "id", Val: builder.IDs.SliderBack},
			html.Attribute{Key: "class", Val: fmt.Sprintf("%s%d", classBackColorLevelPrefix, 0)},
		},
	}
	textNode := &html.Node{
		Type: html.TextNode,
		Data: html.UnescapeString("&#11176;"),
	}
	backButton.AppendChild(textNode)
	slider.AppendChild(backButton)
	// add the slider panel collection to the slider
	// lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`<div id="%s">`, builder.IDs.SliderCollection))
	sliderCollection := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			html.Attribute{Key: "id", Val: builder.IDs.SliderCollection},
		},
	}
	slider.AppendChild(sliderCollection)
	// add each service's slider panels to the collection.
	// builder.sliderPanelIndent = builder.incIndent(indent2)
	// slider panels for every button
	locations := make([]string, 0, 10)
	for _, s := range builder.Services {
		//locations[0] = s.Button.Location
		//lines = append(lines, builder.toButtonSliderPanelsHTML(s.Name, s.Button, locations, builder.IDs.Home, indent3, true, addLocations))
		sliderButtonPanels := builder.toButtonSliderPanelsHTML(s.Name, s.Button, locations, builder.IDs.Home, true, addLocations)
		for _, p := range sliderButtonPanels {
			sliderCollection.AppendChild(p)
		}
	}
	// close slider panel collection
	//lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`</div> <!-- end of #%s -->`, builder.IDs.SliderCollection))
	// close slider
	//lines = append(lines, indentationString[:indent]+fmt.Sprintf(`</div> <!-- end of #%s -->`, builder.IDs.Slider))
	return slider
}

func (builder *Builder) toSliderButtonPadPanelHTML(serviceName string, panel *Panel, locations []string, heading string, seen bool, addLocations bool) *html.Node {
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
	// lines = append(lines, indentationString[:builder.sliderPanelIndent]+fmt.Sprintf(`<div id="%s" class="%s %s" %s="%s%s">`, panel.HTMLID, classSliderPanel, visibility, attributeBackColorLevel, classBackColorLevelPrefix, backLevel))
	sliderPanel := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			html.Attribute{Key: "id", Val: panel.HTMLID},
			html.Attribute{Key: "class", Val: fmt.Sprintf("%s %s", classSliderPanel, visibility)},
			html.Attribute{Key: attributeBackColorLevel, Val: fmt.Sprintf("%s%s", classBackColorLevelPrefix, backLevel)},
		},
	}
	// add the cookie crumbs inside the slider
	if addLocations {
		// lines = append(lines, builder.cookieCrumbs(serviceName, locations, indent2))
		if cc := builder.cookieCrumbs(serviceName, locations); cc != nil {
			sliderPanel.AppendChild(cc)
		}
	}
	// add the heading inside the slider
	// lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`<h2 class="%s %s%s">%s</h2>`, classPanelHeading, classPanelHeadingLevelPrefix, serviceName, heading))
	h2 := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.H2,
		Data:     "h2",
		Attr: []html.Attribute{
			html.Attribute{Key: "class", Val: fmt.Sprintf("%s %s%s", classPanelHeading, classPanelHeadingLevelPrefix, serviceName)},
		},
	}
	textNode := &html.Node{
		Type: html.TextNode,
		Data: heading,
	}
	h2.AppendChild(textNode)
	sliderPanel.AppendChild(h2)
	// add the inner panel inside the slider
	innerID := panel.HTMLID + dashInnerString
	// lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`<div id="%s" class="%s  %s%s">`, innerID, classSliderPanelInner, classPadColorLevelPrefix, serviceName))
	innerPanel := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			html.Attribute{Key: "id", Val: innerID},
			html.Attribute{Key: "class", Val: fmt.Sprintf("%s %s%s", classSliderPanelInner, classPadColorLevelPrefix, serviceName)},
		},
	}
	sliderPanel.AppendChild(innerPanel)
	// add the button pad inside the inner panel
	//lines = append(lines, indentationString[:indent3]+fmt.Sprintf(`<div id="%s" class="%s %s%s">`, buttonPadID, classSliderButtonPad, classPadColorLevelPrefix, serviceName))
	buttonPadID := innerID + dashButtonPadString
	buttonPad := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			html.Attribute{Key: "id", Val: buttonPadID},
			html.Attribute{Key: "class", Val: fmt.Sprintf("%s %s%s", classSliderButtonPad, classPadColorLevelPrefix, serviceName)},
		},
	}
	innerPanel.AppendChild(buttonPad)
	// add the buttons inside the button pad
	for _, b := range panel.Buttons {
		//lines = append(lines, b.toButtonHTML(panel.HTMLID, indent4, panel.HTMLID, serviceName))
		button := b.toButtonHTML(panel.HTMLID, panel.HTMLID, serviceName)
		buttonPad.AppendChild(button)
	}
	// close button pad
	//lines = append(lines, indentationString[:indent3]+fmt.Sprintf(`</div> <!-- end of #%s -->`, buttonPadID))
	// close inner
	//lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`</div> <!-- end of #%s -->`, innerID))
	// close slider panel
	//lines = append(lines, indentationString[:builder.sliderPanelIndent]+fmt.Sprintf(`</div> <!-- end of #%s -->`, panel.HTMLID))
	return sliderPanel
}

func (builder *Builder) toSliderMarkupPanelHTML(serviceName string, panel *Panel, button *Button, locations []string, seen bool, addLocations bool) (*html.Node, string) {
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
	//lines = append(lines, indentationString[:builder.sliderPanelIndent]+fmt.Sprintf(`<div id="%s" class="%s %s" %s="%s%s">`, panel.HTMLID, classSliderPanel, visibility, attributeBackColorLevel, classBackColorLevelPrefix, backLevel))
	sliderMarkupPanel := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			html.Attribute{Key: "id", Val: panel.HTMLID},
			html.Attribute{Key: "class", Val: fmt.Sprintf("%s %s", classSliderPanel, visibility)},
			html.Attribute{Key: attributeBackColorLevel, Val: fmt.Sprintf("%s%s", classBackColorLevelPrefix, backLevel)},
		},
	}
	// add the cookie crumbs inside the slider markup panel.
	if addLocations {
		//lines = append(lines, builder.cookieCrumbs(serviceName, locations, indent2))
		if cc := builder.cookieCrumbs(serviceName, locations); cc != nil {
			sliderMarkupPanel.AppendChild(cc)
		}
	}
	// add the h2 inside the slider markup panel.
	//lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`<h2 class="%s %s%s">%s</h2>`, classPanelHeading, classPanelHeadingLevelPrefix, serviceName, button.Heading))
	h2 := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.H2,
		Data:     "h2",
		Attr: []html.Attribute{
			html.Attribute{Key: "class", Val: fmt.Sprintf("%s %s%s", classPanelHeading, classPanelHeadingLevelPrefix, serviceName)},
		},
	}
	textNode := &html.Node{
		Type: html.TextNode,
		Data: button.Heading,
	}
	h2.AppendChild(textNode)
	sliderMarkupPanel.AppendChild(h2)
	// add the inner panel inside the slider markup panel.
	innerID := panel.HTMLID + dashInnerString
	//lines = append(lines, indentationString[:indent3]+fmt.Sprintf(`<div id="%s" class="%s %s%s">`, innerID, classSliderPanelInner, classPadColorLevelPrefix, serviceName))
	innerPanel := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			html.Attribute{Key: "id", Val: innerID},
			html.Attribute{Key: "class", Val: fmt.Sprintf("%s %s%s", classSliderPanelInner, classPadColorLevelPrefix, serviceName)},
		},
	}
	sliderMarkupPanel.AppendChild(innerPanel)
	// add the user markup panel inside the inner panel.
	contentID := innerID + dashContentString
	//lines = append(lines, indentationString[:indent4]+fmt.Sprintf(`<div id="%s" class="%s">`, contentID, classUserContent))
	markupPanel := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			html.Attribute{Key: "id", Val: contentID},
			html.Attribute{Key: "class", Val: classUserContent},
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
	//lines = append(lines, fmt.Sprintf(`{{template "%s.tmpl"}}`, panel.Name))
	// Put template code linking to the template file in this markup panel.
	templateLink := &html.Node{
		Type: html.TextNode,
		Data: html.UnescapeString(fmt.Sprintf(`{{template "%s.tmpl"}}`, panel.Name)),
	}
	markupPanel.AppendChild(templateLink)
	//lines = append(lines, indentationString[:indent4]+fmt.Sprintf(`</div> <!-- end of #%s -->`, contentID))
	// close inner
	//lines = append(lines, indentationString[:indent3]+fmt.Sprintf(`</div> <!-- end of #%s -->`, innerID))
	// close slider panel
	//lines = append(lines, indentationString[:builder.sliderPanelIndent]+fmt.Sprintf(`</div> <!-- end of #%s -->`, panel.HTMLID))
	return sliderMarkupPanel, innerID
}

func (builder *Builder) toSliderTabBarPanelHTML(serviceName string, panel *Panel, locations []string, heading string, seen bool, addLocations bool) *html.Node {
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
	//lines = append(lines, indentationString[:builder.sliderPanelIndent]+fmt.Sprintf(`<div id="%s" class="%s %s" %s="%s%s">`, panel.HTMLID, classSliderPanel, visibility, attributeBackColorLevel, classBackColorLevelPrefix, backLevel))
	// indent2 := builder.incIndent(builder.sliderPanelIndent)
	sliderPanel := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			html.Attribute{Key: "id", Val: panel.HTMLID},
			html.Attribute{Key: "class", Val: fmt.Sprintf("%s %s", classSliderPanel, visibility)},
			html.Attribute{Key: attributeBackColorLevel, Val: fmt.Sprintf("%s%s", classBackColorLevelPrefix, backLevel)},
		},
	}
	// add the cookie crumbs inside the slider panel
	if addLocations && len(locations) > 0 {
		//lines = append(lines, builder.cookieCrumbs(serviceName, locations, indent2))
		if cc := builder.cookieCrumbs(serviceName, locations); cc != nil {
			sliderPanel.AppendChild(cc)
		}
	}
	// add the heading inside the slider panel
	//lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`<h2 class="%s %s%s">%s</h2>`, classPanelHeading, classPanelHeadingLevelPrefix, serviceName, heading))
	h2 := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.H2,
		Data:     "h2",
		Attr: []html.Attribute{
			html.Attribute{Key: "class", Val: fmt.Sprintf("%s %s%s", classPanelHeading, classPanelHeadingLevelPrefix, serviceName)},
		},
	}
	textNode := &html.Node{
		Type: html.TextNode,
		Data: heading,
	}
	h2.AppendChild(textNode)
	sliderPanel.AppendChild(h2)
	// add the inner panel inside the slider panel
	innerID := panel.HTMLID + dashInnerString
	//lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`<div id="%s" class="%s %s%s">`, innerID, classSliderPanelInner, classPadColorLevelPrefix, serviceName))
	innerPanel := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			html.Attribute{Key: "id", Val: innerID},
			html.Attribute{Key: "class", Val: fmt.Sprintf("%s %s%s", classSliderPanelInner, classPadColorLevelPrefix, serviceName)},
		},
	}
	sliderPanel.AppendChild(innerPanel)
	// add the tab bar inside the slider panel
	// lines = append(lines, builder.toTabBarHTML(panel, indent3, true))
	tabBar, underTabBar := builder.toTabBarHTML(panel, true)
	innerPanel.AppendChild(tabBar)
	innerPanel.AppendChild(underTabBar)
	// close inner
	//lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`</div> <!-- end of #%s -->`, innerID))
	// close slider panel
	//lines = append(lines, indentationString[:builder.sliderPanelIndent]+fmt.Sprintf(`</div> <!-- end of #%s -->`, panel.HTMLID))
	return sliderPanel
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
	//lines = append(lines, indentationString[:indent]+fmt.Sprintf(`<div id="%s" class="%s %s">`, panel.TabBarHTMLID, classTabBar, visibility))
	tabBarPanel = &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			html.Attribute{Key: "id", Val: panel.TabBarHTMLID},
			html.Attribute{Key: "class", Val: fmt.Sprintf("%s %s", classTabBar, visibility)},
		},
	}
	// insert the buttons inside the tab bar panel
	for i, t := range panel.Tabs {
		//lines = append(lines, t.toButtonHTML(panel.TabBarHTMLID, indent2, (i == 0)))
		button := t.toButtonHTML(panel.TabBarHTMLID, (i == 0))
		tabBarPanel.AppendChild(button)
	}
	// close tab bar
	//lines = append(lines, indentationString[:indent]+fmt.Sprintf(`</div> <!-- end of #%s -->`, panel.TabBarHTMLID))
	// under tab bar panel
	underID := panel.TabBarHTMLID + dashUnderTabBar
	//lines = append(lines, indentationString[:indent]+fmt.Sprintf(`<div id="%s" class="%s">`, underID, classUnderTabBar))
	underTabBarPanel = &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			html.Attribute{Key: "id", Val: underID},
			html.Attribute{Key: "class", Val: classUnderTabBar},
		},
	}
	// add the markup panels to the under tab bar panel.
	for i, t := range panel.Tabs {
		//lines = append(lines, builder.toTabPanelHTML(t, indent2, (i == 0)))
		panel := builder.toTabPanelHTML(t, (i == 0))
		underTabBarPanel.AppendChild(panel)
	}
	// close under tab bar
	//lines = append(lines, indentationString[:indent]+fmt.Sprintf(`</div> <!-- end of #%s -->`, underID))
	return
}

func (builder *Builder) cookieCrumbs(serviceName string, locations []string) *html.Node {
	if len(locations) < 1 {
		return nil
	}
	//lines = append(lines, indentationString[:indent]+fmt.Sprintf(`<div class="%s"> <!-- cookie crumbs -->`, classPanelHeading))
	cc := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			html.Attribute{Key: "class", Val: classPanelHeading},
		},
	}
	l := len(locations)
	for i := 0; i < l; i++ {
		// lines = append(lines, fmt.Sprintf(`<h2 class="%s %s%s">%s</h2>`, classCookieCrumb, classCookieCrumbLevelPrefix, serviceName, locations[i]))
		h2 := &html.Node{
			Type:     html.ElementNode,
			DataAtom: atom.H2,
			Data:     "h2",
			Attr: []html.Attribute{
				html.Attribute{Key: "class", Val: fmt.Sprintf("%s %s%s", classCookieCrumb, classCookieCrumbLevelPrefix, serviceName)},
			},
		}
		textNode := &html.Node{
			Type: html.TextNode,
			Data: locations[i],
		}
		h2.AppendChild(textNode)
		cc.AppendChild(h2)
	}
	//lines = append(lines, strings.Join(innerLines, spaceString))
	//lines = append(lines, indentationString[:indent]+`</div> <!-- end of cookie crumbs -->`)
	return cc
}
