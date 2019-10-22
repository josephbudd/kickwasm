// +build js, wasm

package viewtools

// LockButtons blocks the tab and back buttons from working.
func (tools *Tools) LockButtons() {
	tools.buttonsLocked = true
	tools.buttonsLockedMessageTitle = ""
	tools.buttonsLockedMessageText = ""
}

// LockButtonsWithMessage blocks the tab and back buttons from working.
// It also displays a message to the user when the user clicks a tab or back button.
func (tools *Tools) LockButtonsWithMessage(message, title string) {
	tools.buttonsLocked = true
	tools.buttonsLockedMessageTitle = title
	tools.buttonsLockedMessageText = message
}

// UnLockButtons lets tab and back buttons work.
func (tools *Tools) UnLockButtons() {
	tools.buttonsLocked = false
}

// HandleButtonClick takes care of a button click returning if the button is locked.
// Returns if the button clicked.
func (tools *Tools) HandleButtonClick() (clicked bool) {
	if tools.buttonsLocked {
		if len(tools.buttonsLockedMessageText) > 0 {
			tools.GoModal(tools.buttonsLockedMessageText, tools.buttonsLockedMessageTitle, nil)
		}
	}
	clicked = !tools.buttonsLocked
	return
}
