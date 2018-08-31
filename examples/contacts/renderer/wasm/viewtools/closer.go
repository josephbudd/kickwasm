package viewtools

import (
	"syscall/js"
)

// Initialize inititializes the closer.
func (tools *Tools) initializeCloser() {
	notjs := tools.notjs
	// closer master view close button
	cb := notjs.RegisterCallBack(func([]js.Value) {
		tools.Quit()
	})
	button := notjs.GetElementByID("closerMasterView-close")
	notjs.SetOnClick(button, cb)
	// closer master view cancel button
	cb = notjs.RegisterCallBack(tools.toggleCloser)
	button = notjs.GetElementByID("closerMasterView-cancel")
	notjs.SetOnClick(button, cb)
}

// ToggleCloser toggles the closer master view.
func (tools *Tools) toggleCloser([]js.Value) {
	notjs := tools.notjs
	if !tools.ElementIsShown(tools.closerMasterView) {
		// closer view is not visible
		// so hide the other main div that is visible
		// and make the closer main div visible.
		children := notjs.ChildrenSlice(tools.body)
		for _, ch := range children {
			if notjs.TagName(ch) == "DIV" {
				if !notjs.ClassListContains(ch, UnSeenClassName) {
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
	tools.notjs.CloseCallBacks()
	tools.Global.Call("close")
}
