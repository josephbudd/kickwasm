package mainprocess

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/mainprocess/templates"
	"github.com/josephbudd/kickwasm/paths"
)

func createCallsGo(appPaths paths.ApplicationPathsI, data *templateData) error {
	folderpaths := appPaths.GetPaths()
	fname := "calls.go"
	oPath := filepath.Join(folderpaths.OutputMainProcessTransportsCalls, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.CallsGo, data, appPaths); err != nil {
		return err
	}
	fname = "log.go"
	oPath = filepath.Join(folderpaths.OutputMainProcessTransportsCalls, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.CallsLogGo, data, appPaths); err != nil {
		return err
	}
	fname = "lpc.go"
	oPath = filepath.Join(folderpaths.OutputMainProcessTransportsCalls, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.CallsLPCGo, data, appPaths); err != nil {
		return err
	}
	fname = "example_go.txt"
	oPath = filepath.Join(folderpaths.OutputMainProcessTransportsCalls, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.CallsExampleGo, data, appPaths); err != nil {
		return err
	}
	return nil
}
