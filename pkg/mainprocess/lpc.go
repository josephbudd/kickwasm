package mainprocess

import (
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/pkg/mainprocess/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

// RebuildLPCChannels builds the lpc/channels.go file.
func RebuildLPCChannels(appPaths paths.ApplicationPathsI, importPath string, lpcNames []string) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()

	data := struct {
		ApplicationGitPath     string
		ImportDomainLPC        string
		ImportDomainLPCMessage string
		LPCNames               []string
		Inc                    func(int) int
	}{
		ApplicationGitPath:     importPath,
		ImportDomainLPC:        folderpaths.ImportDomainLPC,
		ImportDomainLPCMessage: folderpaths.ImportDomainLPCMessage,
		LPCNames:               lpcNames,
		Inc:                    func(i int) int { return i + 1 },
	}

	fname := fileNames.ChannelsDotGo
	oPath := filepath.Join(folderpaths.OutputMainProcessLPC, fname)
	if err = os.Remove(oPath); err != nil && !os.IsNotExist(err) {
		return
	}
	err = templates.ProcessTemplate(fname, oPath, templates.LPCChannelsGo, data, appPaths)
	return
}
func createLPC(appPaths paths.ApplicationPathsI, data *templateData) (err error) {

	if err = RebuildLPCChannels(appPaths, data.ApplicationGitPath, make([]string, 0)); err != nil {
		return
	}

	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()

	var fname string
	var oPath string
	fname = fileNames.LockedDotGo
	oPath = filepath.Join(folderpaths.OutputMainProcessLPC, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.LPCLockedGo, data, appPaths); err != nil {
		return
	}
	fname = fileNames.ServerDotGo
	oPath = filepath.Join(folderpaths.OutputMainProcessLPC, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.LPCServerGo, data, appPaths); err != nil {
		return
	}
	fname = fileNames.RunDotGo
	oPath = filepath.Join(folderpaths.OutputMainProcessLPC, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.LPCRunGo, data, appPaths); err != nil {
		return
	}
	fname = fileNames.WebSocketDotGo
	oPath = filepath.Join(folderpaths.OutputMainProcessLPC, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.LPCWebsocketGo, data, appPaths); err != nil {
		return
	}
	return nil
}
