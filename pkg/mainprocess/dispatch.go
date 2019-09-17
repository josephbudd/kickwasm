package mainprocess

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickwasm/pkg/cases"
	"github.com/josephbudd/kickwasm/pkg/mainprocess/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

func createDispatch(appPaths paths.ApplicationPathsI, data *templateData) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := paths.GetFileNames()

	RebuildDispatchDotGo(appPaths, data.ApplicationGitPath, nil)

	var fname, oPath string

	// Init message.
	fname = fileNames.InitDotGo
	oPath = filepath.Join(folderpaths.OutputMainProcessLPCDispatch, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.DispatchInitGo, data, appPaths)

	// Log message.
	fname = fileNames.LogDotGo
	oPath = filepath.Join(folderpaths.OutputMainProcessLPCDispatch, fname)
	err = templates.ProcessTemplate(fname, oPath, templates.DispatchLogGo, data, appPaths)
	return
}

// CreateCustomDispatch creates a custom dispatch go file.
func CreateCustomDispatch(appPaths paths.ApplicationPathsI, importGitPath string, lpcName string) (err error) {
	folderpaths := appPaths.GetPaths()

	lpcName = cases.CamelCase(lpcName)
	fname := lpcName + ".go"
	oPath := filepath.Join(folderpaths.OutputMainProcessLPCDispatch, fname)
	data := struct {
		ApplicationGitPath     string
		ImportMainProcessLPC   string
		ImportDomainLPCMessage string
		ImportDomainStore      string
		MessageName            string
	}{
		ApplicationGitPath:     importGitPath,
		ImportMainProcessLPC:   folderpaths.ImportMainProcessLPC,
		ImportDomainLPCMessage: folderpaths.ImportDomainLPCMessage,
		ImportDomainStore:      folderpaths.ImportDomainStore,
		MessageName:            lpcName,
	}
	err = templates.ProcessTemplate(fname, oPath, templates.DispatchMessageGo, data, appPaths)
	return
}

// DeleteCustomDispatch deletes a custom dispatch go file.
func DeleteCustomDispatch(appPaths paths.ApplicationPathsI, lpcName string) (err error) {
	folderpaths := appPaths.GetPaths()

	fname := cases.CamelCase(lpcName) + ".go"
	path := filepath.Join(folderpaths.OutputMainProcessLPCDispatch, fname)
	err = os.Remove(path)
	return
}

// RebuildDispatchDotGo rebuilds dispatch.go.
func RebuildDispatchDotGo(appPaths paths.ApplicationPathsI, importGitPath string, lpcNames []string) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := appPaths.GetFileNames()
	fixedLPCNames := make([]string, 0, len(lpcNames))
	// build the unwanted lpc file names.
	unwanted := make([]string, 2, 2)
	parts := strings.Split(fileNames.InstructionsDotTXT, ".")
	unwanted[0] = parts[0]
	parts = strings.Split(fileNames.LogDotGo, ".")
	unwanted[1] = parts[0]
	// filter the lpc names
	for _, lpcName := range lpcNames {
		found := false
		for _, unw := range unwanted {
			if lpcName == unw {
				found = true
				break
			}
		}
		if !found {
			fixedLPCNames = append(fixedLPCNames, lpcName)
		}
	}
	// build the data
	data := struct {
		ApplicationGitPath        string
		ImportDomainDataLogLevels string
		ImportDomainLPCMessage    string
		ImportDomainStore         string
		ImportMainProcessLPC      string
		LPCNames                  []string
	}{
		ApplicationGitPath:        importGitPath,
		ImportDomainDataLogLevels: folderpaths.ImportDomainDataLogLevels,
		ImportDomainLPCMessage:    folderpaths.ImportDomainLPCMessage,
		ImportDomainStore:         folderpaths.ImportDomainStore,
		ImportMainProcessLPC:      folderpaths.ImportMainProcessLPC,
		LPCNames:                  fixedLPCNames,
	}
	// dispatch.go
	fname := fileNames.DispatchDotGo
	oPath := filepath.Join(folderpaths.OutputMainProcessLPCDispatch, fname)
	if err = os.Remove(oPath); err != nil && !os.IsNotExist(err) {
		return
	}
	if err = templates.ProcessTemplate(fname, oPath, templates.RebuildDispatchGo, data, appPaths); err != nil {
		return
	}
	// instructions
	fname = fileNames.InstructionsDotTXT
	oPath = filepath.Join(folderpaths.OutputMainProcessLPCDispatch, fname)
	if err = os.Remove(oPath); err != nil && !os.IsNotExist(err) {
		return
	}
	if err = templates.ProcessTemplate(fname, oPath, templates.DispatchInstructions, data, appPaths); err != nil {
		return
	}
	return
}
