# rekickwasm

## October 22, 2019

Refactored for kickwasm version 11.0.0.

Now rekickwasm will refactor if you change a markup panel's **hvscroll** in the kickwams.yaml file.

## Oct 01, 2019

Retested with kickwasm v 10.
Rewrote all dif tests.

## Sept 2, 2019

Fixed a couple of bugs. One related to the renderer/spawnPanels/ folder and another dealing with the cookie crumbs and renderer/templates/main.tmpl. Did that while testing with kickwasm version 8.0.0.

[![Go Report Card](https://goreportcard.com/badge/github.com/josephbudd/rekickwasm)](https://goreportcard.com/report/github.com/josephbudd/rekickwasm)

## Part of the kickwasm tool chain

A tool that lets you refactor the renderer process of an application built with kickwasm.

rekickwasm...

* Makes a copy of your application's current kickwasm.yaml file. It lets you edit the file adding, subtracting and moving around panels, buttons and tabs for the GUI.
* Then it makes a refactored copy of your application's renderer process source code.
* You can then import the refactored copy of your application's renderer process source code into our application.

In a terminal, download, install and view **how-to** information with

``` shell

  go get -u github.com/josephbudd/rekickwasm
  cd $GOPATH/src/github.com/josephbudd/rekickwasm
  go install
  rekickwasm -? | less

```

## New July 30, 2019

* Tested with kickwasm version 5.0.0 by building the [CRUD application](https://github.com/josephbudd/crud).
* Then performed the following tests.
  1. Add Home Buttons and Markup Panels.
  1. Swap Normal tab and Swap tab in same tab bar.
  1. Add Tab bar with spawn tab.
  1. Move a tab and spawn tab from one tab bar panel to another.
  1. Add tab and spawn tab to a tab bar that already contains a tab and spawn tab.
  1. Move around tabs and spawn tabs in the same tab bar.
  1. Undo
  1. Swap spawn tabs so that normal tabs spawn spawns tabs in another tab bar.
  1. Move Homes Buttons.
  1. Swap Home Buttons.
  1. Swap Home Button Panels.
  1. Add a panel to a group.
  1. Swap a panel in a group.
  1. Rename a panel in a group.

## Using rekickwasm

The CRUD wiki page titled (The Edit Home YAML)[https://github.com/josephbudd/crud.wiki/The-Edit-Home-YAML.html] contains an actual step by step use of rekickwasm. It may help you to see that example.

Below is the help that **rekickwasm -?** displays.

``` text

Help:

Rekickwasm helps you refactor the renderer part of your application in small steps.
Run rekickwasm in your application's source folder. The one that contains the ./.kickwasm/ folder.

START BY INITIALIZING REKICK WITH "rekickwasm -i"
  * Creates the folder ./rekickwasm/
  * Backs up all of your source code in ./rekickwasm/.backup/
  * Puts a copy of your yaml files in ./rekickwasm/edit/ that you must edit for the refactoring.

EDIT THE YAML FILES IN ./rekickwasm/edit/yaml/ TO MAKE YOUR CHANGES
  * Remove buttons, tabs and panels.
  * Add completely new buttons, tabs and panels.
  * Move the home buttons around.
  * Edit a button name, label, heading, cc.
  * Edit a tab name or label.
  * Resort the panels or panelFiles in a button or tab.
  * Move a button or tab panel or panelFile to a different tab or button. ( May require moving the button's or tab's panel YAML files. )
  * Move a button or tab to a different panel. ( May require moving the button's or tab's panel YAML files. )
EDIT ./rekickwasm/edit/flags.yaml TO MAKE ANOTHER CHANGE
  * Edit flagcc if you want to toggle your previous cookie crumbs flag to "true" or "false".

REFACTOR: "rekickwasm -R"
  * Builds new refactored code in ./rekickwasm/.refactored/
  * Allows you to continue editing the yaml files in ./rekickwasm/edit/ and then make those changes with "rekickwasm -R".

IMPORT THE CHANGES INTO YOUR APPLICATIONS SOURCE CODE: "rekickwasm -I"
  * Imports the new refactored source code.
  * Allows you to continue editing the yaml files in ./rekickwasm/edit/ and then make those changes with "rekickwasm -R".

( "rekickwasm -RI" will refactor and import )

REBUILD THE RENDERER AND RUN YOUR APPLICATION:

============================================================

  cd ./renderer
  ./build.sh
  cd ..
  go build
  ./myapp

============================================================

RESTORE YOUR SOURCE CODE TO IT'S ORIGINAL STATE: "rekickwasm -u"
  * Restores your source code from the ./rekickwasm/.backup/ folder.
  * Allows you to continue editing the yaml files in ./rekickwasm/edit/ and then make those changes with "rekickwasm -R".

QUITTING REKICK: "rekickwasm -x"
  * Deletes the ./rekickwasm/ folder.
  * Without the ./rekickwasm/ folder you can not restore your source code.


WARNING regarding the YAML files in ./rekickwasm/edit/yaml/:
Rekickwasm is not a tool for editing the main process.
So when you edit the files in ./rekickwasm/edit/yaml/

In ./rekickwasm/edit/yaml/kickwasm.yaml
  * Do not edit title or importPath.
  * If you rename a panel rekickwasm will remove it and you will lose it. For example if you rename "MyOldPanel" to "MyNewPanel", rekickwasm will remove the "MyOldPanel" template file and go package files. Then rekickwasm will add a new template file and new go package files for "MyNewPanel"

  AND REMEMBER:
    * Panel names must be unique.
    * Button names must be unique.
    * Tab names must be unique.

In ./rekickwasm/edit/flags.yaml DO NOT EDIT
  * kwversionbreaking, kwversionfeature or kwversionpatch.

```
