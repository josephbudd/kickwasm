// +build js, wasm

package viewtools

import (
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/event"
	"github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/window"
)

// SetPrintTitle sets the document title for printing.
func SetPrintTitle(title string) {
	printTitle = title
}

func beforePrint(e event.Event) (nilReturn interface{}) {
	document.Set("title", printTitle)
	ElementShow(blackMasterView)
	extraHeight = 0
	if !ElementIsShown(tabsMasterview) {
		// Master view is not visible there fore user content is not visilble.
		return
	}
	var userContentDiv js.Value
	var found bool
	if userContentDiv, found = findCurrentUserContent(); !found {
		// User content is not visible. Something else is.
		return
	}
	// User content is visible but is it all visible?
	scrollHeight := userContentDiv.Get("scrollHeight").Float()
	clientHeight := window.InnerHeight(userContentDiv)
	if scrollHeight <= clientHeight {
		// No need to resize because all of the user content is visible.
		return
	}
	// Must resize to make all of the user content visible.
	extraHeight = scrollHeight - clientHeight
	overSizeApp()
	return
}

func afterPrint(e event.Event) (nilReturn interface{}) {
	ElementHide(blackMasterView)
	// Restore the print title incase it was changed.
	printTitle = documentTitle
	document.Set("title", documentTitle)
	if extraHeight == 0 {
		// No need to resize back to the original size.
		return
	}
	SizeApp()
	return
}

func findCurrentUserContent() (div js.Value, found bool) {
	if div, found = findCurrentSliderPanel(); !found {
		return
	}
	id := div.Get("id").String() + "-inner-user-content"
	div = document.Call("getElementById", id)
	found = div != null
	return
}

func findCurrentSliderPanel() (div js.Value, found bool) {
	divs := document.Call("getElementsByTagName", "DIV")
	ldivs := divs.Length()
	for i := 0; i < ldivs; i++ {
		div = divs.Index(i)
		classList := div.Get("classList")
		if found = classList.Call("contains", SliderPanelClassName).Bool() && !classList.Call("contains", UnSeenClassName).Bool(); found {
			return
		}
	}
	return
}
