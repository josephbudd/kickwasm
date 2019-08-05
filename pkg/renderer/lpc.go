package renderer

import (
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createLPC(appPaths paths.ApplicationPathsI, builder *project.Builder) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := appPaths.GetFileNames()

	data := struct {
		ApplicationGitPath      string
		ImportDomainLPC         string
		ImportDomainLPCMessage  string
		ImportRendererViewTools string
		NumberOfMarkupPanels    uint64
		LPCNames                []string
	}{
		ApplicationGitPath:      builder.ImportPath,
		ImportDomainLPC:         folderpaths.ImportDomainLPC,
		ImportDomainLPCMessage:  folderpaths.ImportDomainLPCMessage,
		ImportRendererViewTools: folderpaths.ImportRendererViewTools,
		NumberOfMarkupPanels:    builder.MarkupPanelCount,
		LPCNames:                make([]string, 0),
	}
	fname := fileNames.ChannelsDotGo
	oPath := filepath.Join(folderpaths.OutputRendererLPC, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.ChannelsGo, data, appPaths); err != nil {
		return
	}
	fname = fileNames.ClientDotGo
	oPath = filepath.Join(folderpaths.OutputRendererLPC, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.ClientGo, data, appPaths)
	return
}

// RebuildChannelsDotGo rebuilds renderer/lpc/channels.go
func RebuildChannelsDotGo(appPaths paths.ApplicationPathsI, importPath string, lpcNames []string) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := appPaths.GetFileNames()

	data := struct {
		ApplicationGitPath      string
		ImportDomainLPC         string
		ImportDomainLPCMessage  string
		ImportRendererViewTools string
		LPCNames                []string
		Inc                     func(int) int
	}{
		ApplicationGitPath:      importPath,
		ImportDomainLPC:         folderpaths.ImportDomainLPC,
		ImportDomainLPCMessage:  folderpaths.ImportDomainLPCMessage,
		ImportRendererViewTools: folderpaths.ImportRendererViewTools,
		LPCNames:                lpcNames,
		Inc:                     func(i int) int { return i + 1 },
	}
	fname := fileNames.ChannelsDotGo
	oPath := filepath.Join(folderpaths.OutputRendererLPC, fname)
	if err = os.Remove(oPath); err != nil && !os.IsNotExist(err) {
		return
	}
	err = templates.ProcessTemplate(fname, oPath, templates.ChannelsGo, data, appPaths)
	return
}