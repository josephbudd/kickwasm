package notjs

// ConsoleLog logs to the console.
func (notJS *NotJS) ConsoleLog(message string) {
	notJS.console.Call("log", message)
}

// Alert invokes the browser's alert.
func (notJS *NotJS) Alert(message string) {
	notJS.alert.Invoke(message)
}
