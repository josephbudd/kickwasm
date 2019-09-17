package domain

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickwasm/pkg/cases"
	"github.com/josephbudd/kickwasm/pkg/domain/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

func createLPC(appPaths paths.ApplicationPathsI, data *templateData) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := appPaths.GetFileNames()

	fname := fileNames.PayloadDotGo
	oPath := filepath.Join(folderpaths.OutputDomainLPC, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.LPCPayloadGo, data, appPaths); err != nil {
		return
	}
	// log
	fname = fileNames.LogDotGo
	oPath = filepath.Join(folderpaths.OutputDomainLPCMessage, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.LPCLogGo, data, appPaths); err != nil {
		return
	}
	// init
	fname = fileNames.InitDotGo
	oPath = filepath.Join(folderpaths.OutputDomainLPCMessage, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.LPCInitGo, data, appPaths); err != nil {
		return
	}
	names := make([]string, 2, 2)
	parts := strings.Split(fileNames.LogDotGo, ".")
	names[0] = parts[0]
	parts = strings.Split(fileNames.InitDotGo, ".")
	names[1] = parts[0]
	err = RebuildDomainLPCInstructions(appPaths, data.ApplicationGitPath, names)
	return
}

// CreateCustomLPC creates a custom lpc go file.
func CreateCustomLPC(appPaths paths.ApplicationPathsI, lpcName string) (err error) {
	folderpaths := appPaths.GetPaths()
	// name is camel case.
	lpcName = cases.CamelCase(lpcName)
	fName := lpcName + ".go"
	data := &struct {
		LPCName string
	}{
		LPCName: lpcName,
	}
	oPath := filepath.Join(folderpaths.OutputDomainLPCMessage, fName)
	if err = templates.ProcessTemplate(fName, oPath, templates.MessageGo, data, appPaths); err != nil {
		return
	}
	return
}

// DeleteCustomLPC deletetes a custom lpc go file.
func DeleteCustomLPC(appPaths paths.ApplicationPathsI, lpcName string) (err error) {
	folderpaths := appPaths.GetPaths()
	// name is camel case.
	fname := cases.CamelCase(lpcName) + ".go"
	path := filepath.Join(folderpaths.OutputDomainLPCMessage, fname)
	err = os.Remove(path)
	return
}

// RebuildDomainLPCInstructions rebuilds the instructions.
func RebuildDomainLPCInstructions(appPaths paths.ApplicationPathsI, importGitPath string, lpcNames []string) (err error) {
	folderpaths := appPaths.GetPaths()
	fileNames := appPaths.GetFileNames()
	fixedLPCNames := make([]string, 0, len(lpcNames))
	// build the unwanted lpc file names.
	unwanted := make([]string, 3, 3)
	parts := strings.Split(fileNames.InstructionsDotTXT, ".")
	unwanted[0] = parts[0]
	parts = strings.Split(fileNames.LogDotGo, ".")
	unwanted[1] = parts[0]
	parts = strings.Split(fileNames.InitDotGo, ".")
	unwanted[2] = parts[0]
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
		ApplicationGitPath string
		LPCNames           []string
	}{
		ApplicationGitPath: importGitPath,
		LPCNames:           fixedLPCNames,
	}
	// instructions
	fname := fileNames.InstructionsDotTXT
	oPath := filepath.Join(folderpaths.OutputDomainLPCMessage, fname)
	if err = os.Remove(oPath); err != nil && !os.IsNotExist(err) {
		return
	}
	if err = templates.ProcessTemplate(fname, oPath, templates.LPCInstructions, data, appPaths); err != nil {
		return
	}
	return
}
