package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createViewToolsSliderGo(appPaths paths.ApplicationPathsI) error {
	folderpaths := appPaths.GetPaths()
	fname := "slider.go"
	oPath := filepath.Join(folderpaths.OutputRendererViewTools, fname)
	return appPaths.WriteFile(oPath, []byte(templates.ViewToolsSlider))
}
