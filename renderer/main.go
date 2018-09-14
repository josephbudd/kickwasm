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
		ApplicationGitPath                 string
		Host                               string
		Port                               uint
		Repos                              []string
		ImportRendererCall                 string
		ImportRendererViewTools            string
		ImportDomainImplementationsCalling string
	}{
		ApplicationGitPath:                 builder.ImportPath,
		Host:                               host,
		Port:                               port,
		Repos:                              builder.Repos,
		ImportRendererCall:                 folderpaths.ImportRendererCall,
		ImportRendererViewTools:            folderpaths.ImportRendererViewTools,
		ImportDomainImplementationsCalling: folderpaths.ImportDomainImplementationsCalling,
	}
	fname := "main.go"
	oPath := filepath.Join(folderpaths.OutputRenderer, fname)
	return templates.ProcessTemplate(fname, oPath, templates.MainGo, data, appPaths)
}
