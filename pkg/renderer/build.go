package renderer

import (
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/renderer/templates"
)

func createBuildSH(appPaths paths.ApplicationPathsI) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()
	folderNames := paths.GetFolderNames()

	var fname string
	var oPath string

	data := &struct {
		AppDotWASM  string
		HTTPDotYAML string
		MainDotGo   string
		PanelsDotGo string
		SiteFolder  string
		ServeDotGo  string
	}{
		AppDotWASM:  fileNames.AppDotWASM,
		HTTPDotYAML: fileNames.HTTPDotYAML,
		MainDotGo:   fileNames.MainDotGo,
		PanelsDotGo: fileNames.PanelsDotGo,
		SiteFolder:  folderNames.RendererSite,
		ServeDotGo:  fileNames.ServeDotGo,
	}

	// build.sh
	fname = fileNames.BuildDotSH
	oPath = filepath.Join(folderpaths.OutputRenderer, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.BuildDotSH, data, appPaths); err != nil {
		return
	}
	os.Chmod(oPath, 0744)

	// buildPack.sh
	fname = fileNames.BuildPackDotSH
	oPath = filepath.Join(folderpaths.OutputRenderer, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.BuildPackDotSH, data, appPaths); err != nil {
		return
	}
	os.Chmod(oPath, 0744)
	return
}
