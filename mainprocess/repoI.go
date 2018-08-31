package mainprocess

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickwasm/mainprocess/templates"
	"github.com/josephbudd/kickwasm/paths"
)

func createRepoIGo(appPaths paths.ApplicationPathsI, data *templateData) error {
	folderpaths := appPaths.GetPaths()
	rdata := &struct {
		Repo                         string
		LowerCamelCase               func(string) string
		ApplicationGitPath           string
		ImportMainProcessDataRecords string
	}{
		LowerCamelCase:               data.LowerCamelCase,
		ApplicationGitPath:           data.ApplicationGitPath,
		ImportMainProcessDataRecords: folderpaths.ImportMainProcessDataRecords,
	}
	for _, repo := range data.Repos {
		fname := fmt.Sprintf("%s.go", data.LowerCamelCase(repo))
		oPath := filepath.Join(folderpaths.OutputMainProcessBehaviorRepoI, fname)
		//oPath := filepath.Join(folderpaths.OutputMainProcessRepositoriesBolt, fname)
		rdata.Repo = repo
		if err := templates.ProcessTemplate(fname, oPath, templates.RepoIGo, rdata, appPaths); err != nil {
			return err
		}
	}
	return nil
}
