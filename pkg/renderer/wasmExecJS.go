package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createWASMExecJS(appPaths paths.ApplicationPathsI) error {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()
	oPath := filepath.Join(folderpaths.OutputRendererSite, fileNames.WasmExecJS)
	return appPaths.WriteFile(oPath, []byte(templates.WASMExecJS))
}
