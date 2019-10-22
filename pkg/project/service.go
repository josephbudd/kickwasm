package project

import (
	"strings"
)

func cleanHomeMarkup(homeButton *Button) {
	for _, p := range homeButton.Panels {
		cleanHomePanelMarkup(homeButton, p)
	}
}

func cleanHomePanelMarkup(homeButton *Button, panel *Panel) {
	panel.Markup = strings.Replace(panel.Markup, "{home}", homeButton.ID, -1)
	for _, b := range panel.Buttons {
		for _, p := range b.Panels {
			cleanHomePanelMarkup(homeButton, p)
		}
	}
	for _, t := range panel.Tabs {
		for _, p := range t.Panels {
			cleanHomePanelMarkup(homeButton, p)
		}
	}
}
