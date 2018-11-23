package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
	"github.com/josephbudd/kickwasm/pkg/tap"
)

type colorLevel struct {
	One   int
	Two   int
	Three int
	Four  int
	Five  int
}

type cssTemplateData struct {
	Seen             string
	UnSeen           string
	UnderTabBar      string
	Selected         string
	TabBar           string
	PanelHeading     string
	TabPanel         string
	InnerPanel       string
	PanelWithHeading string
	PanelWithTabBar  string

	SliderPanel       string
	SliderPanelInner  string
	SliderButtonPad   string
	UserContent       string
	ModalUserContent  string
	CloserUserContent string

	IDMaster           string
	IDHome             string
	IDHomePad          string
	IDSlider           string
	IDSliderBack       string
	IDSliderCollection string

	ClassBackColorLevelPrefix      string
	ClassPadColorLevelPrefix       string
	ClassPadButtonColorLevelPrefix string
	ClassCookieCrumbLevelPrefix    string
	ClassPanelHeadingLevelPrefix   string

	ColorLevels      []colorLevel
	ServiceNames     []string
	LastServiceIndex int
	Mod5             func(int) int
}

func createCSS(appPaths paths.ApplicationPathsI, builder *tap.Builder) error {
	n := 1
	if builder.Colors.LastColorLevel > 5 {
		n = int(builder.Colors.LastColorLevel) / 5
		if builder.Colors.LastColorLevel%5 > 0 {
			n++
		}
	}
	colorLevels := make([]colorLevel, n, n)
	for i := 0; i < int(n); i++ {
		colorLevels[i] = colorLevel{
			One:   1 + (i * 5),
			Two:   2 + (i * 5),
			Three: 3 + (i * 5),
			Four:  4 + (i * 5),
			Five:  5 + (i * 5),
		}
	}
	serviceNames := builder.GenerateServiceNames()
	data := &cssTemplateData{
		Seen:              builder.Classes.Seen,
		UnSeen:            builder.Classes.UnSeen,
		UnderTabBar:       builder.Classes.UnderTabBar,
		Selected:          builder.Classes.SelectedTab,
		TabBar:            builder.Classes.TabBar,
		PanelHeading:      builder.Classes.PanelHeading,
		TabPanel:          builder.Classes.TabPanel,
		InnerPanel:        builder.Classes.InnerPanel,
		PanelWithHeading:  builder.Classes.PanelWithHeading,
		PanelWithTabBar:   builder.Classes.PanelWithTabBar,
		SliderPanel:       builder.Classes.SliderPanel,
		SliderPanelInner:  builder.Classes.SliderPanelInner,
		SliderButtonPad:   builder.Classes.SliderButtonPad,
		UserContent:       builder.Classes.UserContent,
		ModalUserContent:  builder.Classes.ModalUserContent,
		CloserUserContent: builder.Classes.CloserUserContent,

		IDMaster:           builder.IDs.Master,
		IDHome:             builder.IDs.Home,
		IDHomePad:          builder.IDs.HomePad,
		IDSlider:           builder.IDs.Slider,
		IDSliderBack:       builder.IDs.SliderBack,
		IDSliderCollection: builder.IDs.SliderCollection,

		ClassBackColorLevelPrefix:      builder.Colors.ClassBackColorLevelPrefix,
		ClassPadColorLevelPrefix:       builder.Colors.ClassPadColorLevelPrefix,
		ClassPadButtonColorLevelPrefix: builder.Colors.ClassPadButtonColorLevelPrefix,
		ClassCookieCrumbLevelPrefix:    builder.Classes.CookieCrumbLevelPrefix,
		ClassPanelHeadingLevelPrefix:   builder.Classes.PanelHeadingLevelPrefix,

		ColorLevels:      colorLevels,
		ServiceNames:     serviceNames,
		LastServiceIndex: len(serviceNames) - 1,
		Mod5: func(i int) int {
			return i % 5
		},
	}
	folderpaths := appPaths.GetPaths()
	fname := "colors.css"
	oPath := filepath.Join(folderpaths.OutputRendererCSS, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.ColorsCSS, data, appPaths); err != nil {
		return err
	}
	fname = "main.css"
	oPath = filepath.Join(folderpaths.OutputRendererCSS, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.MainCSS, data, appPaths); err != nil {
		return err
	}
	return nil
}
