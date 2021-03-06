package project

import (
	"fmt"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Tab is a tab in a tab bar.
type Tab struct {
	ID      string   `yaml:"name"`
	Label   string   `yaml:"label"`
	Heading string   `yaml:"heading"`
	Panels  []*Panel `yaml:"panels"`

	HTMLID           string `yaml:"-"` // the tab's id.
	PanelHTMLID      string `yaml:"-"` // the id of the tab panel.
	PanelH3HTMLID    string `yaml:"-"` // the id of the tab panel's h3.
	PanelInnerHTMLID string `yaml:"-"` // the inner div id of the tab panel.

	Spawn bool `yaml:"spawn"`
}

// GetHTMLID returns the tab's html id.
func (t *Tab) GetHTMLID() string {
	return t.HTMLID
}

// toButtonHTML returns the tab's button html
// a spawnable tab is a template for other tabs.
// a spawnable tab is unseen and it's id contains a unique id replace pattern.
func (t *Tab) toButtonHTML(idPrefix string, selected bool) (button *html.Node) {
	var visibilityClass string
	attributes := make([]html.Attribute, 0, 10)
	if t.Spawn {
		selected = false
		visibilityClass = classUnSeen
		t.HTMLID = fmt.Sprintf("%s-%s-%s", idPrefix, t.ID, SpawnIDReplacePattern)
	} else {
		visibilityClass = classSeen
		t.HTMLID = fmt.Sprintf("%s-%s", idPrefix, t.ID)
	}
	// id attribute
	attributes = append(attributes, html.Attribute{Key: "id", Val: t.HTMLID})
	var selectionClass string
	if selected {
		selectionClass = classSelected
	} else {
		selectionClass = classUnSelected
	}
	// class attribute
	class := fmt.Sprintf("%s %s %s", classTab, selectionClass, visibilityClass)
	attributes = append(attributes, html.Attribute{Key: "class", Val: class})
	// spawnable attribute
	if t.Spawn {
		attributes = append(attributes, html.Attribute{Key: attributeSpawnable, Val: attributeSpawnable})
	}
	button = &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Button,
		Data:     "button",
		Attr:     attributes,
	}
	// button label
	if !t.Spawn {
		textNode := &html.Node{
			Type: html.TextNode,
			Data: t.Label,
		}
		button.AppendChild(textNode)
	}
	return
}

// toTabPanelHTML returns the tab's panel html
// a spawnable tab panel is a template for other tab panels.
// a spawnable tab panel is unseen and it's id contains a unique id replace pattern.
func (builder *Builder) toTabPanelHTML(t *Tab, seen bool) (tabPanel *html.Node) {
	var attributes []html.Attribute
	// the tab panel is bound to the tab button
	t.PanelHTMLID = t.HTMLID + suffixPanel
	var visibilityClass string
	if t.Spawn {
		seen = false
	}
	if seen {
		visibilityClass = classSeen
	} else {
		visibilityClass = classUnSeen
	}
	// the tab panel is bound to its tab
	// it wraps the inner which wraps the inner siblings.
	attributes = make([]html.Attribute, 0, 10)
	attributes = append(attributes, html.Attribute{Key: "id", Val: t.PanelHTMLID})
	attributes = append(attributes, html.Attribute{Key: "class", Val: fmt.Sprintf("%s %s %s", classTabPanel, classPanelWithHeading, visibilityClass)})
	tabPanel = &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr:     attributes,
	}
	// the tab panel has an h3
	t.PanelH3HTMLID = t.PanelHTMLID + dashH3String
	attributes = make([]html.Attribute, 0, 10)
	attributes = append(attributes, html.Attribute{Key: "id", Val: t.PanelH3HTMLID})
	attributes = append(attributes, html.Attribute{Key: "class", Val: classPanelHeading})
	h3 := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.H3,
		Data:     "h3",
		Attr:     attributes,
	}
	if !t.Spawn {
		textNode := &html.Node{
			Type: html.TextNode,
			Data: t.Heading,
		}
		h3.AppendChild(textNode)
	}
	tabPanel.AppendChild(h3)
	// Under the h3 is the group of markup panels.
	// How the group works.
	//  The inner panel wraps the group of content panels.
	//  Each content panel wraps a single markup panel.
	innerID := t.PanelHTMLID + DashInnerString
	t.PanelInnerHTMLID = innerID
	attributes = make([]html.Attribute, 0, 10)
	attributes = append(attributes, html.Attribute{Key: "id", Val: innerID})
	attributes = append(attributes, html.Attribute{Key: "class", Val: classTabPanelGroup})
	innerPanel := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr:     attributes,
	}
	tabPanel.AppendChild(innerPanel)
	// Panel group of one or more user content panels.
	// if more than one then only one if visible at a time.
	// first is visible by default
	// tabs only have markup.
	// each panel in the group is a content panel wrapping a markup panel.
	for i, p := range t.Panels {
		p.H3ID = t.PanelH3HTMLID
		builder.MarkupPanelCount++
		var visibility string
		if i == 0 {
			// by default the 1st panel is visible.
			visibility = classSeen
		} else {
			// by default the other panels are not visible.
			visibility = classUnSeen
		}
		p.HTMLID = innerID + dashString + p.ID
		var scroll string
		if p.HVScroll {
			scroll = classHVScroll
		} else {
			scroll = classVScroll
		}
		attributes = make([]html.Attribute, 0, 10)
		attributes = append(attributes, html.Attribute{Key: "id", Val: p.HTMLID})
		attributes = append(attributes, html.Attribute{Key: "class", Val: fmt.Sprintf("%s %s %s", classUserContent, visibility, scroll)})
		userContentPanel := &html.Node{
			Type:     html.ElementNode,
			DataAtom: atom.Div,
			Data:     "div",
			Attr:     attributes,
		}
		innerPanel.AppendChild(userContentPanel)
		attributes = make([]html.Attribute, 0, 1)
		if !p.HVScroll {
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
		// The markup for this panel is defined in the panel's yaml file.
		// Put template code linking to the template file in this markup panel.
		templateLink := &html.Node{
			Type: html.TextNode,
			Data: html.UnescapeString(fmt.Sprintf(`{{template "%s.tmpl"}}`, p.Name)),
		}
		markupPanel.AppendChild(templateLink)
	}
	// Now that the ids for all of the panels are created make the templates.
	l := len(t.Panels)
	for _, p := range t.Panels {
		var forwhat string
		if l == 1 {
			forwhat = "This panel is displayed when the %q tab button is clicked."
		} else {
			forwhat = fmt.Sprintf("This is one of a group of %d panels displayed when the %%q tab button is clicked.", l)
		}
		// The panel's markup will be in a template file.
		p.Template = p.markItUp(fmt.Sprintf(forwhat, t.Label), t.Panels)
	}
	return
}
