package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/paths"
	"github.com/josephbudd/kickwasm/renderer/templates"
)

func createViewToolsSliderGo(appPaths paths.ApplicationPathsI) error {
	folderpaths := appPaths.GetPaths()
	fname := "slider.go"
	oPath := filepath.Join(folderpaths.OutputRendererWASMViewTools, fname)
	return appPaths.WriteFile(oPath, []byte(templates.ViewToolsSlider))
}
