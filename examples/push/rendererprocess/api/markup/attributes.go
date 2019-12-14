// +build js, wasm

package markup

import (
	"fmt"
	"strconv"
)

// SetAttribute sets the element's attribute.
// Param name is the name of the attribute.
// Param value is the attributes setting. It must not be a pointer.
func (e *Element) SetAttribute(name string, value interface{}) {
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
	e.element.Call(setAttributeMethodName, name, v)
}


// Attribute returns an element's attribute value as a string.
func (e *Element) Attribute(name string) (value string) {
	value = e.element.Call(getAttributeMethodName, name).String()
	return
}

// AttributeInt64 returns an element's attribute value as an int64 and the error.
// If there is an error the value is 0.
func (e *Element) AttributeInt64(name string) (value int64, err error) {
	s := e.element.Call(getAttributeMethodName, name).String()
	if value, err = strconv.ParseInt(s, 10, 64); err != nil {
		value = 0
	}
	return
}

// AttributeUint64 returns an element's attribute value as an uint64 and the error.
// If there is an error the value is 0.
func (e *Element) AttributeUint64(name string) (value uint64, err error) {
	s := e.element.Call(getAttributeMethodName, name).String()
	if value, err = strconv.ParseUint(s, 10, 64); err != nil {
		value = 0
	}
	return
}

// AttributeFloat64 returns an element's attribute value as a float64 and the error.
// If there is an error the value is 0.0.
func (e *Element) AttributeFloat64(name string) (value float64, err error) {
	s := e.element.Call(getAttributeMethodName, name).String()
	if value, err = strconv.ParseFloat(s, 64); err != nil {
		value = 0.0
	}
	return
}
