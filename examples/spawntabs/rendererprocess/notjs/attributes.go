// +build js, wasm

package notjs

import (
	"fmt"
	"strconv"
	"syscall/js"
)

// TagName returns an element's tab name.
func (notjs *NotJS) TagName(el js.Value) string {
	return el.Get("tagName").String()
}

// ID returns an element's id.
func (notjs *NotJS) ID(el js.Value) string {
	return el.Get(idAttributeName).String()
}

// SetID sets an element's id.
func (notjs *NotJS) SetID(el js.Value, id string) {
	el.Set(idAttributeName, id)
}

// SetAttributeHref sets an element's href attribute
func (notjs *NotJS) SetAttributeHref(element js.Value, value string) {
	element.Call(setAttributeMethodName, "href", value)
}

// SetAttribute sets an element's attribute
func (notjs *NotJS) SetAttribute(element js.Value, name, value string) {
	element.Call(setAttributeMethodName, name, value)
}

// SetAttributeInt sets an element's attribute to an int value
func (notjs *NotJS) SetAttributeInt(element js.Value, name string, value int) {
	element.Call(setAttributeMethodName, name, fmt.Sprintf("%d", value))
}

// SetAttributeUint sets an element's attribute to an uint value
func (notjs *NotJS) SetAttributeUint(element js.Value, name string, value uint) {
	element.Call(setAttributeMethodName, name, fmt.Sprintf("%d", value))
}

// SetAttributeInt64 sets an element's attribute to an int64 value
func (notjs *NotJS) SetAttributeInt64(element js.Value, name string, value int64) {
	element.Call(setAttributeMethodName, name, fmt.Sprintf("%d", value))
}

// SetAttributeUint64 sets an element's attribute to an uint64 value
func (notjs *NotJS) SetAttributeUint64(element js.Value, name string, value uint64) {
	element.Call(setAttributeMethodName, name, fmt.Sprintf("%d", value))
}

// SetAttributeFloat64 sets an element's attribute to a float64 value
func (notjs *NotJS) SetAttributeFloat64(element js.Value, name string, value float64) {
	element.Call(setAttributeMethodName, name, fmt.Sprintf("%f", value))
}

// GetAttribute sets an element's attribute
func (notjs *NotJS) GetAttribute(element js.Value, name string) string {
	return element.Call(getAttributeMethodName, name).String()
}

// GetAttributeInt gets an element's attribute as an int value.
func (notjs *NotJS) GetAttributeInt(element js.Value, name string) int {
	s := element.Call(getAttributeMethodName, name).String()
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		i = 0
	}
	return int(i)
}

// GetAttributeUint gets an element's attribute as an uint value.
func (notjs *NotJS) GetAttributeUint(element js.Value, name string) uint {
	s := element.Call(getAttributeMethodName, name).String()
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		i = 0
	}
	return uint(i)
}

// GetAttributeInt64 gets an element's attribute as an int64 value.
func (notjs *NotJS) GetAttributeInt64(element js.Value, name string) int64 {
	s := element.Call(getAttributeMethodName, name).String()
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		i = 0
	}
	return int64(i)
}

// GetAttributeUint64 gets an element's attribute as an uint64 value.
func (notjs *NotJS) GetAttributeUint64(element js.Value, name string) uint64 {
	s := element.Call(getAttributeMethodName, name).String()
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		i = 0
	}
	return uint64(i)
}

// GetAttributeFloat64 gets an element's attribute as an float64 value.
func (notjs *NotJS) GetAttributeFloat64(element js.Value, name string) float64 {
	s := element.Call(getAttributeMethodName, name).String()
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		i = 0.0
	}
	return float64(i)
}
