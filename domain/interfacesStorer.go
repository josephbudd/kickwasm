package domain

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickwasm/domain/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

func createInterfacesStoreInterfaceGo(appPaths paths.ApplicationPathsI, data *templateData) (err error) {
	folderpaths := appPaths.GetPaths()
	rdata := &struct {
		Store              string
		LowerCamelCase     func(string) string
		ApplicationGitPath string
		ImportDomainTypes  string
	}{
		LowerCamelCase:     data.LowerCamelCase,
		ApplicationGitPath: data.ApplicationGitPath,
		ImportDomainTypes:  folderpaths.ImportDomainTypes,
	}
	for _, store := range data.Stores {
		fname := fmt.Sprintf("%s.go", data.LowerCamelCase(store))
		oPath := filepath.Join(folderpaths.OutputDomainInterfacesStorers, fname)
		rdata.Store = store
		if err = templates.ProcessTemplate(fname, oPath, templates.StorerGo, rdata, appPaths); err != nil {
			return
		}
	}
	return
}
