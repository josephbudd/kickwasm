package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/paths"
	"github.com/josephbudd/kickwasm/renderer/templates"
)

func createWASMExecJS(appPaths paths.ApplicationPathsI) error {
	folderpaths := appPaths.GetPaths()
	oPath := filepath.Join(folderpaths.OutputRendererWASM, "wasm_exec.js")
	return appPaths.WriteFile(oPath, []byte(templates.WASMExecJS))
}
