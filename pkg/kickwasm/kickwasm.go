package kickwasm

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickwasm/domain"
	"github.com/josephbudd/kickwasm/foldercp"
	"github.com/josephbudd/kickwasm/pkg/flagdata"
	"github.com/josephbudd/kickwasm/pkg/mainprocess"
	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer"
	"github.com/josephbudd/kickwasm/pkg/slurp"
)

// Do builds the source code and .kickwasm/ into the output folder.
// Returns the paths.ApplicationPathsI and the error.
func Do(pwd, outputFolder, yamlpath string, addLocations bool, host string, port uint, headTemplateFile string) (appPaths *paths.ApplicationPaths, err error) {
	sl := slurp.NewSlurper()
	builder, err := sl.Gulp(yamlpath)
	if err != nil {
		err = fmt.Errorf("Tried to slurp the YAML file(s) but counldn't, %s", err.Error())
		return
	}
	parts := strings.Split(builder.ImportPath, "/")
	appName := parts[len(parts)-1]
	appPaths = &paths.ApplicationPaths{}
	appPaths.Initialize(pwd, outputFolder, appName)
	if err = appPaths.MakeOutput(); err != nil {
		return
	}
	if err = create(appPaths, builder, addLocations, headTemplateFile); err != nil {
		return
	}
	folderPaths := appPaths.GetPaths()
	// create the .kickwasm/flags.yaml file
	flagsPath := filepath.Join(folderPaths.OutputDotKickwasm, "flags.yaml")
	yamlStartFileName := filepath.Base(yamlpath)
	if err = flagdata.SaveFlags(flagsPath, addLocations, yamlStartFileName); err != nil {
		return
	}
	// build the panel file paths
	appYAMLFilePath := filepath.Join(pwd, yamlpath)
	panelFilePaths := sl.GetPanelFilePaths()
	for i, p := range panelFilePaths {
		panelFilePaths[i] = filepath.Join(pwd, p)
	}
	foldercp.CopyYAML(appPaths, appYAMLFilePath, panelFilePaths)
	return
}

func create(appPaths paths.ApplicationPathsI, builder *project.Builder, addLocations bool, headTemplateFile string) (err error) {
	// get the framework name from the import path.
	if err = renderer.Create(appPaths, builder, addLocations, headTemplateFile); err != nil {
		return
	}
	if err = mainprocess.Create(appPaths, builder, headTemplateFile); err != nil {
		return
	}
	if err = domain.Create(appPaths, builder, headTemplateFile); err != nil {
		return
	}
	return
}
