// +build js, wasm

package viewtools

// Alert inokes the browser's alert.
func Alert(message string) {
	alert.Invoke(message)
}

// Success displays a message titled "Success"
func Success(message string) {
	GoModal(message, "Success", nil)
}

// Error displays a message titled "Error"
func Error(message string) {
	GoModal(message, "Error", nil)
}
