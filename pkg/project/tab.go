package project

import (
	"fmt"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Tab is a tab in a tab bar.
type Tab struct {
	ID     string   `yaml:"name"`
	Label  string   `yaml:"label"`
	Panels []*Panel `yaml:"panels"`

	HTMLID           string `yaml:"-"` // the tab's id
	PanelHTMLID      string `yaml:"-"` // the id of the tab panel
	PanelInnerHTMLID string `yaml:"-"` // the inner div id of the tab panel
}

// GetHTMLID returns the tab's html id.
func (t *Tab) GetHTMLID() string {
	return t.HTMLID
}

// toButtonHTML returns the tab's button html
func (t *Tab) toButtonHTML(idPrefix string, selected bool) (button *html.Node) {
	t.HTMLID = fmt.Sprintf("%s-%s", idPrefix, t.ID)
	var class string
	if selected {
		class = fmt.Sprintf("%s %s", classTab, classSelected)
	} else {
		class = fmt.Sprintf("%s %s", classTab, classUnSelected)
	}
	button = &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Button,
		Data:     "button",
		Attr: []html.Attribute{
			html.Attribute{Key: "id", Val: t.HTMLID},
			html.Attribute{Key: "class", Val: class},
		},
	}
	textNode := &html.Node{
		Type: html.TextNode,
		Data: t.Label,
	}
	button.AppendChild(textNode)
	return
}

// toTabPanelHTML returns the tab's panel html
func (builder *Builder) toTabPanelHTML(t *Tab, seen bool) (tabPanel *html.Node) {
	// the tab panel is bound to the tab button
	t.PanelHTMLID = t.HTMLID + suffixPanel
	var visibility string
	if seen {
		visibility = classSeen
	} else {
		visibility = classUnSeen
	}
	// the tab panel is bound to its tab
	// it wraps the inner which wraps the inner siblings.
	tabPanel = &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			html.Attribute{Key: "id", Val: t.PanelHTMLID},
			html.Attribute{Key: "class", Val: fmt.Sprintf("%s %s %s", classTabPanel, classPanelWithHeading, visibility)},
		},
	}
	// the tab panel has an h3
	h3 := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.H3,
		Data:     "h3",
		Attr: []html.Attribute{
			html.Attribute{Key: "class", Val: classPanelHeading},
		},
	}
	textNode := &html.Node{
		Type: html.TextNode,
		Data: t.Label,
	}
	h3.AppendChild(textNode)
	tabPanel.AppendChild(h3)
	// the tab panel has the inner panel under the header.
	// the inner panel wraps inner sibling panels which are the panels in the panel group.
	innerID := t.PanelHTMLID + dashInnerString
	t.PanelInnerHTMLID = innerID
	innerPanel := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			html.Attribute{Key: "id", Val: innerID},
			html.Attribute{Key: "class", Val: fmt.Sprintf("%s %s", classInnerPanel, classUserContent)},
		},
	}
	tabPanel.AppendChild(innerPanel)
	// the inner sibling panels.
	// one or more panels.
	// if more than one then only one if visible at a time.
	// first is visible by default
	l := len(t.Panels)
	var forwhat string
	if l == 1 {
		forwhat = "This panel is displayed when the %q tab button is clicked."
	} else {
		forwhat = fmt.Sprintf("This is one of a group of %d panels displayed when the %%q tab button is clicked.", l)
	}
	// each tab panel has markup.
	// each inner sibling panel is a markup panel.
	for i, p := range t.Panels {
		var visibility string
		if i == 0 {
			// by default the 1st panel is visible.
			visibility = classSeen
		} else {
			// by default the other panels are not visible.
			visibility = classUnSeen
		}
		p.HTMLID = innerID + dashString + p.ID
		innerSiblingMarkupPanel := &html.Node{
			Type:     html.ElementNode,
			DataAtom: atom.Div,
			Data:     "div",
			Attr: []html.Attribute{
				html.Attribute{Key: "id", Val: p.HTMLID},
				html.Attribute{Key: "class", Val: fmt.Sprintf("%s %s", classSliderPanelInnerSibling, visibility)},
			},
		}
		innerPanel.AppendChild(innerSiblingMarkupPanel)
		//markup := p.markItUp(fmt.Sprintf(forwhat, t.Label), t.Panels)
		// The markup for this panel is defined in the panel's yaml file.
		// The panel's markup will be in a template file.
		p.Template = p.markItUp(fmt.Sprintf(forwhat, t.Label), t.Panels)
		// Put template code linking to the template file in this markup panel.
		templateLink := &html.Node{
			Type: html.TextNode,
			Data: html.UnescapeString(fmt.Sprintf(`{{template "%s.tmpl"}}`, p.Name)),
		}
		innerSiblingMarkupPanel.AppendChild(templateLink)
	}
	// close inner
	// close panel bound to tab, with header
	return
}
