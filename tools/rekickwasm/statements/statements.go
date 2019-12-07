package statements

// DotKickNotFoundFormat is a message format that states the .kickwasm folder does not exist.
const (
	DotKickNotFoundFormat = `Oops!

There is no ./.kickwasm folder in %s
`
	// WrongVersionFormat is a message format that states the current application being refactored was built with the wrong version of kickwasm.
	WrongVersionFormat = `Oops!
  
  The application was built with kickwasm version %d but rekickwasm only works with builds from version %d on.
  `

	// Help is verbose how to.
	Help = `Help:

Rekickwasm helps you refactor the renderer part of your application in small steps.
Run rekickwasm in your application's source folder. The one that contains the ./.kickwasm/ folder.

START BY INITIALIZING REKICK WITH "rekickwasm -i"
  * Creates the folder ./rekickwasm/
  * Backs up all of your source code in ./rekickwasm/.backup/
  * Puts a copy of your yaml files in ./rekickwasm/edit/ that you must edit for the refactoring.

EDIT THE YAML FILES IN ./rekickwasm/edit/yaml/ TO MAKE YOUR CHANGES
  * Remove buttons, tabs and panels.
  * Add completely new buttons, tabs and panels.
  * Rename a home.
  * Move the buttons around.
  * Edit a button name, label, heading, cc.
  * Edit a tab name or label.
  * Resort the panels or panelFiles in a button or tab.
  * Move a button or tab panel or panelFile to a different tab or button. ( May require moving the button's or tab's panel YAML files. )
  * Move a button or tab to a different panel. ( May require moving the button's or tab's panel YAML files. )
EDIT ./rekickwasm/edit/flags.yaml TO MAKE ANOTHER CHANGE
  * Edit flagcc if you want to toggle your previous cookie crumbs flag to "true" or "false".

REFACTOR: "rekickwasm -R"
  * Refactors your source code.
  * Allows you to continue editing the yaml files in ./rekickwasm/edit/ and then make those changes with "rekickwasm -R".

REBUILD THE RENDERER AND RUN YOUR APPLICATION:

============================================================

  $ kickbuild -rp -mp -run

============================================================

RESTORE YOUR SOURCE CODE TO IT'S ORIGINAL STATE: "rekickwasm -u"
  * Imports the backup ./panelMap.go
  * Imports the backup ./rendererprocess/css/
  * Imports the backup ./rendererprocess/panels.go
  * Imports the backup ./rendererprocess/panels/
  * Imports the backup ./rendererprocess/spawnPanels/
  * Imports the backup ./site/templates/
  * Imports the backup ./site/spawnTemplates/
  * Imports the backup ./rendererprocess/viewtools/
  * Imports the backup ./.kickwasm/
  * Allows you to continue editing the yaml files in ./rekickwasm/edit/ and then make those changes with "rekickwasm -R".

QUITTING REKICK: "rekickwasm -x"
  * Deletes the ./rekickwasm/ folder.
  * Without the ./rekickwasm/ folder you can not restore your source code.

`

	// Warning is another verbose help.
	Warning = `WARNING regarding the YAML files in ./rekickwasm/edit/yaml/:
Rekickwasm is not a tool for editing the main process.
So when you edit the files in ./rekickwasm/edit/yaml/

In ./rekickwasm/edit/yaml/kickwasm.yaml
  * Do not edit title or importPath.
  * If you rename a panel rekickwasm will remove it and you will lose it.
    For example if you rename "MyOldPanel" to "MyNewPanel",
     rekickwasm will remove the "MyOldPanel" template file and go package files.
    Then rekickwasm will add a new template file and new go package files for "MyNewPanel".

  AND REMEMBER:
    * Home names must be unique.
    * Panel names must be unique.
    * Button names must be unique.
    * Tab names must be unique.

In ./rekickwasm/edit/flags.yaml DO NOT EDIT
  * kwversionbreaking, kwversionfeature or kwversionpatch.
`

	// AlreadyInit is a message that states the current folder has already been initialized.
	AlreadyInit = `Initialize: Oops!
	
rekickwasm stopped without doing anything because rekickwasm has already been initialized in the current folder.

So, if you haven't already, you need to edit the YAML files in ./rekickwasm/edit/

On the other hand...
If you want to overwrite the changes you have made to the yaml files in ./rekickwasm/edit/
Then you can do, "rekickwasm -yaml"

`

	// SuccessInit is a message that states the current folder has been initialized successfully.
	SuccessInit = `Initialize: Success!

Rekickwasm just copied the current yaml files to ./rekickwasm/edit/ for you to edit for the refactoring.
So now you can
  * Edit the YAML files in ./rekickwasm/edit/yaml/ to show the changes that you want to make.
  * In the YAML file ./rekickwasm/edit/flags.yaml, edit the bool value of flagCC to turn cookie crumbs on or off.

If for some reason, you want to overwrite the changes you have made to the yaml files in ./rekickwasm/edit/yaml/
Then you can do, "rekickwasm -yaml"

`

	// ErrorInitFormat is a message format that states the error which occurred while initializing the current folder.
	ErrorInitFormat = `Initialize: Oops!

The following error stopped the initialization process.

%s

`

	// NotInitialized is a message that states the .rekickwasm/ folder does not exist in the current folder.
	NotInitialized = `Oops!

rekickwasm stopped without doing anything because rekickwasm has not been initialized in this current folder.

You're getting ahead of yourself so slow down...
Start with "rekickwasm -i" to create a copy of the yaml files at ./rekickwasm/edit/.
Then you can
  * Edit the YAML files in ./rekickwasm/edit/yaml/ to show the changes that you want to make.
  * In the YAML file ./rekickwasm/edit/flags.yaml, edit the bool value of flagCC to turn cookie crumbs on or off.

`

	// SuccessImport is a message that states the changes have been imported into the original source code.
	SuccessImport = `Refactor Import: Success!

Your application is now refactored.
Also, the ./.kickwasm/ folder contains the new YAML files that you edited for this refactoring.

`

	// ErrorImportFormat is a message format that states the error while importing the changes.
	ErrorImportFormat = `Refactor Import: Oops!

The following error stopped the importation of the refactored renderer code into your source code.

%s

You should "rekickwasm -u" to restore your source code.

`

	// SuccessRefactor is a message that states that the edited yaml file was successfully built and merged with a copy of the original source code.
	SuccessRefactor = `Refactor: Success!`

	// ErrorRefactorFormat is a message format that states the error while refactoring.
	ErrorRefactorFormat = `Refactor: Oops!

The following error(s) stopped the refactoring of the copy of your source code.:

%s

`

	// NothingRefactored is a message that states that there were no changes to refactor.
	NothingRefactored = `There were no changes so there is nothing to refactor.`

	// SuccessUndo is a message that states that the original source code has been restored to its original state.
	SuccessUndo = `Refactor Undo: Success!

The application's source code is now restored to it's oringinal state.
The application's ./.kickwasm/ folder is now restored to it's oringinal state.

`

	// ErrorUndoFormat is a message format that states error which occured while attempting to undo.
	ErrorUndoFormat = `Refactor-Undo: Oops!

The following error stopped the restoration of your source code.

%s

I hope you have your own backup of your source code.

`

	// SuccessClean is a message that states the .rekickwasm folder has been deleted.
	SuccessClean = `Exit: Success!

./rekickwasm/ is now deleted.

`

	// ErrorCleanFormat is a message format that states the error that occurred while attempting to delete the ./rekickwasm folder.
	ErrorCleanFormat = `Exit: Oops!

The following error stopped the deletion of the ./rekickwasm/ folder.:

%s

`

	// ErrYAMLFormat is  is a message format that states the error that occurred while attempting to restore the editable yaml files to their original unedited state.
	ErrYAMLFormat = `YAML: Oops!

The following error stopped the restoration of ./rekickwasm/edit/:

%s

`

	// SuccessClean is a message that states the editable yaml files have been restored to their original unedited state.
	SuccessYAML = `YAML: Success!

The files in ./rekickwasm/edit/ have been restored from the backup.
So you can
  * Edit the YAML files in ./rekickwasm/edit/yaml/ to show the changes that you want to make.
  * In the YAML file ./rekickwasm/edit/flags.yaml, edit the bool value of flagCC to turn cookie crumbs on or off.

`
)
