package tap

import (
	"fmt"
	"strings"
)

// Tab is a tab in a tab bar.
type Tab struct {
	ID        string   `yaml:"name"`
	Label     string   `yaml:"label"`
	Panels    []*Panel `yaml:"panels"`
	Generated bool     `yaml:"-"`

	HTMLID           string `yaml:"-"` // the tab's id
	PanelHTMLID      string `yaml:"-"` // the id of the tab panel
	PanelInnerHTMLID string `yaml:"-"` // the inner div id of the tab panel
}

// GetHTMLID returns the tab's html id.
func (t *Tab) GetHTMLID() string {
	return t.HTMLID
}

// toButtonHTML returns the tab's button html
func (t *Tab) toButtonHTML(idPrefix string, indent uint, selected bool) string {
	t.HTMLID = fmt.Sprintf("%s-%s", idPrefix, t.ID)
	var class string
	if selected {
		class = fmt.Sprintf("%s %s", classTab, classSelected)
	} else {
		class = fmt.Sprintf("%s %s", classTab, classUnSelected)
	}
	return indentationString[:indent] + fmt.Sprintf(`<button id="%s" class="%s">%s</button>`, t.HTMLID, class, t.Label)
}

// toTabPanelHTML returns the tab's panel html
func (builder *Builder) toTabPanelHTML(t *Tab, indent uint, seen bool) string {
	lines := make([]string, 0, 5)
	// the tab panel is bound to the tab button
	t.PanelHTMLID = t.HTMLID + suffixPanel
	var visibility string
	if seen {
		visibility = classSeen
	} else {
		visibility = classUnSeen
	}
	// open panel bound to tab, with header
	lines = append(lines, indentationString[:indent]+fmt.Sprintf(`<div id="%s" class="%s %s %s">`, t.PanelHTMLID, classTabPanel, classPanelWithHeading, visibility))
	indent2 := builder.incIndent(indent)
	// the tab panel has a header
	header := indentationString[:indent2] + fmt.Sprintf(`<h3 class="%s">%s</h3>`, classPanelHeading, t.Label)
	lines = append(lines, header)
	// under the header is the inner panel
	innerID := t.PanelHTMLID + dashInnerString
	t.PanelInnerHTMLID = innerID
	// open inner
	lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`<div id="%s" class="%s %s">`, innerID, classInnerPanel, classUserContent))
	// set the panel html id for each panel in this group.
	for _, p := range t.Panels {
		p.HTMLID = innerID + dashString + p.ID
	}
	// one or more panels.
	// if more than one then only one if visible at a time.
	// first is visible by default
	indent3 := builder.incIndent(indent2)
	l := len(t.Panels)
	var forwhat string
	if l == 1 {
		forwhat = "This panel is displayed when the %q tab button is clicked."
	} else {
		forwhat = fmt.Sprintf("This is one of a group of %d panels displayed when the %%q tab button is clicked.", l)
	}
	// each tab panel has markup.
	for i, p := range t.Panels {
		var visibility string
		if i == 0 {
			visibility = classSeen
		} else {
			visibility = classUnSeen
		}
		lines = append(lines, indentationString[:indent3]+fmt.Sprintf(`<div id="%s" class="%s %s">`, p.HTMLID, classSliderPanelInnerSibling, visibility))
		markup := p.markItUp(fmt.Sprintf(forwhat, t.Label), t.Panels)
		if !p.Generated {
			p.Template = markup
			lines = append(lines, fmt.Sprintf(`{{template "%s.tmpl"}}`, p.Name))
		} else {
			lines = append(lines, markup)
		}
		lines = append(lines, indentationString[:indent3]+fmt.Sprintf(`</div> <!-- end of #%s -->`, p.HTMLID))
	}
	// close inner
	lines = append(lines, indentationString[:indent2]+fmt.Sprintf(`</div> <!-- end of #%s -->`, innerID))
	// close panel bound to tab, with header
	lines = append(lines, indentationString[:indent]+fmt.Sprintf(`</div> <!-- end of #%s -->`, t.PanelHTMLID))
	return strings.Join(lines, newline)
}
