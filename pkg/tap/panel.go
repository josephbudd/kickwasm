package tap

import (
	"fmt"
	"strings"
)

// JSData is the name and document id of javascript Controler or Presenter member.
type JSData struct {
	Name string `yaml:"name"`
	ID   string `yaml:"id"`
}

// JSHandler is the controler element's name, event and handler.
type JSHandler struct {
	Element  string `yaml:"element"`
	Event    string `yaml:"event"`
	Function string `yaml:"function"`
}

// Panel is a panel under a tab
type Panel struct {
	ID      string    `yaml:"id"`
	Name    string    `yaml:"name"`
	Tabs    []*Tab    `yaml:"tabs,omitempty"`
	Buttons []*Button `yaml:"buttons,omitempty"`
	Note    string    `yaml:"note"`

	Markup string `yaml:"markup,omitempty"`

	HTMLID            string `yaml:"HTMLID"`            // "-"
	TabBarHTMLID      string `yaml:"TabBarHTMLID"`      // "-"
	UnderTabBarHTMLID string `yaml:"UnderTabBarHTMLID"` // "-"

	Level    uint   `yaml:"-"`
	Template string `yaml:"-"`
}

// newPanel constructs a new Panel
func newPanel() *Panel {
	return &Panel{
		Tabs:    make([]*Tab, 0, 5),
		Buttons: make([]*Button, 0, 5),
	}
}

func (panel *Panel) markItUp(forwhat string, group []*Panel) string {
	lines := make([]string, 0, 5)
	var l int
	if group != nil {
		l = len(group)
	}
	// this panel comes from the user's json file.
	// add comments in the html and then comments renderered.
	lines = append(lines, emptyString)
	// html comments
	lines = append(lines, "<!--")
	lines = append(lines, emptyString)
	lines = append(lines, fmt.Sprintf("Panel name: %q", panel.Name))
	lines = append(lines, emptyString)
	lines = append(lines, fmt.Sprintf("Panel note: %s", panel.Note))
	lines = append(lines, emptyString)
	lines = append(lines, forwhat)
	lines = append(lines, emptyString)
	if l == 1 {
		lines = append(lines, "This panel is the only panel in it's panel group.")
	} else {
		lines = append(lines, fmt.Sprintf("This panel is just 1 in a group of %d panels.", l))
		if l == 2 {
			lines = append(lines, "Your other panel in this group is")
		} else {
			lines = append(lines, fmt.Sprintf("Your other %d panels in this group are", (l-1)))
		}
		if group != nil {
			for _, p := range group {
				if p != panel {
					lines = append(lines, emptyString)
					if len(p.Buttons) > 0 {
						lines = append(lines, fmt.Sprintf(`  * The button pad <div #%s`, p.innerID()))
						lines = append(lines, "  * But the panel is a button pad so you won't be adding any content there.")
					} else if len(p.Tabs) > 0 {
						lines = append(lines, fmt.Sprintf(`  * The tab bar <div #%s`, p.innerID()))
						lines = append(lines, "  * But the panel is a tab bar so you won't be adding any content there.")
					} else {
						lines = append(lines, fmt.Sprintf(`  * The content panel <div #%s`, p.innerID()))
					}
					if len(p.Buttons) == 0 && len(p.Tabs) == 0 {
						lines = append(lines, fmt.Sprint("  * Name: ", p.Name))
						lines = append(lines, fmt.Sprint("  * Note: ", p.Note))
					}
				}
			}
		}
	}
	lines = append(lines, emptyString)
	lines = append(lines, "-->")
	lines = append(lines, emptyString)
	// render markup
	lines = append(lines, panel.Markup)
	lines = append(lines, emptyString)
	return strings.Join(lines, newline)
}

func (panel *Panel) innerID() string {
	if len(panel.Buttons) > 0 {
		return panel.HTMLID + dashInnerString + dashButtonPadString
	}
	if len(panel.Tabs) > 0 {
		return strings.Replace(panel.HTMLID+dashTabBar, dashString, underscoreString, -1)
	}
	return panel.HTMLID + dashInnerString + dashContentString
}

func (panel *Panel) innerComment() string {
	lines := make([]string, 0, 5)
	if len(panel.Buttons) > 0 {
		lines = append(lines, fmt.Sprintf(`The button pad <div #%s`, panel.innerID()))
		lines = append(lines, "But the panel is a button pad so you won't be adding any content there.")
		return strings.Join(lines, newline)
	}
	if len(panel.Tabs) > 0 {
		lines = append(lines, fmt.Sprintf(`The tab bar <div #%s`, panel.innerID()))
		lines = append(lines, "But the panel is a tab bar so you won't be adding any content there.")
		return strings.Join(lines, newline)
	}
	lines = append(lines, fmt.Sprintf(`The content panel <div #%s`, panel.innerID()))
	return strings.Join(lines, newline)
}

func (panel *Panel) innerHTMLComment() string {
	lines := make([]string, 0, 5)
	if len(panel.Buttons) > 0 {
		lines = append(lines, fmt.Sprintf(`The button pad &lt;div #%s`, panel.innerID()))
		lines = append(lines, "<br/>But the panel is a button pad so you won't be adding any content there.")
		return strings.Join(lines, newline)
	}
	if len(panel.Tabs) > 0 {
		lines = append(lines, fmt.Sprintf(`The tab bar &lt;div #%s`, panel.innerID()))
		lines = append(lines, "<br/>But the panel is a tab bar so you won't be adding any content there.")
		return strings.Join(lines, newline)
	}
	lines = append(lines, fmt.Sprintf(`The content panel &lt;div #%s`, panel.innerID()))
	return strings.Join(lines, newline)
}

func (panel *Panel) tabSubPanelComment() string {
	return fmt.Sprintf(`  * The content panel <div #%s`, panel.HTMLID)
}

func (panel *Panel) tabSubPanelHTMLComment() string {
	return fmt.Sprintf(`  * The content panel &lt;div #%s`, panel.HTMLID)
}
