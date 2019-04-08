package project

import (
	"fmt"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Button is a button in a button pad.
type Button struct {
	ID       string   `yaml:"name"`
	Label    string   `yaml:"label"`
	Heading  string   `yaml:"heading"`
	Location string   `yaml:"cc,omitempty"`
	Panels   []*Panel `yaml:"panels,omitempty"` // sub panels within the tab panel

	HTMLID           string `yaml:"-"` // the tab's id
	PanelHTMLID      string `yaml:"-"` // the id of the tab panel
	PanelInnerHTMLID string `yaml:"-"` // the inner div id of the tab panel
}

// GetHTMLID returns the tab's html id.
func (b *Button) GetHTMLID() string {
	return b.HTMLID
}

// toButtonHTML returns the button's button html
func (b *Button) toButtonHTML(idPrefix string, backid string, colorLevel string) *html.Node {
	b.HTMLID = fmt.Sprintf("%s-%s", idPrefix, b.ID)
	button := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Button,
		Data:     "button",
		Attr: []html.Attribute{
			{Key: "id", Val: b.HTMLID},
			{Key: "class", Val: fmt.Sprintf("%s%s", classPadButtonColorLevelPrefix, colorLevel)},
			{Key: attributeBackID, Val: backid},
		},
	}
	textNode := &html.Node{
		Type: html.TextNode,
		Data: b.Label,
	}
	button.AppendChild(textNode)
	return button
}

func (builder *Builder) toButtonSliderPanelsHTML(serviceName string, b *Button, locations []string, backid string, seen bool, addLocations bool) []*html.Node {
	//markupLocations := locations[:]
	locations = append(locations, b.Location)
	// notes
	if len(b.Panels) == 1 {
		p := b.Panels[0]
		if len(p.Buttons) > 0 {
			// panel is a button bar
			if len(p.Note) == 0 {
				p.Note = fmt.Sprintf("This is the button pad displayed when the %s button is clicked.", b.Label)
			}
		} else if len(p.Tabs) > 0 {
			// panel is a tab bar
			if len(p.Note) == 0 {
				p.Note = fmt.Sprintf("This is the tab bar displayed when the %s button is clicked.", b.Label)
			}
		}
	}
	for _, p := range b.Panels {
		p.HTMLID = b.HTMLID + dashString + p.ID
	}
	panels := make([]*html.Node, 0, 20)
	for i, p := range b.Panels {
		// panel group the first in each group must be tobeseen
		if len(p.Buttons) > 0 {
			// this panel is a button pad.
			buttonPadPanel := builder.toSliderButtonPadPanelHTML(serviceName, p, locations, b.Heading, (seen && (i == 0)), addLocations)
			panels = append(panels, buttonPadPanel)
			seen = false
			for _, b2 := range p.Buttons {
				buttonPanels := builder.toButtonSliderPanelsHTML(serviceName, b2, locations, p.HTMLID, true, addLocations)
				panels = append(panels, buttonPanels...)
			}
		} else if len(p.Tabs) > 0 {
			// this panel is a tab bar
			tabBarPanel := builder.toSliderTabBarPanelHTML(serviceName, p, locations, b.Heading, seen, addLocations)
			panels = append(panels, tabBarPanel)
			seen = false
		} else {
			// this panel is a markup panel.
			markupPanel, innerid := builder.toSliderMarkupPanelHTML(serviceName, p, b, locations, (seen && (i == 0)), addLocations)
			b.PanelInnerHTMLID = innerid
			panels = append(panels, markupPanel)
			seen = false
		}
	}
	return panels
}
