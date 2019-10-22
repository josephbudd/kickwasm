// +build js, wasm

package notjs

// ConsoleLog logs to the console.
func (notjs *NotJS) ConsoleLog(message string) {
	notjs.console.Call("log", message)
}

// Alert invokes the browser's alert.
func (notjs *NotJS) Alert(message string) {
	notjs.alert.Invoke(message)
}
