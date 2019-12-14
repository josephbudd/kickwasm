// +build js, wasm

package markup

import (
	"fmt"
	"strconv"
)

// ClearValue sets an element's value
func (e *Element) ClearValue() {
	e.element.Set(valueAttributeName, "")
}

// SetValue sets an element's value
func (e *Element) SetValue(value interface{}) {
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
	e.element.Set(valueAttributeName, v)
}

// Value gets an element's value as a string.
// Value
func (e *Element) Value() (value string) {
	value = e.element.Get(valueAttributeName).String()
	return
}

// ValueInt64 gets an element's value as an int64.
// If there is an error the value is 0.
func (e *Element) ValueInt64() (value int64, err error) {
	s := e.element.Get(valueAttributeName).String()
	value, err = strconv.ParseInt(s, 10, 64)
	if err != nil {
		value = 0
	}
	return
}

// ValueUint64 gets an element's value as an uint64.
// If there is an error the value is 0.
func (e *Element) ValueUint64() (value uint64, err error) {
	s := e.element.Get(valueAttributeName).String()
	if value, err = strconv.ParseUint(s, 10, 64); err != nil {
		value = 0
	}
	return
}

// ValueFloat64 gets an element's value as a float64.
// If there is an error the value is 0.0.
func (e *Element) ValueFloat64() (value float64, err error) {
	s := e.element.Get(valueAttributeName).String()
	if value, err = strconv.ParseFloat(s, 64); err != nil {
		value = 0.0
	}
	return
}
