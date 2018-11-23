package tap

import (
	"strings"
)

// Service is an application service.
type Service struct {
	Name   string
	Button *Button
}

func (service *Service) cleanMarkup() {
	for _, p := range service.Button.Panels {
		service.cleanPanelMarkup(p)
	}
}

func (service *Service) cleanPanelMarkup(panel *Panel) {
	panel.Markup = strings.Replace(panel.Markup, "{service}", service.Name, -1)
	for _, b := range panel.Buttons {
		for _, p := range b.Panels {
			service.cleanPanelMarkup(p)
		}
	}
	for _, t := range panel.Tabs {
		for _, p := range t.Panels {
			service.cleanPanelMarkup(p)
		}
	}
}
