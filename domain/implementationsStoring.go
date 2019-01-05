package domain

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickwasm/domain/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

// Create the files in the domain/implementations/store/boltstoring/ folder.
func createStoreBoltStoresGo(appPaths paths.ApplicationPathsI, data *templateData) (err error) {
	folderpaths := appPaths.GetPaths()
	fname := "error.go"
	oPath := filepath.Join(folderpaths.OutputDomainImplementationsStoringBolt, fname)
	if err = templates.ProcessTemplate(fname, oPath, templates.ImplementationsStoringBoltErrorsGo, data, appPaths); err != nil {
		return
	}
	data2 := struct {
		Store              string
		ApplicationGitPath string
		ImportDomainTypes  string
		LowerCamelCase     func(string) string
	}{
		ApplicationGitPath: data.ApplicationGitPath,
		ImportDomainTypes:  data.ImportDomainTypes,
		LowerCamelCase:     data.LowerCamelCase,
	}
	for _, store := range data.Stores {
		data2.Store = store
		fname := fmt.Sprintf("%sStore.go", store)
		oPath := filepath.Join(folderpaths.OutputDomainImplementationsStoringBolt, fname)
		if err = templates.ProcessTemplate(fname, oPath, templates.ImplementationsStoringBoltStoringGo, data2, appPaths); err != nil {
			return
		}
	}
	return
}
