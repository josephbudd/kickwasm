// +build js, wasm

package jsvalue

import (
	"syscall/js"
)

// SetNotPrintable sets the element to not be printed.
func SetNotPrintable(e js.Value) {
	classList := e.Get("classList")
	classList.Call("add", "do-not-print")
}
