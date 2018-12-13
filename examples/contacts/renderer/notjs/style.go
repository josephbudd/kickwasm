package notjs

import (
	"fmt"
	"syscall/js"
)

// SetStyleHeight sets an element's style height.
func (notJS *NotJS) SetStyleHeight(element js.Value, height float64) {
	style := element.Get(styleMemberName)
	style.Set("height", fmt.Sprintf(pxFormatter, height))
}

// SetStyleWidth sets an element's style width.
func (notJS *NotJS) SetStyleWidth(element js.Value, width float64) {
	style := element.Get(styleMemberName)
	style.Set("width", fmt.Sprintf(pxFormatter, width))
}
