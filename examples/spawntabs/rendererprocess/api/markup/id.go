// +build js, wasm

package markup

import (
	"fmt"
)

// SetID sets the element's id.
// Param value is the attributes setting. It must not be a pointer.
func (e *Element) SetID(value interface{}) {
	var v string
	switch value := value.(type) {
	case string:
		v = value
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		v = fmt.Sprintf("%d", value)
	case float32, float64:
		v = fmt.Sprintf("%f", value)
	default:
		v = ""
	}
	e.element.Set(idAttributeName, v)
}

// ID returns the element's id.
func (e *Element) ID() (id string) {
	id = e.element.Get(idAttributeName).String()
	return
}
