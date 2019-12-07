// +build js, wasm

package display

import (
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/push/rendererprocess/framework/location"
	"github.com/josephbudd/kickwasm/examples/push/rendererprocess/framework/viewtools"
)

var printFunc js.Value

func init() {
	printFunc = js.Global().Get("print")
}

// Alert invokes the browser's alert.
func Alert(message string) {
	viewtools.Alert(message)
}

// Success displays a message titled "Success" using the display's overlay.
func Success(message string) {
	viewtools.GoModal(message, "Success", nil)
}

// Error displays a message titled "Error" using the display's overlay.
func Error(message string) {
	viewtools.GoModal(message, "Error", nil)
}

// Inform displays a title and message using the display's overlay.
func Inform(message, title string, callback func()) {
	viewtools.GoModal(message, title, callback)
}

// InformHTML displays a title and html message using the display's overlay.
// Param htmlMessage is html.
// Param title is plain text.
func InformHTML(htmlMessage, title string, callback func()) {
	viewtools.GoModalHTML(htmlMessage, title, callback)
}

// Back simulates a click on the tall back button at the left of slider panels.
func Back() {
	viewtools.Back()
}

// ForceTabButtonClick implements the behavior of a tab button being clicked by the user.
func ForceTabButtonClick(button js.Value) {
	viewtools.ForceTabButtonClick(button)
}

// HostPort returns the document location host and port.
func HostPort() (host string, port uint64) {
	host, port = location.HostPort()
	return
}

// BlockButtons blocks the tab and back buttons from working.
func BlockButtons() {
	viewtools.LockButtons()
	return
}

// BlockButtonsWithMessage blocks the tab and back buttons from working.
// It also displays a message to the user when the user clicks a tab or back button.
func BlockButtonsWithMessage(message, title string) {
	viewtools.LockButtonsWithMessage(message, title)
}

// UnBlockButtons lets tab and back buttons work again.
func UnBlockButtons() {
	viewtools.UnLockButtons()
}

// Resize resizes the entire GUI layout to fit the window.
func Resize() {
	viewtools.SizeApp()
}

// SpawnID
func SpawnID(brokenID string, spawnPanelID uint64) (fixedID string) {
	fixedID = viewtools.FixSpawnID(brokenID, spawnPanelID)
	return
}

// NewSpawnWidgetUniqueID returns a new id for a widget in a spawned panel.
func NewSpawnWidgetUniqueID() (spawnWidgetID uint64) {
	spawnWidgetID = viewtools.NewSpawnWidgetUniqueID()
	return
}

// SpawnWidget spawns a widget.
func SpawnWidget(spawnWidgetID uint64, widget, parent js.Value) {
	viewtools.SpawnWidget(spawnWidgetID, widget, parent)
}

// UnSpawnWidget unspawns a widget.
func UnSpawnWidget(spawnWidgetID uint64) {
	viewtools.UnSpawnWidget(spawnWidgetID)
}

// Print prints the appliction to the printer.
// Param title is the unique title for the printed page.
// The styles for printing are in site/mycss/Usercontent.css.
// Use it to print your markup panels.
func Print(title string) {
	viewtools.SetPrintTitle(title)
	printFunc.Invoke()
}
