package viewtools

import (
	"syscall/js"
)

// Initialize inititializes the closer.
func (tools *Tools) initializeCloser() {
	notJS := tools.notJS
	// closer master view close button
	cb := notJS.RegisterCallBack(func([]js.Value) {
		tools.Quit()
	})
	button := notJS.GetElementByID("closerMasterView-close")
	notJS.SetOnClick(button, cb)
	// closer master view cancel button
	cb = notJS.RegisterCallBack(tools.toggleCloser)
	button = notJS.GetElementByID("closerMasterView-cancel")
	notJS.SetOnClick(button, cb)
}

// ToggleCloser toggles the closer master view.
func (tools *Tools) toggleCloser([]js.Value) {
	notJS := tools.notJS
	if !tools.ElementIsShown(tools.closerMasterView) {
		// closer view is not visible
		// so hide the other main div that is visible
		// and make the closer main div visible.
		children := notJS.ChildrenSlice(tools.body)
		for _, ch := range children {
			if notJS.TagName(ch) == "DIV" {
				if !notJS.ClassListContains(ch, UnSeenClassName) {
					// closer if visible so hide it
					tools.lastMasterView = ch
					tools.ElementHide(ch)
					break
				}
			}
		}
		// show the closer main div
		tools.ElementShow(tools.closerMasterView)
		return
	}
	// closer view is visible
	// so hide the closer view and show the last main div.
	tools.ElementHide(tools.closerMasterView)
	tools.ElementShow(tools.lastMasterView)
}

// Quit closes the application renderer.
func (tools *Tools) Quit() {
	tools.notJS.CloseCallBacks()
	tools.Global.Call("close")
}
