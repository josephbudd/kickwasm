// +build js, wasm

package viewtools

// ConsoleLog logs to the console.
func (tools *Tools) ConsoleLog(message string) {
	tools.console.Call("log", message)
}

// Alert inokes the browser's alert.
func (tools *Tools) Alert(message string) {
	tools.alert.Invoke(message)
}

// Success displays a message titled "Success"
func (tools *Tools) Success(message string) {
	tools.GoModal(message, "Success", nil)
}

// Error displays a message titled "Error"
func (tools *Tools) Error(message string) {
	tools.GoModal(message, "Error", nil)
}
