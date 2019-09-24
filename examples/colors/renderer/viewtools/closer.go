package viewtools

// Initialize inititializes the closer.
func (tools *Tools) initializeCloser() {
	notJS := tools.NotJS
	// closer master view close button
	f := func(e Event) (nilReturn interface{}) {
		tools.Quit()
		return
	}
	button := notJS.GetElementByID("closerMasterView-close")
	tools.AddEventHandler(f, button, "click", false)
	// closer master view cancel button
	button = notJS.GetElementByID("closerMasterView-cancel")
	tools.AddEventHandler(tools.toggleCloser, button, "click", false)
}

// ToggleCloser toggles the closer master view.
func (tools *Tools) toggleCloser(e Event) (nilReturn interface{}) {
	notJS := tools.NotJS
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
	return
}

// Quit closes the application renderer.
func (tools *Tools) Quit() {
	tools.CloseCallBacks()
	tools.Global.Call("close")
}
