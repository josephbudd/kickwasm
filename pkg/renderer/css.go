package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
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
	TabPanelGroup    string
	PanelWithHeading string
	PanelWithTabBar  string

	SliderPanel       string
	SliderPanelPad    string
	SliderButtonPad   string
	UserContent       string
	ModalUserContent  string
	CloserUserContent string
	UserMarkup        string

	VScroll  string
	HVScroll string

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
	HomeNames     []string
	LastHomeIndex int
	Mod5             func(int) int
}

func createCSS(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
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
	homeNames := builder.GenerateHomeButtonNames()
	data := &cssTemplateData{
		Seen:              builder.Classes.Seen,
		UnSeen:            builder.Classes.UnSeen,
		UnderTabBar:       builder.Classes.UnderTabBar,
		Selected:          builder.Classes.SelectedTab,
		TabBar:            builder.Classes.TabBar,
		PanelHeading:      builder.Classes.PanelHeading,
		TabPanel:          builder.Classes.TabPanel,
		TabPanelGroup:     builder.Classes.TabPanelGroup,
		PanelWithHeading:  builder.Classes.PanelWithHeading,
		PanelWithTabBar:   builder.Classes.PanelWithTabBar,
		SliderPanel:       builder.Classes.SliderPanel,
		SliderPanelPad:    builder.Classes.SliderPanelPad,
		SliderButtonPad:   builder.Classes.SliderButtonPad,
		UserContent:       builder.Classes.UserContent,
		ModalUserContent:  builder.Classes.ModalUserContent,
		CloserUserContent: builder.Classes.CloserUserContent,
		VScroll:           builder.Classes.VScroll,
		HVScroll:          builder.Classes.HVScroll,
		UserMarkup:        builder.Classes.UserMarkup,

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
		HomeNames:     homeNames,
		LastHomeIndex: len(homeNames) - 1,
		Mod5: func(i int) int {
			return i % 5
		},
	}
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()
	fname := fileNames.ColorsDotCSS
	oPath := filepath.Join(folderpaths.OutputRendererCSS, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.ColorsCSS, data, appPaths); err != nil {
		return
	}
	fname = fileNames.MainDotCSS
	oPath = filepath.Join(folderpaths.OutputRendererCSS, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.MainCSS, data, appPaths); err != nil {
		return
	}
	fname = fileNames.UserContentDotCSS
	oPath = filepath.Join(folderpaths.OutputRendererMyCSS, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.MyCSS, data, appPaths); err != nil {
		return
	}
	return
}
