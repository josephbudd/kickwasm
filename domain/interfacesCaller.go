package domain

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/domain/templates"
	"github.com/josephbudd/kickwasm/pkg/paths"
)

func createInterfacesCallInterfaceGo(appPaths paths.ApplicationPathsI) error {
	folderpaths := appPaths.GetPaths()
	fname := "callInterface.go"
	oPath := filepath.Join(folderpaths.OutputDomainInterfacesCallers, fname)
	return templates.ProcessTemplate(fname, oPath, templates.CallInterfaceGo, nil, appPaths)
}
