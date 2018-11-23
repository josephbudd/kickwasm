package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createWASMExecJS(appPaths paths.ApplicationPathsI) error {
	folderpaths := appPaths.GetPaths()
	oPath := filepath.Join(folderpaths.OutputRenderer, "wasm_exec.js")
	return appPaths.WriteFile(oPath, []byte(templates.WASMExecJS))
}
