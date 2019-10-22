package store

import (
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickwasm/pkg/paths"
)

// Manager manages lpc stores.
type Manager struct {
	storerFolder  string
	importGitPath string
	appPaths      paths.ApplicationPathsI
	instructions  string
	manifest      *Manifest
}

// NewManager constructs a new Message.
func NewManager(pwd string, appName string, importGitPath string) (manager *Manager, err error) {
	appPaths := &paths.ApplicationPaths{}
	appPaths.Initialize(pwd, "", appName)
	paths := appPaths.GetPaths()
	fileNames := appPaths.GetFileNames()
	instructionsParts := strings.Split(fileNames.InstructionsDotTXT, ".")
	manifestPath := filepath.Join(paths.OutputDotKickstore, fileNames.StoresDotYAML)
	var manifest *Manifest
	if manifest, err = NewManifest(manifestPath, appPaths.GetFMode()); err != nil {
		return
	}
	manager = &Manager{
		storerFolder:  paths.OutputDomainStoreStorer,
		importGitPath: importGitPath,
		appPaths:      appPaths,
		instructions:  instructionsParts[0],
		manifest:      manifest,
	}
	return
}
