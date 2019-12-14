package refactor

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/tools/rekickwasm/ftools"
	"github.com/josephbudd/kickwasm/tools/rekickwasm/statements"

	"github.com/josephbudd/kickwasm/pkg/flagdata"
	"github.com/josephbudd/kickwasm/pkg/mainprocess"
	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/renderer"
	"github.com/josephbudd/kickwasm/pkg/slurp"

	"github.com/josephbudd/kickwasm/tools/rekickwasm/repath"
)

const (
	panelDotGoTxtFile     = "original_panel_dot_go.txt"
	templateFileExtension = ".tmpl"
)

// Refactorer refactors the application.
type Refactorer struct {
	rp *repath.RePaths
}

// NewRefactorer constructs a new Refactorer
func NewRefactorer(rp *repath.RePaths) *Refactorer {
	return &Refactorer{rp}
}

// Refactor returns error
// 0.  Prep.
// 1.  Make the refactor builder.
// 2.  Build the changes source code in the changes folder.
// 2.a Copy the edited yaml to changes.
// 2.b Make the changes builder.
// 2.c Build the changes renderer source.
// 3   If errors then return.
// 4.  If no differcences then return.
// 5.  Merge Changes source code into Refactor source code (original source code).
// 6.  Cleanup
func (r *Refactorer) Refactor() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf(statements.ErrorRefactorFormat, err.Error())
		}
	}()

	// Step 0: Prep.
	// Prep the changes folder for the changes build.
	if err = r.rp.InitializeChanges(); err != nil {
		return
	}
	// Prep the refactor folder for the refactoring.
	if err = r.rp.InitializeRefactor(); err != nil {
		return
	}

	mergePaths := r.rp.Refactor.GetPaths()
	// Step 1:  Make the refactor builder.
	var dcopy *ftools.DCopy
	var mergeFlags *flagdata.Flags
	if mergeFlags, err = flagdata.GetFlags(mergePaths.OutputDotKickwasmFlags); err != nil {
		return
	}
	mergeYAMLPath := filepath.Join(mergePaths.OutputDotKickwasmYAML, mergeFlags.YAMLStartFileName)
	sl := slurp.NewSlurper()
	var mergeBuilder *project.Builder
	if mergeBuilder, err = sl.Gulp(mergeYAMLPath); err != nil {
		return
	}
	_ = mergeBuilder.ToHTMLNode("", mergeFlags.FlagCC)
	// Step 2: Build the changes source code in the changes folder.
	// Step 2.a: Copy the edited yaml to changes.
	changesPaths := r.rp.Changes.GetPaths()
	dcopy, err = ftools.NewDCopy(
		r.rp.RekickWasmEdit,
		changesPaths.OutputDotKickwasm,
		true, true,
		nil,
	)
	if err != nil {
		return
	}
	if err = dcopy.Copy(); err != nil {
		return
	}
	// Step 2.b: Make the changes builder.
	sl = slurp.NewSlurper()
	var changesFlags *flagdata.Flags
	if changesFlags, err = flagdata.GetFlags(changesPaths.OutputDotKickwasmFlags); err != nil {
		return
	}
	if err = checkVersion(changesFlags, mergeFlags); err != nil {
		return
	}
	changesYAMLPath := filepath.Join(changesPaths.OutputDotKickwasmYAML, changesFlags.YAMLStartFileName)
	var changesBuilder *project.Builder
	if changesBuilder, err = sl.Gulp(changesYAMLPath); err != nil {
		return
	}
	// If the user edited anything that is main process it is an error.
	if err = checkMainProcess(changesBuilder, mergeBuilder); err != nil {
		return
	}
	// Step 2.c: Build the changes renderer source.
	if err = renderer.Create(r.rp.Changes, changesBuilder, changesFlags.FlagCC); err != nil {
		return
	}
	// Step 3: If errors then return.
	if err == nil {
		if err = ButtonPanelErrors(changesBuilder, mergeBuilder); err == nil {
			err = TabPanelErrors(changesBuilder, mergeBuilder)
		}
	}
	if err != nil {
		return
	}
	// Step 4: If no differences then return.
	var removals, additions map[string]SpawnPath
	var moves map[string]MoveSpawnPath
	var differences, scrollChanged bool
	removals, additions, moves, differences, scrollChanged = DifPanelPaths(changesBuilder, mergeBuilder)
	if !differences {
		if differences = (mergeFlags.FlagCC != changesFlags.FlagCC) || len(removals) > 0 || len(additions) > 0 || len(moves) > 0; !differences {
			if differences = DifHomePositionsButtons(changesBuilder, mergeBuilder); !differences {
				if differences = DifButtonPanelPositions(changesBuilder, mergeBuilder); !differences {
					differences = DifTabPanelPositions(changesBuilder, mergeBuilder)
				}
			}
		}
	}
	if !differences && !scrollChanged {
		err = fmt.Errorf(statements.ErrorRefactorFormat, statements.NothingRefactored)
		return
	}
	// Step 5: merge changes folder into refactor ( original ) folder.
	//         panels/, spawnPanels/, templates/ and spawnTemplates/,
	//         proofs/, viewtools/.
	// Step 5.a panels/ and spawnPanels/ templates/ spawnTemplates/.
	if err = r.refactorPanels(changesBuilder, mergeBuilder, removals, additions, moves); err != nil {
		return
	}
	// Step 5.b Copy uneditable files.
	if err = r.refactorPanelGroupFiles(changesBuilder, mergeBuilder, removals, additions, moves); err != nil {
		return
	}
	if err = r.refactorSpawnPanelUneditableFiles(changesBuilder, mergeBuilder, removals, additions, moves); err != nil {
		return
	}
	// Step 5.c Remove unused panel and template folders.
	if err = r.removeUnusedPanelTemplateFolders(); err != nil {
		return
	}

	// Copy the user edited yaml to the refactor folder.
	dcopy, err = ftools.NewDCopy(
		r.rp.RekickWasmEdit,
		mergePaths.OutputDotKickwasm,
		true, true,
		nil,
	)
	if err != nil {
		return
	}
	if err = dcopy.Copy(); err != nil {
		return
	}
	// Copy renderer/lpc/client.go
	refactorPaths := r.rp.Refactor.GetPaths()
	filenames := r.rp.Refactor.GetFileNames()
	src := filepath.Join(changesPaths.OutputRendererLPC, filenames.ClientDotGo)
	dst := filepath.Join(refactorPaths.OutputRendererLPC, filenames.ClientDotGo)
	if err = os.Remove(dst); err != nil {
		return
	}
	if err = ftools.CopyFile(src, dst); err != nil {
		return
	}
	// Copy site/templates/main.html
	src = filepath.Join(changesPaths.OutputRendererTemplates, filenames.MainDotTMPL)
	dst = filepath.Join(refactorPaths.OutputRendererTemplates, filenames.MainDotTMPL)
	if err = os.Remove(dst); err != nil {
		return
	}
	if err = ftools.CopyFile(src, dst); err != nil {
		return
	}
	// Step 6: Cleanup
	// Remove the changes folder
	if err = os.RemoveAll(changesPaths.Output); err != nil {
		return
	}
	return
}

func checkVersion(changesFlags, mergeFlags *flagdata.Flags) (err error) {
	if changesFlags.KWVersionBreaking != mergeFlags.KWVersionBreaking {
		err = fmt.Errorf("you altered the kwversionbreaking in your yaml file")
		return
	}
	if changesFlags.KWVersionFeature != mergeFlags.KWVersionFeature {
		err = fmt.Errorf("you altered the kwversionfeature in your yaml file")
		return
	}
	if changesFlags.KWVersionPatch != mergeFlags.KWVersionPatch {
		err = fmt.Errorf("you altered the kwversionpatch in your yaml file")
		return
	}
	return
}

// checkMainProcess returns an error if the user edited any main process parts of the main yaml file.
func checkMainProcess(changesBuilder, mergeBuilder *project.Builder) (err error) {
	if changesBuilder.Title != mergeBuilder.Title {
		err = fmt.Errorf("you altered the title in your yaml file")
		return
	}
	if changesBuilder.ImportPath != mergeBuilder.ImportPath {
		err = fmt.Errorf("you altered the import path in your yaml file")
		return
	}
	return
}

// refactorPanels rebuilds the application preserving what must not be changed.
// The refactor folder is a copy of the original source.
// The changes folder is the source built from the new yaml edits.
// If the user is modifying the frammerge's renderer
//  then the changes source code is merged into the refactor source code
//  so that the refactor source is the refactored source code.
// Returns if modified and the error.
func (r *Refactorer) refactorPanels(changesBuilder, mergeBuilder *project.Builder, removals, additions map[string]SpawnPath, moves map[string]MoveSpawnPath) (err error) {
	// The changes folder contains the changes that the user wants to make.
	// The refactor folder contains a copy of the current source code.
	// The changes will be applied to the refactor folder
	//  so that the refactor folder will be the refactored source code.
	//
	// Step 1: refactor site/templates/ and site/spawnTemplates
	//         refactor renderer/panels/ and renderer/spawnPanels/.
	if err = r.refactorPanelPaths(changesBuilder, mergeBuilder, removals, additions, moves); err != nil {
		return
	}
	// Step 2: NA
	// Step 3: build the changes main process source.
	if err = mainprocess.Create(r.rp.Changes, changesBuilder); err != nil {
		return
	}

	mergePaths := r.rp.Refactor.GetPaths()
	changesPaths := r.rp.Changes.GetPaths()
	fileNames := paths.GetFileNames()
	var dc *ftools.DCopy
	// Step 4.a: ./rendererprocess/viewtools/
	if dc, err = ftools.NewDCopy(
		changesPaths.OutputRendererViewTools, mergePaths.OutputRendererViewTools,
		true, false, nil); err != nil {
		return
	}
	if err = dc.Copy(); err != nil {
		return
	}
	// Step 4.b: ./rendererprocess/proofs/
	if dc, err = ftools.NewDCopy(
		changesPaths.OutputRendererProofs, mergePaths.OutputRendererProofs,
		true, false, nil); err != nil {
		return
	}
	if err = dc.Copy(); err != nil {
		return
	}

	// Step 5: ./site/templates/main.tmpl
	src := filepath.Join(changesPaths.OutputRendererTemplates, fileNames.MainDotTMPL)
	dst := filepath.Join(mergePaths.OutputRendererTemplates, fileNames.MainDotTMPL)
	if err = ftools.CopyFile(src, dst); err != nil {
		return
	}
	// Step 6: ./rendererprocess/framework/panels.go
	src = filepath.Join(changesPaths.OutputRendererFramework, fileNames.PanelsDotGo)
	dst = filepath.Join(mergePaths.OutputRendererFramework, fileNames.PanelsDotGo)
	if err = ftools.CopyFile(src, dst); err != nil {
		return
	}
	// Step 7: ./rendererprocess/css/*
	if dc, err = ftools.NewDCopy(
		changesPaths.OutputRendererCSS, mergePaths.OutputRendererCSS,
		true, false, nil); err != nil {
		return
	}
	if err = dc.Copy(); err != nil {
		return
	}

	// Step 8: ./panelMap.go
	src = filepath.Join(changesPaths.OutputMainProcess, fileNames.PanelMapDotGo)
	dst = filepath.Join(mergePaths.OutputMainProcess, fileNames.PanelMapDotGo)
	if err = ftools.CopyFile(src, dst); err != nil {
		return
	}
	// Step 9: yaml files.
	if dc, err = ftools.NewDCopy(
		r.rp.RekickWasmEdit, mergePaths.OutputDotKickwasm,
		true, false, nil); err != nil {
		return
	}
	if err = dc.Copy(); err != nil {
		return
	}
	return
}

// refactorPanelPaths does a folder and file management.
func (r *Refactorer) refactorPanelPaths(changesBuilder, mergeBuilder *project.Builder, removals, additions map[string]SpawnPath, moves map[string]MoveSpawnPath) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("refactorPanelPaths: %w", err)
		}
	}()

	var dcopy *ftools.DCopy
	var fm *ftools.FMove
	var src, dst string
	var changesStartFolder string
	var mergeStartFolder string
	changesPaths := r.rp.Changes.GetPaths()
	mergePaths := r.rp.Refactor.GetPaths()

	// Additions
	for panelName, spawnPath := range additions {
		// renderer/panels/ folders
		// Copy the entire changes package folder to the refactor panels/ folder.
		if spawnPath.Spawn {
			changesStartFolder = changesPaths.OutputRendererSpawns
			mergeStartFolder = mergePaths.OutputRendererSpawns
		} else {
			changesStartFolder = changesPaths.OutputRendererPanels
			mergeStartFolder = mergePaths.OutputRendererPanels
		}

		// Markup panels folders.
		src = filepath.Join(changesStartFolder, spawnPath.Path, panelName)
		dst = filepath.Join(mergeStartFolder, spawnPath.Path, panelName)
		if dcopy, err = ftools.NewDCopy(src, dst, false, false, nil); err != nil {
			err = fmt.Errorf("additions dcopy: %w", err)
			return
		}
		if err = dcopy.Copy(); err != nil {
			err = fmt.Errorf("additions copy: %w", err)
			return
		}
		// Template files.
		if spawnPath.Spawn {
			changesStartFolder = changesPaths.OutputRendererSpawnTemplates
			mergeStartFolder = mergePaths.OutputRendererSpawnTemplates
		} else {
			changesStartFolder = changesPaths.OutputRendererTemplates
			mergeStartFolder = mergePaths.OutputRendererTemplates
		}
		fname := panelName + templateFileExtension
		src = filepath.Join(changesStartFolder, spawnPath.Path, fname)
		dst = filepath.Join(mergeStartFolder, spawnPath.Path, fname)
		if err = ftools.CopyFile(src, dst); err != nil {
			err = fmt.Errorf("additions: %w", err)
			return
		}
	}

	// Moves
	dmode := r.rp.Changes.GetDMode()
	for panelName, moveSpawnPath := range moves {
		// renderer/panels/ folders
		// Move the entire refactor package folder to the new refactor panels/ folder.
		if moveSpawnPath.From.Spawn {
			mergeStartFolder = mergePaths.OutputRendererSpawns
		} else {
			mergeStartFolder = mergePaths.OutputRendererPanels
		}
		src = filepath.Join(mergeStartFolder, moveSpawnPath.From.Path, panelName)
		dst = filepath.Join(mergeStartFolder, moveSpawnPath.To.Path, panelName)
		var dm *ftools.DMove
		if dm, err = ftools.NewDMove(src, dst, false, false, nil); err != nil {
			err = fmt.Errorf("moves: %w", err)
			return
		}
		if err = dm.Move(); err != nil {
			err = fmt.Errorf("moves: %w", err)
			return
		}

		// Template files.
		if moveSpawnPath.From.Spawn {
			mergeStartFolder = mergePaths.OutputRendererSpawnTemplates
		} else {
			mergeStartFolder = mergePaths.OutputRendererTemplates
		}
		fname := panelName + templateFileExtension
		src = filepath.Join(mergeStartFolder, moveSpawnPath.From.Path, fname)
		dst = filepath.Join(mergeStartFolder, moveSpawnPath.To.Path, fname)
		fm = ftools.NewFMove(src, dst, true, dmode)
		if err = fm.Move(); err != nil {
			err = fmt.Errorf("moves: %w", err)
			return
		}
	}

	// Removals
	for panelName, removeSpawnPath := range removals {
		// renderer/panels/ folders
		if removeSpawnPath.Spawn {
			mergeStartFolder = mergePaths.OutputRendererSpawns
		} else {
			mergeStartFolder = mergePaths.OutputRendererPanels
		}
		dst = filepath.Join(mergeStartFolder, removeSpawnPath.Path, panelName)
		if err = os.RemoveAll(dst); err != nil {
			err = fmt.Errorf("removals: %w", err)
			return
		}

		// Template files.
		if removeSpawnPath.Spawn {
			mergeStartFolder = mergePaths.OutputRendererSpawnTemplates
		} else {
			mergeStartFolder = mergePaths.OutputRendererTemplates
		}
		fname := panelName + templateFileExtension
		src = filepath.Join(mergeStartFolder, removeSpawnPath.Path, fname)
		if err = os.RemoveAll(src); err != nil {
			err = fmt.Errorf("removals: %w", err)
			return
		}
	}
	return
}

// refactorPanelGroupFiles copies group files.
func (r *Refactorer) refactorPanelGroupFiles(changesBuilder, mergeBuilder *project.Builder, removals, additions map[string]SpawnPath, moves map[string]MoveSpawnPath) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("refactorPanelGroupFiles: %w", err)
		}
	}()

	var src, dst string
	var gpsrc, gpdst string
	changesPaths := r.rp.Changes.GetPaths()
	mergePaths := r.rp.Refactor.GetPaths()
	fileNames := r.rp.Changes.GetFileNames()
	redundancy := make(map[string]bool)
	var found bool

	var panelNames []string
	var panelName string

	// Additions
	for _, spawnPath := range additions {
		if spawnPath.Spawn {
			src = filepath.Join(changesPaths.OutputRendererSpawns, spawnPath.Path)
			dst = filepath.Join(mergePaths.OutputRendererSpawns, spawnPath.Path)
		} else {
			src = filepath.Join(changesPaths.OutputRendererPanels, spawnPath.Path)
			dst = filepath.Join(mergePaths.OutputRendererPanels, spawnPath.Path)
		}
		// Build the source and destination parent folders.
		// Undeditables are always copied from changes to merge.
		// Redudancy check.
		if _, found = redundancy[src]; found {
			continue
		}
		redundancy[src] = true
		// Get the panel names ( folder names ) from the source.
		if panelNames, err = getPanelNamesInFolder(src); err != nil {
			err = fmt.Errorf("additions getPanelNamesInFolder: %w", err)
			return
		}
		// Copy group files from src to dst.
		for _, panelName = range panelNames {
			gpsrc = filepath.Join(src, panelName, fileNames.PanelGroupDotGo)
			gpdst = filepath.Join(dst, panelName, fileNames.PanelGroupDotGo)
			if err = ftools.CopyFile(gpsrc, gpdst); err != nil {
				err = fmt.Errorf("additions: %w", err)
				return
			}
		}
	}

	// Moves
	for panelName, moveSpawnPath := range moves {
		if moveSpawnPath.From.Spawn {
			src = filepath.Join(changesPaths.OutputRendererSpawns, moveSpawnPath.To.Path)
			dst = filepath.Join(mergePaths.OutputRendererSpawns, moveSpawnPath.To.Path)
		} else {
			src = filepath.Join(changesPaths.OutputRendererPanels, moveSpawnPath.To.Path)
			dst = filepath.Join(mergePaths.OutputRendererPanels, moveSpawnPath.To.Path)
		}
		// Build the source and destination parent folders.
		// Undeditables are always copied from changes to merge.
		// Redudancy check.
		if _, found = redundancy[src]; found {
			continue
		}
		redundancy[src] = true
		dst = filepath.Join(mergePaths.OutputRendererPanels, moveSpawnPath.To.Path)
		// Get the panel names ( folder names ) from the source.
		if panelNames, err = getPanelNamesInFolder(src); err != nil {
			err = fmt.Errorf("moves getPanelNamesInFolder: %w", err)
			return
		}
		// Copy group files from src to dst.
		for _, panelName = range panelNames {
			gpsrc = filepath.Join(src, panelName, fileNames.PanelGroupDotGo)
			gpdst = filepath.Join(dst, panelName, fileNames.PanelGroupDotGo)
			if err = ftools.CopyFile(gpsrc, gpdst); err != nil {
				err = fmt.Errorf("moves: %w", err)
				return
			}
		}
	}

	// Removals
	for panelName, removeSpawnPath := range removals {
		if removeSpawnPath.Spawn {
			src = filepath.Join(changesPaths.OutputRendererSpawns, removeSpawnPath.Path)
			dst = filepath.Join(mergePaths.OutputRendererSpawns, removeSpawnPath.Path)
		} else {
			src = filepath.Join(changesPaths.OutputRendererPanels, removeSpawnPath.Path)
			dst = filepath.Join(mergePaths.OutputRendererPanels, removeSpawnPath.Path)
		}
		// Build the source and destination parent folders.
		// Undeditables are always copied from changes to merge.
		src = filepath.Join(changesPaths.OutputRendererPanels, removeSpawnPath.Path)
		// Redudancy check.
		if _, found = redundancy[src]; found {
			continue
		}
		redundancy[src] = true
		dst = filepath.Join(mergePaths.OutputRendererPanels, removeSpawnPath.Path)
		// Get the panel names ( folder names ) from the source.
		if panelNames, err = getPanelNamesInFolder(src); err != nil {
			err = fmt.Errorf("removes getPanelNamesInFolder: %w", err)
			return
		}
		// Copy group files from src to dst.
		for _, panelName = range panelNames {
			gpsrc = filepath.Join(src, panelName, fileNames.PanelGroupDotGo)
			gpdst = filepath.Join(dst, panelName, fileNames.PanelGroupDotGo)
			if err = ftools.CopyFile(gpsrc, gpdst); err != nil {
				err = fmt.Errorf("removals: %w", err)
				return
			}
		}
	}

	return
}

func (r *Refactorer) refactorSpawnPanelUneditableFiles(changesBuilder, mergeBuilder *project.Builder, removals, additions map[string]SpawnPath, moves map[string]MoveSpawnPath) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("refactorSpawnPanelUneditableFiles: %w", err)
		}
	}()

	var src, dst string
	changesPaths := r.rp.Changes.GetPaths()
	mergePaths := r.rp.Refactor.GetPaths()
	fileNames := r.rp.Changes.GetFileNames()
	redundancy := make(map[string]bool)

	// Additions
	for panelName, spawnPath := range additions {
		if !spawnPath.Spawn {
			continue
		}
		// Build the source and destination tab folders.
		// Undeditables are always copied from changes to merge.
		src = filepath.Join(changesPaths.OutputRendererSpawns, spawnPath.Path, panelName)
		dst = filepath.Join(mergePaths.OutputRendererSpawns, spawnPath.Path, panelName)
		if err = copySpawnUnEditables(src, dst, fileNames, redundancy); err != nil {
			err = fmt.Errorf("additions: %w", err)
			return
		}
	}

	// Moves
	for panelName, moveSpawnPath := range moves {
		if !moveSpawnPath.From.Spawn {
			continue
		}
		// Fix the move to folder.
		// Build the source and destination tab folders.
		// Undeditables are always copied from changes to merge.
		src = filepath.Join(changesPaths.OutputRendererSpawns, moveSpawnPath.To.Path, panelName)
		dst = filepath.Join(mergePaths.OutputRendererSpawns, moveSpawnPath.To.Path, panelName)
		if err = copySpawnUnEditables(src, dst, fileNames, redundancy); err != nil {
			err = fmt.Errorf("moves to: %w", err)
			return
		}
		// Fix the move from folder.
		// Build the source and destination tab folders.
		// Undeditables are always copied from changes to merge.
		src = filepath.Join(changesPaths.OutputRendererSpawns, moveSpawnPath.From.Path, panelName)
		dst = filepath.Join(mergePaths.OutputRendererSpawns, moveSpawnPath.From.Path, panelName)
		if err = copySpawnUnEditables(src, dst, fileNames, redundancy); err != nil {
			err = fmt.Errorf("moves from: %w", err)
			return
		}
	}

	// Removals
	for panelName, removeSpawnPath := range removals {
		if !removeSpawnPath.Spawn {
			continue
		}
		// Build the source and destination tab folders.
		// Undeditables are always copied from changes to merge.
		src = filepath.Join(changesPaths.OutputRendererSpawns, removeSpawnPath.Path, panelName)
		dst = filepath.Join(mergePaths.OutputRendererSpawns, removeSpawnPath.Path, panelName)
		if err = copySpawnUnEditables(src, dst, fileNames, redundancy); err != nil {
			err = fmt.Errorf("removals: %w", err)
			return
		}
	}
	return
}

func (r *Refactorer) removeUnusedPanelTemplateFolders() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("Refactorer.removeUnusedPanelTemplateFolders(): %w", err)
		}
	}()

	changesPaths := r.rp.Changes.GetPaths()
	mergePaths := r.rp.Refactor.GetPaths()
	// Panel packages.
	if err = walkFolders(changesPaths.OutputRendererPanels, mergePaths.OutputRendererPanels); err != nil {
		err = fmt.Errorf("OutputRendererPanels: %w", err)
		return
	}
	if err = walkFolders(changesPaths.OutputRendererSpawns, mergePaths.OutputRendererSpawns); err != nil {
		err = fmt.Errorf("OutputRendererSpawns: %w", err)
		return
	}
	// Panel templates.
	if err = walkFolders(changesPaths.OutputRendererTemplates, mergePaths.OutputRendererTemplates); err != nil {
		err = fmt.Errorf("OutputRendererTemplates: %w", err)
		return
	}
	if err = walkFolders(changesPaths.OutputRendererSpawnTemplates, mergePaths.OutputRendererSpawnTemplates); err != nil {
		err = fmt.Errorf("OutputRendererSpawnTemplates: %w", err)
	}
	return
}

func walkFolders(changesPath, mergePath string) (err error) {
	haveChangesPath := true
	_, err = os.Stat(changesPath)
	if os.IsNotExist(err) {
		haveChangesPath = false
	}
	haveMergePath := true
	_, err = os.Stat(mergePath)
	if os.IsNotExist(err) {
		haveMergePath = false
	}
	if !haveChangesPath && !haveMergePath {
		err = nil
		return
	}
	if !haveChangesPath && haveMergePath {
		// This folder does not exist in changes.
		//   so try to remove it from merge.
		err = os.RemoveAll(mergePath)
		return
	}
	// Both the changesPath and the mergePath exist.
	offset := len(mergePath)
	err = filepath.Walk(
		mergePath,
		func(path string, info os.FileInfo, er error) (err error) {
			if er != nil {
				err = er
				return
			}
			if info.IsDir() {
				// Find this folder in changes.
				ch := filepath.Join(changesPath, path[offset:])
				if _, err = os.Stat(ch); os.IsNotExist(err) {
					// This folder does not exist in changes.
					//   so try to remove it from merge.
					// If it does not exist in merge that's ok.
					if err = os.RemoveAll(path); err != nil && !os.IsNotExist(err) {
						return
					}
					// The dir does not exist anymore so skip it.
					err = filepath.SkipDir
					return
				}
			}
			return
		})
	return
}

func copySpawnUnEditables(panelsrc, paneldst string, fileNames *paths.FileNames, redundancy map[string]bool) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("copySpawnUnEditables: %w", err)
		}
	}()

	var fsrc, fdst string
	fmt.Printf("copySpawnUnEditables: %s \n  to  %s\n", panelsrc, paneldst)
	// The spawn panel folder.
	// api.go
	if haveFolders(panelsrc, paneldst) {
		fsrc = filepath.Join(panelsrc, fileNames.APIDotGo)
		fdst = filepath.Join(paneldst, fileNames.APIDotGo)
		if err = ftools.CopyFile(fsrc, fdst); err != nil {
			return
		}
		// group.go
		fsrc = filepath.Join(panelsrc, fileNames.PanelGroupDotGo)
		fdst = filepath.Join(paneldst, fileNames.PanelGroupDotGo)
		if err = ftools.CopyFile(fsrc, fdst); err != nil {
			return
		}
	}

	// The spawn tab folders.
	var tabsrc, tabdst string
	tabsrc = filepath.Dir(panelsrc)
	tabdst = filepath.Dir(paneldst)
	fmt.Printf("  * copySpawnUnEditables: %s \n  to  %s\n", tabsrc, tabdst)
	if haveFolders(tabsrc, tabdst) {
		// prepare.go
		fsrc = filepath.Join(tabsrc, fileNames.PrepareDotGo)
		fdst = filepath.Join(tabdst, fileNames.PrepareDotGo)
		if err = ftools.CopyFile(fsrc, fdst); err != nil {
			return
		}
		// spawn.go
		fsrc = filepath.Join(tabsrc, fileNames.SpawnDotGo)
		fdst = filepath.Join(tabdst, fileNames.SpawnDotGo)
		if err = ftools.CopyFile(fsrc, fdst); err != nil {
			return
		}
	}

	// The spawn tab bar folder.
	var tabbarsrc, tabbardst string
	var found bool
	tabbarsrc = filepath.Dir(tabsrc)
	_, found = redundancy[tabbarsrc]
	if found {
		fmt.Printf("  * found %s\n", tabbarsrc)
	}
	if !found {
		redundancy[tabbarsrc] = true
		tabbardst = filepath.Dir(tabdst)
		if haveFolders(tabbarsrc, tabbardst) {
			// prepare.go
			fsrc = filepath.Join(tabbarsrc, fileNames.PrepareDotGo)
			fdst = filepath.Join(tabbardst, fileNames.PrepareDotGo)
			fmt.Printf("  * copying %s to %s\n", fsrc, fdst)
			if err = ftools.CopyFile(fsrc, fdst); err != nil {
				return
			}
		}
	}
	return
}

func getPanelNamesInFolder(path string) (panelNames []string, err error) {
	var dir *os.File
	if dir, err = os.Open(path); err != nil {
		return
	}

	var fis []os.FileInfo
	if fis, err = dir.Readdir(-1); err != nil {
		return
	}
	panelNames = make([]string, 0, len(fis))
	var fi os.FileInfo
	for _, fi = range fis {
		if fi.IsDir() {
			panelNames = append(panelNames, fi.Name())
		}
	}
	return
}

func haveFolders(paths ...string) (haveFolders bool) {
	haveFolders = true
	for _, path := range paths {
		if _, err := os.Stat(path); err != nil {
			haveFolders = false
			return
		}
	}
	return
}
