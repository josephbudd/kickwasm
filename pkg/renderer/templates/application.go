package templates

// ApplicationGo is the rendererprocess/application/application.go file.
const ApplicationGo = `// +build js, wasm

package application

import (
	"context"

	"{{.ApplicationGitPath}}{{.ImportRendererAPIDisplay}}"
	"{{.ApplicationGitPath}}{{.ImportRendererAPIEvent}}"
	"{{.ApplicationGitPath}}{{.ImportRendererFrameworkCallBack}}"
)

// NewGracefullyCloseHandler makes an event handler which gracefully closes the application for you.
// Use it in your panel controllers to handle your own application closing buttons.
// Param cancelFunc is the renderer process's context cancel func.
func NewGracefullyCloseHandler(cancelFunc context.CancelFunc) (handler func(e event.Event) (nilReturn interface{})) {
	handler = func(e event.Event) (nilReturn interface{}) {
		GracefullyClose(cancelFunc)
		return
	}
	return
}

// GracefullyClose gracefully closes the application for you.
// Use it in your panel controllers.
// Param cancelFunc is the renderer process's context cancel func.
func GracefullyClose(cancelFunc context.CancelFunc) {
	callback.RemoveApplicationOnCloseHandler()
	title := "Closing"
	msg := "Closing <q>{{.Title}}</q>."
	display.Inform(msg, title, cancelFunc)
	return
}
`
