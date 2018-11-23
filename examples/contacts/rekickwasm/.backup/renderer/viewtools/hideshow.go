package viewtools

import (
	"syscall/js"
)

// IDIsShown returns is the element with the id is shown.
func (tools *Tools) IDIsShown(id string) bool {
	notJS := tools.notJS
	element := notJS.GetElementByID(id)
	unseen := notJS.ClassListContains(element, UnSeenClassName)
	return !unseen
}

// ElementIsShown returns if the element is shown.
func (tools *Tools) ElementIsShown(element js.Value) bool {
	unseen := tools.notJS.ClassListContains(element, UnSeenClassName)
	return !unseen
}

// IDShow shows the element with the id.
func (tools *Tools) IDShow(id string) {
	tools.elementShow(tools.notJS.GetElementByID(id), id)
}

// ElementShow shows the element.
func (tools *Tools) ElementShow(element js.Value) {
	tools.elementShow(element, tools.notJS.ID(element))
}

func (tools *Tools) elementShow(element js.Value, id string) {
	if id != "tabsMasterView-home-slider-collection" && id != "tabsMasterView-home" && id != "closerMasterView" {
		isSlider, _ := tools.toBeShownInGroup(element)
		if isSlider {
			return
		}
	}
	tools.notJS.ClassListReplaceClass(element, UnSeenClassName, SeenClassName)
}

// IDHide hides the element with the id.
func (tools *Tools) IDHide(id string) {
	tools.elementHide(tools.notJS.GetElementByID(id), id)
}

// ElementHide hides the element.
func (tools *Tools) ElementHide(element js.Value) {
	tools.elementHide(element, tools.notJS.ID(element))
}

func (tools *Tools) elementHide(element js.Value, id string) {
	notJS := tools.notJS
	notJS.ConsoleLog("hiding #" + id)
	if id != "tabsMasterView-home-slider-collection" && id != "tabsMasterView-home" && id != "closerMasterView" && tools.toBeHiddenInGroup(element) {
		return
	}
	notJS.ClassListReplaceClass(element, SeenClassName, UnSeenClassName)
}
