package templates

// ViewToolsHideShow is the renderer/viewtools/hideshow.go file.
const ViewToolsHideShow = `package viewtools

import (
	"syscall/js"
)

// IDIsShown returns is the element with the id is shown.
func (tools *Tools) IDIsShown(id string) bool {
	notjs := tools.notjs
	element := notjs.GetElementByID(id)
	unseen := notjs.ClassListContains(element, UnSeenClassName)
	return !unseen
}

// ElementIsShown returns if the element is shown.
func (tools *Tools) ElementIsShown(element js.Value) bool {
	unseen := tools.notjs.ClassListContains(element, UnSeenClassName)
	return !unseen
}

// IDShow shows the element with the id.
func (tools *Tools) IDShow(id string) {
	tools.elementShow(tools.notjs.GetElementByID(id), id)
}

// ElementShow shows the element.
func (tools *Tools) ElementShow(element js.Value) {
	tools.elementShow(element, tools.notjs.ID(element))
}

func (tools *Tools) elementShow(element js.Value, id string) {
	if id != "tabsMasterView-home-slider-collection" && id != "tabsMasterView-home" && id != "closerMasterView" {
		isSlider, _ := tools.toBeShownInGroup(element)
		if isSlider {
			return
		}
	}
	tools.notjs.ClassListReplaceClass(element, UnSeenClassName, SeenClassName)
}

// IDHide hides the element with the id.
func (tools *Tools) IDHide(id string) {
	tools.elementHide(tools.notjs.GetElementByID(id), id)
}

// ElementHide hides the element.
func (tools *Tools) ElementHide(element js.Value) {
	tools.elementHide(element, tools.notjs.ID(element))
}

func (tools *Tools) elementHide(element js.Value, id string) {
	notjs := tools.notjs
	notjs.ConsoleLog("hiding #" + id)
	if id != "tabsMasterView-home-slider-collection" && id != "tabsMasterView-home" && id != "closerMasterView" && tools.toBeHiddenInGroup(element) {
		return
	}
	notjs.ClassListReplaceClass(element, SeenClassName, UnSeenClassName)
}
`
