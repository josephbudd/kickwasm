package project

import (
	"errors"
	"fmt"
	"strings"

	"github.com/josephbudd/kickwasm/pkg/cases"
)

type checker struct {
	serviceNames   []string
	panelNames     []string
	panelButtonIDs map[string][]string
	panelTabIDs    map[string][]string
}

func (ch *checker) isNewServiceName(name string) bool {
	for _, n := range ch.serviceNames {
		if n == name {
			return false
		}
	}
	ch.serviceNames = append(ch.serviceNames, name)
	return true
}

func (ch *checker) checkPanelID(panel *Panel) (err error) {
	ccName := cases.CamelCase(panel.Name)
	if ccName != panel.Name {
		err = errors.New("is not CamelCased")
		return
	}
	if !strings.HasSuffix(ccName, suffixPanel) {
		err = fmt.Errorf("should end with the suffix %q", suffixPanel)
		return
	}
	for _, n := range ch.panelNames {
		if n == ccName {
			err = errors.New("is not a new name")
			return
		}
	}
	panel.ID = ccName
	ch.panelNames = append(ch.panelNames, ccName)
	return
}

func (ch *checker) isNewButtonID(panel *Panel, button *Button) bool {
	_, found := ch.panelButtonIDs[panel.Name]
	if !found {
		ch.panelButtonIDs[panel.Name] = make([]string, 0, len(panel.Buttons))
	}
	for _, n := range ch.panelButtonIDs[panel.Name] {
		if n == button.ID {
			return false
		}
	}
	ch.panelButtonIDs[panel.Name] = append(ch.panelButtonIDs[panel.Name], button.ID)
	return true
}

func (ch *checker) isNewTabID(panel *Panel, tab *Tab) bool {
	_, found := ch.panelTabIDs[panel.Name]
	if !found {
		ch.panelTabIDs[panel.Name] = make([]string, 0, len(panel.Tabs))
	}
	for _, n := range ch.panelTabIDs[panel.Name] {
		if n == tab.ID {
			return false
		}
	}
	ch.panelTabIDs[panel.Name] = append(ch.panelTabIDs[panel.Name], tab.ID)
	return true
}

// BuildFromServices builds from services.
func (builder *Builder) BuildFromServices(services []*Service) (err error) {
	ch := &checker{
		panelNames:     make([]string, 0, 5),
		serviceNames:   make([]string, 0, 5),
		panelButtonIDs: make(map[string][]string),
		panelTabIDs:    make(map[string][]string),
	}
	for _, s := range services {
		if err = checkServiceValidity(s, 0, ch); err != nil {
			return
		}
	}
	builder.Services = services
	return
}

func checkServiceValidity(s *Service, level uint, ch *checker) (err error) {
	if err = isValidServiceName(s.Name); err != nil {
		return err
	}
	if !ch.isNewServiceName(s.Name) {
		return fmt.Errorf("the service name %q is not new", s.Name)
	}
	if s.Button == nil {
		return fmt.Errorf("the serviced name %q has no button", s.Name)
	}
	return checkButtonValidity(fmt.Sprintf("the %s service panel", s.Name), s.Button, 0, ch)
}

func isValidServiceName(name string) (err error) {
	if len(name) == 0 {
		return errors.New("a service is missing a name")
	}
	cs := cases.CamelCase(name)
	if cs != name {
		return fmt.Errorf("the service name %q is not camel cased", name)
	}
	return nil
}

func isValidButtonID(id string) (err error) {
	if len(strings.TrimSpace(id)) == 0 {
		return errors.New("missing a name")
	}
	ccID := cases.CamelCase(id)
	if id != ccID {
		return fmt.Errorf("button name %q is not CamelCased. It should be %q", id, ccID)
	}
	if !strings.HasSuffix(ccID, suffixButton) {
		return fmt.Errorf("button name %q should end with the suffix %q", id, suffixButton)
	}
	return nil
}

func isValidTabID(id string) (err error) {
	if len(strings.TrimSpace(id)) == 0 {
		return errors.New("missing a name")
	}
	ccID := cases.CamelCase(id)
	if id != ccID {
		return fmt.Errorf("tab name %q is not CamelCased. It should be %q", id, ccID)
	}
	if !strings.HasSuffix(ccID, suffixTab) {
		return fmt.Errorf("tab name %q should end with the suffix %q", id, suffixTab)
	}
	return nil
}

func checkButtonValidity(panelDesc string, b *Button, level uint, ch *checker) (err error) {
	if err = isValidButtonID(b.ID); err != nil {
		err = fmt.Errorf("%s %s", panelDesc, err.Error())
		return
	}
	if len(b.Heading) == 0 {
		b.Heading = b.Label
	}
	b.Label = strings.TrimSpace(b.Label)
	if len(b.Label) == 0 {
		err = fmt.Errorf("%s button is missing a label. The label is the button text", panelDesc)
		return
	}
	if len(b.Location) == 0 {
		b.Location = b.Heading
	}
	if len(b.Panels) == 0 {
		err = fmt.Errorf(`%s button with label %q has no panel files`, panelDesc, b.Label)
		return
	}
	for _, panel := range b.Panels {
		panel.Level = level + 1
		if err = checkButtonPanelValidity(panel, ch); err != nil {
			err = fmt.Errorf("%s: %s", panelDesc, err.Error())
			return
		}
	}
	return
}

func checkTabValidity(panelDesc string, tab *Tab, ch *checker) (err error) {
	if err = isValidTabID(tab.ID); err != nil {
		err = fmt.Errorf("%s %s", panelDesc, err.Error())
		return
	}
	if len(tab.Label) == 0 {
		err = fmt.Errorf("%s the tab named %q is missing a label", panelDesc, tab.ID)
		return
	}
	for _, panel := range tab.Panels {
		if err = checkTabPanelValidity(panel, ch); err != nil {
			err = fmt.Errorf("%s: %s", panelDesc, err.Error())
			return
		}
	}
	return
}

func checkTabPanelValidity(panel *Panel, ch *checker) (err error) {
	if err = ch.checkPanelID(panel); err != nil {
		err = fmt.Errorf("the tab panel name %q %s", panel.Name, err.Error())
	}
	return
}

func checkButtonPanelValidity(panel *Panel, ch *checker) (err error) {
	if err = ch.checkPanelID(panel); err != nil {
		err = fmt.Errorf("the button panel name %q %s", panel.Name, err.Error())
		return
	}
	for _, tab := range panel.Tabs {
		if err = checkTabValidity(fmt.Sprintf("the panel named %q", panel.Name), tab, ch); err != nil {
			return
		}
		if isnew := ch.isNewTabID(panel, tab); !isnew {
			err = fmt.Errorf("the tab panel named %q has more than one tab named %q", panel.Name, tab.ID)
			return
		}
	}
	for _, button := range panel.Buttons {
		if err = checkButtonValidity(fmt.Sprintf("the panel named %q", panel.Name), button, panel.Level, ch); err != nil {
			return
		}
		if isnew := ch.isNewButtonID(panel, button); !isnew {
			err = fmt.Errorf("the button panel named %q has more than one button named %q", panel.Name, button.ID)
			return
		}
	}
	return
}
