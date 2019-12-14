// +build js, wasm

package markup

import (
	"github.com/josephbudd/kickwasm/examples/push/rendererprocess/api/window"
)

// InnerWidth returns the innermost width.
func (e *Element) InnerWidth() (width float64) {
	width = window.InnerWidth(e.element)
	return
}

// InnerHeight returns the innermost height.
func (e *Element) InnerHeight() (height float64) {
	height = window.InnerHeight(e.element)
	return
}

// OuterWidth returns the total width.
func (e *Element) OuterWidth() (width float64) {
	width = window.OuterWidth(e.element)
	return
}

// OuterHeight returns the total height.
func (e *Element) OuterHeight() (height float64) {
	height = window.OuterHeight(e.element)
	return
}

// OutlineWidth return the total outline width.
func (e *Element) OutlineWidth() (width float64) {
	width = window.OutlineWidth(e.element)
	return
}

// WidthExtras returns the total width that is not the innermost width.
func (e *Element) WidthExtras() (width float64) {
	width = window.WidthExtras(e.element)
	return
}

// HeightExtras returns the total height that is not the innermost height.
func (e *Element) HeightExtras() (height float64) {
	height = window.HeightExtras(e.element)
	return
}

// PaddingWidth returns the total padding width.
func (e *Element) PaddingWidth() (width float64) {
	width = window.PaddingWidth(e.element)
	return
}

// PaddingHeight returns the total padding height.
func (e *Element) PaddingHeight() (height float64) {
	height = window.PaddingHeight(e.element)
	return
}

// MarginWidth returns the total margin width.
func (e *Element) MarginWidth() (width float64) {
	width = window.MarginWidth(e.element)
	return
}

// MarginHeight returns the total margin height.
func (e *Element) MarginHeight() (height float64) {
	height = window.MarginHeight(e.element)
	return
}

// BorderWidth returns the total border width.
func (e *Element) BorderWidth() (width float64) {
	width = window.BorderWidth(e.element)
	return
}

// BorderHeight returns the total border height.
func (e *Element) BorderHeight() (height float64) {
	height = window.BorderHeight(e.element)
	return
}

// Style

// SetStyleHeight sets an element's style height.
func (e *Element) SetStyleHeight(height float64) {
	window.SetStyleHeight(e.element, height)
}

// SetStyleWidth sets an element's style width.
func (e *Element) SetStyleWidth(width float64) {
	window.SetStyleWidth(e.element, width)
}

// Resizing to full width.

// SetResizeable allows the block element to be resized to fit perfectly inside it's parent's width.
func (e *Element) SetResizeable() {
	e.AddClass(resizeMeWidthClassName)
}
