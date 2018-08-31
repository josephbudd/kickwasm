package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/paths"
	"github.com/josephbudd/kickwasm/renderer/templates"
	"github.com/josephbudd/kickwasm/tap"
)

func createMainGo(host string, port uint, appPaths paths.ApplicationPathsI, builder *tap.Builder) error {
	folderpaths := appPaths.GetPaths()
	data := &struct {
		ApplicationGitPath               string
		Host                             string
		Port                             uint
		Repos                            []string
		ImportMainProcessTransportsCalls string
		ImportRendererWASMCall           string
		ImportRendererWASMViewTools      string
	}{
		ApplicationGitPath: builder.ImportPath,
		Host:               host,
		Port:               port,
		Repos:              builder.Repos,
		ImportMainProcessTransportsCalls: folderpaths.ImportMainProcessTransportsCalls,
		ImportRendererWASMCall:           folderpaths.ImportRendererWASMCall,
		ImportRendererWASMViewTools:      folderpaths.ImportRendererWASMViewTools,
	}
	fname := "main.go"
	oPath := filepath.Join(folderpaths.OutputRendererWASM, fname)
	return templates.ProcessTemplate(fname, oPath, templates.MainGo, data, appPaths)
}
