package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickwasm/tools/common"
	"github.com/josephbudd/kickwasm/tools/rekickwasm/statements"

	"github.com/josephbudd/kickwasm/tools/rekickwasm/refactor"
	"github.com/josephbudd/kickwasm/tools/rekickwasm/repath"
	"github.com/josephbudd/kickwasm/tools/script"
)

const (
	applicationName        = "rekickwasm"
	versionBreaking        = 16 // Kicwasm Breaking Version. (Backwards compatibility.)
	versionFeature         = 0  // Added features. Still backwards compatible.
	versionPatch           = 0  // Bug fix. No added features.
	minumunKickwasmVersion = 14 // Minumum kickwasm version.
)

// CleanFlag means remove the ./rekickwasm/ folder
var CleanFlag bool

// YAMLFlag means reinitialize ./rekickwasm/yaml/
var YAMLFlag bool

// VersionFlag means show the version.
var VersionFlag bool

// RefactorFlag means refactor the application using the changes in ./rekickwasm/yaml/
var RefactorFlag bool

// RefactorImportFlag means
//   1. refactor the application using the changes in ./rekickwasm/yaml/
//   2. import the refactored code into the original code.
var RefactorImportFlag bool

// UndoFlag means undo the refactoring.
var UndoFlag bool

// InitFlag means create the initial yaml copy if it doesn't alreay exist.
var InitFlag bool

// EditFlag is the flag to the editor to use to edit ./rekickwasm/edit/yaml/kickwasm.yaml
var EditFlag bool

// EditorFlag is the flag to the editor to use to edit ./rekickwasm/edit/yaml/kickwasm.yaml
var EditorFlag string

// DoNotEditFlag forces no editing of ./rekickwasm/edit/yaml/kickwasm.yaml
var DoNotEditFlag bool

// HelpFlag means display the help.
var HelpFlag bool

func main() {

	var err error
	defer func() {
		if err != nil {
			os.Exit(1)
		}
	}()

	// initialize the flags
	flag.BoolVar(&InitFlag, "i", false, "Initializes. Backs up your source code and yaml files in ./rekickwasm/backup/. Creates yaml files for you to edit in ./rekickwasm/edit/.")
	flag.BoolVar(&EditFlag, "e", false, "Use with -i. edit. ./rekickwasm/edit/yaml/kickwasm.yaml with the default editor.")
	flag.StringVar(&EditorFlag, "E", "", "Use with -i. The editor to edit ./rekickwasm/edit/yaml/kickwasm.yaml with.")
	flag.BoolVar(&DoNotEditFlag, "dne", false, "Use with -i. Do not edit ./rekickwasm/edit/yaml/kickwasm.yaml with.")
	flag.BoolVar(&YAMLFlag, "yaml", false, "Restores ./rekickwasm/edit/yaml/ removing your edits.")
	flag.BoolVar(&RefactorImportFlag, "R", false, "Refactor using your edits in ./rekickwasm/edit/.")
	flag.BoolVar(&UndoFlag, "u", false, "Undo the refactoring")
	flag.BoolVar(&CleanFlag, "x", false, "Delete the ./rekickwasm/ folder.")
	flag.BoolVar(&HelpFlag, "?", false, "help")
	flag.BoolVar(&VersionFlag, "v", false, "version")
	flag.Parse()
	if len(EditorFlag) > 0 {
		EditFlag = true
	}
	if DoNotEditFlag {
		EditFlag = false
	}
	// The version and help flags work alone.
	// If they are used then process them and quit.
	if VersionFlag {
		fmt.Println(common.Version(applicationName, versionBreaking, versionFeature, versionPatch))
		return
	}
	if HelpFlag {
		help()
		return
	}
	// Other flags are required in order to actually use rekickwasm.
	if !InitFlag && !YAMLFlag && !RefactorImportFlag && !UndoFlag && !CleanFlag && !EditFlag {
		help()
		return
	}
	// This tools must run in the framework's root folder.
	var rootFolderPath string
	var isRoot bool
	if rootFolderPath, isRoot, err = common.IsRootFolder(); !isRoot {
		if err != nil {
			fmt.Println(err.Error())
		} else {
			help()
		}
		return
	}
	// This framework must have been built with a recent version of kickwasm.
	if kwversion := common.AppKickwasmVersion(); kwversion < minumunKickwasmVersion {
		common.PrintWrongVersion(applicationName, kwversion, minumunKickwasmVersion)
		return
	}

	var rp *repath.RePaths
	rp, err = repath.NewRePaths(rootFolderPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if InitFlag {
		// Initialize: Setup rekickwasm so the user can begin editing the yaml files.
		if err = rp.InitializeWorkingDirectory(); err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(statements.SuccessInit)
		if DoNotEditFlag {
			return
		}
		path := filepath.Join(rp.RekickWasmEditYAML, "kickwasm.yaml")
		if len(EditorFlag) > 0 {
			script.Run(EditorFlag, path)
		} else {
			script.Edit(path)
		}
		return
	}
	// The rest of these actions require that rekickwasm be initilized into the current folder.
	if !rp.Initialized() {
		fmt.Println(statements.NotInitialized)
		err = fmt.Errorf("not initializied")
		return
	}
	if EditFlag {
		path := filepath.Join(rp.RekickWasmEditYAML, "kickwasm.yaml")
		if len(EditorFlag) > 0 {
			script.Run(EditorFlag, path)
		} else {
			script.Edit(path)
		}
		return
	}
	if RefactorImportFlag {
		// Refactor
		// Make the refactorer
		var ref *refactor.Refactorer
		ref = refactor.NewRefactorer(rp)
		// Refactor into the refactor folder.
		// May return err no changes.
		if err = ref.Refactor(); err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(statements.SuccessRefactor)
		if err = rp.ImportRefactor(); err != nil {
			fmt.Printf(statements.ErrorImportFormat, err.Error())
			return
		}
		// import finished
		fmt.Println(statements.SuccessImport)
	}
	if UndoFlag {
		// undo the backup
		if err = rp.RestoreOriginal(); err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(statements.SuccessUndo)
	}
	if CleanFlag {
		// delete ./rekickwasm
		if err = os.RemoveAll(rp.RekickWasm); err != nil {
			fmt.Printf(statements.ErrorCleanFormat, err.Error())
			return
		}
		fmt.Println(statements.SuccessClean)
		return
	}
	if YAMLFlag {
		// continue
		if err := rp.RestoreYAML(); err != nil {
			fmt.Printf(statements.ErrYAMLFormat, err.Error())
			return
		}
		fmt.Println(statements.SuccessYAML)
	}
}

func help() {
	fmt.Println(common.Version(applicationName, versionBreaking, versionFeature, versionPatch))
	fmt.Println(common.UseItInRoot)
	flag.Usage()
}
