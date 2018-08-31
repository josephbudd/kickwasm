package tap

import (
	"fmt"
	"strings"
)

// Button is a button in a button pad.
type Button struct {
	ID        string   `yaml:"name"`
	Label     string   `yaml:"label"`
	Heading   string   `yaml:"heading"`
	Location  string   `yaml:"cc,omitempty"`
	Panels    []*Panel `yaml:"panels,omitempty"` // sub panels within the tab panel
	Generated bool     `yaml:"-"`

	HTMLID           string `yaml:"-"` // the tab's id
	PanelHTMLID      string `yaml:"-"` // the id of the tab panel
	PanelInnerHTMLID string `yaml:"-"` // the inner div id of the tab panel
}

// GetHTMLID returns the tab's html id.
func (b *Button) GetHTMLID() string {
	return b.HTMLID
}

// toButtonHTML returns the button's button html
func (b *Button) toButtonHTML(idPrefix string, indent uint, backid string, colorLevel string) string {
	b.HTMLID = fmt.Sprintf("%s-%s", idPrefix, b.ID)
	return indentationString[:indent] + fmt.Sprintf(`<button id="%s" %s="%s" class="%s%s">%s</button>`, b.HTMLID, attributeBackID, backid, classPadButtonColorLevelPrefix, colorLevel, b.Label)
}

func (builder *Builder) toButtonSliderPanelsHTML(serviceName string, b *Button, locations []string, backid string, indent uint, seen bool, addLocations bool) string {
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
	lines := make([]string, 0, 5)
	indent2 := builder.incIndent(indent)
	for _, p := range b.Panels {
		p.HTMLID = b.HTMLID + dashString + p.ID
	}
	for i, p := range b.Panels {
		// panel group the first in each group must be tobeseen
		if len(p.Buttons) > 0 {
			lines = append(lines, builder.toSliderButtonPadPanelHTML(serviceName, p, locations, b.Heading, (seen && (i == 0)), addLocations))
			seen = false
			for _, b2 := range p.Buttons {
				lines = append(lines, builder.toButtonSliderPanelsHTML(serviceName, b2, locations, p.HTMLID, indent2, true, addLocations))
			}
		} else if len(p.Tabs) > 0 {
			lines = append(lines, builder.toSliderTabBarPanelHTML(serviceName, p, locations, b.Heading, seen, addLocations))
			seen = false
		} else {
			//line, innerid := builder.toSliderMarkupPanelHTML(serviceName, p, b, markupLocations, (seen && (i == 0)), addLocations)
			line, innerid := builder.toSliderMarkupPanelHTML(serviceName, p, b, locations, (seen && (i == 0)), addLocations)
			lines = append(lines, line)
			b.PanelInnerHTMLID = innerid
			seen = false
		}
	}
	return strings.Join(lines, newline)
}
