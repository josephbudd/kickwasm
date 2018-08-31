package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/paths"
	"github.com/josephbudd/kickwasm/renderer/templates"
	"github.com/josephbudd/kickwasm/tap"
)

func createCallerGo(appPaths paths.ApplicationPathsI, builder *tap.Builder) error {
	folderpaths := appPaths.GetPaths()
	fname := "call.go"
	oPath := filepath.Join(folderpaths.OutputRendererWASMCaller, fname)
	data := struct {
		ApplicationGitPath               string
		ImportMainProcessTransportsCalls string
		ImportRendererWASMViewTools      string
	}{
		ApplicationGitPath:               builder.ImportPath,
		ImportMainProcessTransportsCalls: folderpaths.ImportMainProcessTransportsCalls,
		ImportRendererWASMViewTools:      folderpaths.ImportRendererWASMViewTools,
	}
	return templates.ProcessTemplate(fname, oPath, templates.ClientGo, data, appPaths)
}
