package templates

// MarkupHideShowGo is the rendererprocess/markup/hideshow.go file.
const MarkupHideShowGo = `// +build js, wasm

package markup

import (
	"{{.ApplicationGitPath}}{{.ImportRendererFrameworkViewTools}}"
)

// Show makes the element visible.
func (e *Element) Show() {
	viewtools.ElementShow(e.element)
}

// Hide makes the element not visible.
func (e *Element) Hide() {
	viewtools.ElementHide(e.element)
}

// IsShown returns is the element is visible.
func (e *Element) IsShown() (isShown bool) {
	isShown = viewtools.ElementIsShown(e.element)
	return
}
`
