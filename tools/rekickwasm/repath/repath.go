package repath

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/tools/rekickwasm/ftools"
	"github.com/josephbudd/kickwasm/tools/rekickwasm/statements"
)

// folder names
const (
	EditFolder       = "edit"
	RekickWasmFolder = "rekickwasm"
	BackupFolder     = ".backup"
	ChangesFolder    = ".changes"
	RefactorFolder   = ".refactored"
	initializedFile  = ".initialized"
	noFolder         = ""
)

// RePaths are the folder paths
type RePaths struct {
	totallyInitialized                                      bool
	originalFolderPath                                      string
	DotKickwasm                                             string
	RekickWasm                                              string
	RekickWasmInitialized                                   string
	RekickWasmEdit, RekickWasmEditFlags, RekickWasmEditYAML string
	Original, Backup, Changes, Refactor                     *paths.ApplicationPaths
}

// NewRePaths constructs a new RePath.
// Param originalFolderPath is the path of the original source code.
func NewRePaths(originalFolderPath string) (rp *RePaths, err error) {
	folderNames := paths.GetFolderNames()
	fileNames := paths.GetFileNames()
	rp = &RePaths{}
	rp.originalFolderPath = originalFolderPath
	rp.DotKickwasm = filepath.Join(originalFolderPath, folderNames.DotKickwasm)
	rp.RekickWasm = filepath.Join(originalFolderPath, RekickWasmFolder)
	rp.RekickWasmInitialized = filepath.Join(rp.RekickWasm, initializedFile)
	rp.RekickWasmEdit = filepath.Join(rp.RekickWasm, EditFolder)
	rp.RekickWasmEditFlags = filepath.Join(rp.RekickWasmEdit, fileNames.FlagDotYAML)
	rp.RekickWasmEditYAML = filepath.Join(rp.RekickWasmEdit, folderNames.YAML)
	return
}

func (rp *RePaths) totallyInitialize() {
	if !rp.totallyInitialized {
		rp.totallyInitialized = true
		// Finish initializing rp.
		rp.Original = &paths.ApplicationPaths{}
		rp.Original.Initialize(rp.originalFolderPath, noFolder, noFolder)
		rp.Backup = &paths.ApplicationPaths{}
		rp.Backup.Initialize(rp.originalFolderPath, RekickWasmFolder, BackupFolder)
		rp.Changes = &paths.ApplicationPaths{}
		rp.Changes.Initialize(rp.originalFolderPath, RekickWasmFolder, ChangesFolder)
		rp.Refactor = &paths.ApplicationPaths{}
		rp.Refactor.Initialize(rp.originalFolderPath, RekickWasmFolder, RefactorFolder)
	}
}

// Initialized determines if rekickwasm has been initialized into the working folder.
func (rp *RePaths) Initialized() (initialized bool) {
	if _, err := os.Stat(rp.RekickWasmInitialized); err == nil {
		initialized = true
		// Finish the constructor.
		rp.totallyInitialize()
	}
	return
}

// InitializeWorkingDirectory creates the folders.
func (rp *RePaths) InitializeWorkingDirectory() (err error) {
	if rp.Initialized() {
		err = errors.New(statements.AlreadyInit)
		return
	}
	// Finish the constructor.
	rp.totallyInitialize()

	// Initialize the working folder for rekickwasm.
	originalPaths := rp.Original.GetPaths()
	backupPaths := rp.Backup.GetPaths()
	changesPaths := rp.Changes.GetPaths()
	mergePaths := rp.Refactor.GetPaths()
	var fc *ftools.DCopy
	// backup the original source to the .rekickwasm/.backup/ folder. Do not copy the rekickwasm folder to the backup.
	fc, err = ftools.NewDCopy(
		originalPaths.Output,
		backupPaths.Output,
		true, true,
		[]string{RekickWasmFolder},
	)
	if err != nil {
		msg := fmt.Sprintf(statements.ErrorInitFormat, err.Error())
		err = errors.New(msg)
		return
	}
	if err = fc.Copy(); err != nil {
		msg := fmt.Sprintf(statements.ErrorInitFormat, err.Error())
		err = errors.New(msg)
		return
	}
	// remove the .rekickwasm/.changes/ folder if it exists.
	if err = os.RemoveAll(changesPaths.Output); err != nil {
		msg := fmt.Sprintf(statements.ErrorInitFormat, err.Error())
		err = errors.New(msg)
		return
	}
	// remove the .rekickwasm/refactor/ folder it if exists.
	if err = os.RemoveAll(mergePaths.Output); err != nil {
		msg := fmt.Sprintf(statements.ErrorInitFormat, err.Error())
		err = errors.New(msg)
		return
	}
	// copy original .kickwasm/ to ./rekickwasm/edit/
	// the user will edit the yaml files in ./rekickwasm/edit/
	fc, err = ftools.NewDCopy(
		backupPaths.OutputDotKickwasm,
		rp.RekickWasmEdit,
		true, true,
		nil,
	)
	if err != nil {
		msg := fmt.Sprintf(statements.ErrorInitFormat, err.Error())
		err = errors.New(msg)
		return
	}
	if err = fc.Copy(); err != nil {
		msg := fmt.Sprintf(statements.ErrorInitFormat, err.Error())
		err = errors.New(msg)
		return
	}

	// Write the init file.
	var f *os.File
	if f, err = os.Create(rp.RekickWasmInitialized); err != nil {
		return
	}
	err = f.Close()
	return
}

// InitializeChanges empties the changes folder and copies the user edited yaml to the changes folder.
// Note:
// After this copy the changes source code can be built.
// The changes source code is the source code showing the changes.
// The changes must be merged with the original source code in the refactor folder.
func (rp *RePaths) InitializeChanges() (err error) {
	// empty the changes folder by removing it and then re-creating it.
	changesPaths := rp.Changes.GetPaths()
	if err = os.RemoveAll(changesPaths.Output); err != nil {
		return
	}
	if err = rp.Changes.MakeOutput(); err != nil {
		return
	}
	// copy the user edited yaml to the changes .kickwasm folder.
	var fc *ftools.DCopy
	fc, err = ftools.NewDCopy(
		rp.RekickWasmEdit,
		changesPaths.OutputDotKickwasm,
		true, true,
		nil,
	)
	if err != nil {
		return
	}
	err = fc.Copy()
	return
}

// InitializeRefactor copies the backup to the refactor folder
// After this the changes source code can be merged with into the refactor folder.
func (rp *RePaths) InitializeRefactor() (err error) {
	backupPaths := rp.Backup.GetPaths()
	mergePaths := rp.Refactor.GetPaths()
	// Begin with a fresh refactor folder and copy backup to it.
	var fc *ftools.DCopy
	fc, err = ftools.NewDCopy(
		backupPaths.Output,
		mergePaths.Output,
		true, true,
		[]string{RekickWasmFolder},
	)
	if err != nil {
		return
	}
	err = fc.Copy()
	return
}

// ImportRefactor imports the refactored refactor/ folder into the original folder.
func (rp *RePaths) ImportRefactor() (err error) {
	return copyFolder(rp.Refactor, rp.Original)
}

// RestoreOriginal imports the backup folder into the original folder.
func (rp *RePaths) RestoreOriginal() (err error) {
	return copyFolder(rp.Backup, rp.Original)
}

// RestoreYAML imports the backup .kickwasm/ folder into the refactor .kickwasm/edit/ folder.
func (rp *RePaths) RestoreYAML() (err error) {
	backupPaths := rp.Backup.GetPaths()
	var dcopy *ftools.DCopy
	// .kickwasm/
	dcopy, err = ftools.NewDCopy(
		backupPaths.OutputDotKickwasm,
		rp.RekickWasmEdit,
		true, true,
		nil,
	)
	if err != nil {
		return err
	}
	err = dcopy.Copy()
	return
}

// copyFolder copies the folder src to dst.
func copyFolder(srcApplicationPaths, dstApplicationPaths *paths.ApplicationPaths) (err error) {
	/*
		./mainprocess/panelMap.go

		./rendererprocess/css/
		./rendererprocess/framework/panels.go
		./rendererprocess/framework/lpc/client.go
		./rendererprocess/panels/
		./rendererprocess/spawnPanels/
		./rendererprocess/framework/viewtools/
		./rendererprocess/framework/proofs/
		./rendererprocess/application/Application.go

		./site/templates/
		./site/spawnTemplates/

		./.kickwasm/
	*/

	fileNames := paths.GetFileNames()
	srcPaths := srcApplicationPaths.GetPaths()
	dstPaths := dstApplicationPaths.GetPaths()
	var src, dst string
	// Files to be copied.

	// File in the ./mainprocess/ folder.
	// ./mainprocess/panelMap.go
	src = filepath.Join(srcPaths.OutputMainProcess, fileNames.PanelMapDotGo)
	dst = filepath.Join(dstPaths.OutputMainProcess, fileNames.PanelMapDotGo)
	if err = ftools.CopyFile(src, dst); err != nil {
		return
	}

	// Files and folders in the ./rendererprocess/ folder.
	// ./rendererprocess/framework/panels.go
	src = filepath.Join(srcPaths.OutputRendererFramework, fileNames.PanelsDotGo)
	dst = filepath.Join(dstPaths.OutputRendererFramework, fileNames.PanelsDotGo)
	if err = ftools.CopyFile(src, dst); err != nil {
		return
	}
	// ./rendererprocess/framework/lpc/client.go
	src = filepath.Join(srcPaths.OutputRendererLPC, fileNames.ClientDotGo)
	dst = filepath.Join(dstPaths.OutputRendererLPC, fileNames.ClientDotGo)
	if err = ftools.CopyFile(src, dst); err != nil {
		return
	}
	// ./rendererprocess/panels/
	if err = _copyFolder(srcPaths.OutputRendererPanels, dstPaths.OutputRendererPanels); err != nil {
		return
	}
	// ./rendererprocess/spawnPanels/
	// The renderer/spawnPanels folder might not exist.
	if err = _copyFolder(srcPaths.OutputRendererSpawns, dstPaths.OutputRendererSpawns); err != nil {
		return
	}
	// ./rendererprocess/framework/viewtools/
	if err = _copyFolder(srcPaths.OutputRendererViewTools, dstPaths.OutputRendererViewTools); err != nil {
		return
	}
	// ./rendererprocess/framework/proofs/
	if err = _copyFolder(srcPaths.OutputRendererProofs, dstPaths.OutputRendererProofs); err != nil {
		return
	}
	// ./rendererprocess/application/Application.go
	src = filepath.Join(srcPaths.OutputRendererApplication, fileNames.ApplicationDotGo)
	dst = filepath.Join(dstPaths.OutputRendererApplication, fileNames.ApplicationDotGo)
	if err = ftools.CopyFile(src, dst); err != nil {
		return
	}

	// Files and folders in the .site/ folder.
	// ./site/css/
	if err = _copyFolder(srcPaths.OutputRendererCSS, dstPaths.OutputRendererCSS); err != nil {
		return
	}
	// ./site/templates/
	if err = _copyFolder(srcPaths.OutputRendererTemplates, dstPaths.OutputRendererTemplates); err != nil {
		return
	}
	// ./site/spawnTemplates/
	// The site/spawnTemplates folder might not exist.
	if err = _copyFolder(srcPaths.OutputRendererSpawnTemplates, dstPaths.OutputRendererSpawnTemplates); err != nil {
		return
	}

	// The ./.kickwasm/ folder.
	// ./.kickwasm
	if err = _copyFolder(srcPaths.OutputDotKickwasm, dstPaths.OutputDotKickwasm); err != nil {
		return
	}
	return
}

func _copyFolder(src, dst string) (err error) {
	// The source folder must exist.
	s := filepath.Clean(src)
	if _, err = os.Stat(s); err != nil {
		if os.IsNotExist(err) {
			// not exist error so ignore this folder
			err = nil
			return
		}
		return
	}
	// The source folder exists.
	var dcopy *ftools.DCopy
	if dcopy, err = ftools.NewDCopy(src, dst, true, true, nil); err != nil {
		return
	}
	err = dcopy.Copy()
	return
}
