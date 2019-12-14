// +build js, wasm

package viewtools

// LockButtons blocks the tab and back buttons from working.
func LockButtons() {
	buttonsLocked = true
	buttonsLockedMessageTitle = ""
	buttonsLockedMessageText = ""
}

// LockButtonsWithMessage blocks the tab and back buttons from working.
// It also displays a message to the user when the user clicks a tab or back button.
func LockButtonsWithMessage(message, title string) {
	buttonsLocked = true
	buttonsLockedMessageTitle = title
	buttonsLockedMessageText = message
}

// UnLockButtons lets tab and back buttons work.
func UnLockButtons() {
	buttonsLocked = false
}

// HandleButtonClick takes care of a button click returning if the button is locked.
// Returns if the button clicked.
func HandleButtonClick() (clicked bool) {
	if buttonsLocked {
		if len(buttonsLockedMessageText) > 0 {
			GoModal(buttonsLockedMessageText, buttonsLockedMessageTitle, nil)
		}
	}
	clicked = !buttonsLocked
	return
}
