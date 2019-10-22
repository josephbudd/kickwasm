# Notes on rekickwasm

## Tests

1. ok: Undo.
1. ok: Move Homes.
1. ok: Swap Home Buttons.
1. ok: Swap Home Button Panels.
1. ok: Add About Home with tabs.
1. ok: Rotate tabs inside About Home.
1. ok: Add another tab inside About Home.
1. ok: Swap tabs between totally separate panels.
1. ok: Add apanel to agroup.
1. ok: Swap a panel in a group.
1. ok: Rename a panel in a group.

## Don't allow

* move a button to a panel which already has a button with the same name.
* move a tab to a panel which already has a button with the same name.

## must not be changed in /rendererprocess/

* calls/
* panelHelping/
* interfaces/
* viewtools/
* app.wasm
* main.go
* wasm_exec.js

## must be changed in /rendererprocess/

* css/
* panels/
* templates/
* panels.go
