package templates

// JSValueHideShowGo is the rendererprocess/api/jsvalue/hideshow.go file template.
const JSValueHideShowGo = `// +build js, wasm

package jsvalue

import (
	"syscall/js"

	"{{.ApplicationGitPath}}{{.ImportRendererFrameworkViewTools}}"
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
`
