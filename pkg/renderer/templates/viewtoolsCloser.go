package templates

// ViewToolsCloser is the rendererprocess/viewtools/closer.go file.
const ViewToolsCloser = `// +build js, wasm

package viewtools

import (
	"{{.ApplicationGitPath}}{{.ImportRendererFrameworkCallBack}}"
)

// Quit closes the application renderer.
func Quit() {
	callback.CloseCallBacks()
	global.Call("close")
}

// quit closes the application renderer.
func quit() {
	callback.CloseCallBacks()
	global.Call("close")
}
`
