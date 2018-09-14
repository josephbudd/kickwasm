package domain

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickwasm/domain/templates"
	"github.com/josephbudd/kickwasm/paths"
)

// Create the files in the domain/implementations/store/boltstoring/ folder.
func createStoreBoltStoresGo(appPaths paths.ApplicationPathsI, data *templateData) error {
	folderpaths := appPaths.GetPaths()
	fname := "error.go"
	oPath := filepath.Join(folderpaths.OutputDomainImplementationsStoringBolt, fname)
	if err := templates.ProcessTemplate(fname, oPath, templates.ImplementationsStoringBoltErrorsGo, data, appPaths); err != nil {
		return err
	}
	data2 := struct {
		Repo               string
		ApplicationGitPath string
		ImportDomainTypes  string
		LowerCamelCase     func(string) string
	}{
		ApplicationGitPath: data.ApplicationGitPath,
		ImportDomainTypes:  data.ImportDomainTypes,
		LowerCamelCase:     data.LowerCamelCase,
	}
	for _, repo := range data.Repos {
		data2.Repo = repo
		fname := fmt.Sprintf("%sStore.go", repo)
		oPath := filepath.Join(folderpaths.OutputDomainImplementationsStoringBolt, fname)
		if err := templates.ProcessTemplate(fname, oPath, templates.ImplementationsStoringBoltStoringGo, data2, appPaths); err != nil {
			return err
		}
	}
	return nil
}
