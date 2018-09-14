package domain

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickwasm/domain/templates"
	"github.com/josephbudd/kickwasm/paths"
)

func createInterfacesStoreInterfaceGo(appPaths paths.ApplicationPathsI, data *templateData) error {
	folderpaths := appPaths.GetPaths()
	rdata := &struct {
		Repo               string
		LowerCamelCase     func(string) string
		ApplicationGitPath string
		ImportDomainTypes  string
	}{
		LowerCamelCase:     data.LowerCamelCase,
		ApplicationGitPath: data.ApplicationGitPath,
		ImportDomainTypes:  folderpaths.ImportDomainTypes,
	}
	for _, repo := range data.Repos {
		fname := fmt.Sprintf("%s.go", data.LowerCamelCase(repo))
		oPath := filepath.Join(folderpaths.OutputDomainInterfacesStorers, fname)
		rdata.Repo = repo
		if err := templates.ProcessTemplate(fname, oPath, templates.StorerGo, rdata, appPaths); err != nil {
			return err
		}
	}
	return nil
}
