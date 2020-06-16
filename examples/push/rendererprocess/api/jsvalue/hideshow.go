// +build js, wasm

package jsvalue

import (
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/push/rendererprocess/framework/viewtools"
)

// Show makes the element visible.
func Show(e js.Value) {
	viewtools.ElementShow(e)
}

// Hide makes the element not visible.
func Hide(e js.Value) {
	viewtools.ElementHide(e)
}

// IsShown returns if the element is visible.
func IsShown(e js.Value) (isShown bool) {
	isShown = viewtools.ElementIsShown(e)
	return
}
