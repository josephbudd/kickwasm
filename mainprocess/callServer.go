package mainprocess

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/mainprocess/templates"
	"github.com/josephbudd/kickwasm/paths"
)

func createCallServer(appPaths paths.ApplicationPathsI, data *templateData) error {
	folderpaths := appPaths.GetPaths()
	fname := "locked.go"
	oPath := filepath.Join(folderpaths.OutputMainProcessTransportsCallServer, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.CallServerLockedGo, data, appPaths); err != nil {
		return err
	}
	fname = "callserver.go"
	oPath = filepath.Join(folderpaths.OutputMainProcessTransportsCallServer, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.CallServerGo, data, appPaths); err != nil {
		return err
	}
	fname = "run.go"
	oPath = filepath.Join(folderpaths.OutputMainProcessTransportsCallServer, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.CallServerRunGo, data, appPaths); err != nil {
		return err
	}
	fname = "websocket.go"
	oPath = filepath.Join(folderpaths.OutputMainProcessTransportsCallServer, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.CallServerWebsocketGo, data, appPaths); err != nil {
		return err
	}
	return nil
}
