package mainprocess

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/mainprocess/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

func createCallServer(appPaths paths.ApplicationPathsI, data *templateData) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()
	fname := fileNames.LockedDotGo
	oPath := filepath.Join(folderpaths.OutputMainProcessCallServer, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.CallServerLockedGo, data, appPaths); err != nil {
		return
	}
	fname = fileNames.CallServerDotGo
	oPath = filepath.Join(folderpaths.OutputMainProcessCallServer, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.CallServerGo, data, appPaths); err != nil {
		return
	}
	fname = fileNames.RunDotGo
	oPath = filepath.Join(folderpaths.OutputMainProcessCallServer, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.CallServerRunGo, data, appPaths); err != nil {
		return
	}
	fname = fileNames.WebSocketDotGo
	oPath = filepath.Join(folderpaths.OutputMainProcessCallServer, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.CallServerWebsocketGo, data, appPaths); err != nil {
		return
	}
	return nil
}
