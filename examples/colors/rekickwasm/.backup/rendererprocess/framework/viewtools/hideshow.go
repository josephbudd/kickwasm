// +build js, wasm

package viewtools

import (
	"log"
	"syscall/js"
)

// IDIsShown returns is the element with the id is shown.
func IDIsShown(id string) (isShown bool) {
	e := getElementByID(document, id)
	classList := e.Get("classList")
	isShown = !classList.Call("contains", UnSeenClassName).Bool()
	return
}

// ElementIsShown returns if the element is shown.
func ElementIsShown(e js.Value) (isShown bool) {
	classList := e.Get("classList")
	isShown = !classList.Call("contains", UnSeenClassName).Bool()
	return
}

// IDShow shows the element with the id.
func IDShow(id string) {
	elementShow(getElementByID(document, id), id)
}

// ElementShow shows the element.
func ElementShow(e js.Value) {
	elementShow(e, e.Get("id").String())
}

func elementShow(e js.Value, id string) {
	if id != "mainMasterView-home-slider-collection" && id != "mainMasterView-home" {
		isSlider, _ := ShowInGroup(e, ToBeSeenClassName, ToBeUnSeenClassName)
		if isSlider {
			return
		}
	}
	classList := e.Get("classList")
	if classList.Call("replace", UnSeenClassName, SeenClassName).Bool() {
		return
	}
	if !classList.Call("contains", SeenClassName).Bool() {
		classList.Call("add", SeenClassName)
	}
}

// IDHide hides the element with the id.
func IDHide(id string) {
	elementHide(getElementByID(document, id), id)
}

// ElementHide hides the element.
func ElementHide(e js.Value) {
	elementHide(e, e.Get("id").String())
}

func elementHide(e js.Value, id string) {
	log.Println("hiding #" + id)
	if id != "mainMasterView-home-slider-collection" && id != "mainMasterView-home" && hideInGroup(e, ToBeSeenClassName, ToBeUnSeenClassName) {
		return
	}
	classList := e.Get("classList")
	if classList.Call("replace", SeenClassName, UnSeenClassName).Bool() {
		return
	}
	if !classList.Call("contains", UnSeenClassName).Bool() {
		classList.Call("add", UnSeenClassName)
	}
}
