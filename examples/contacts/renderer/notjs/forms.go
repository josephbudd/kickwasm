package notjs

import (
	"strconv"
	"syscall/js"
)

// ClearValue sets an element's value
func (notJS *NotJS) ClearValue(element js.Value) {
	element.Set(valueAttributeName, "")
}

// SetValue sets an element's value
func (notJS *NotJS) SetValue(element js.Value, value string) {
	element.Set(valueAttributeName, value)
}

// GetValue gets an element's value as a string.
func (notJS *NotJS) GetValue(element js.Value) string {
	return element.Get(valueAttributeName).String()
}

// GetValueInt gets an element's value as an int.
func (notJS *NotJS) GetValueInt(element js.Value) int {
	v := element.Get(valueAttributeName).String()
	n, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return int(0)
	}
	return int(n)
}

// GetValueInt64 gets an element's value as an int64.
func (notJS *NotJS) GetValueInt64(element js.Value) int64 {
	v := element.Get(valueAttributeName).String()
	n, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return int64(0)
	}
	return int64(n)
}

// GetValueUint gets an element's value as an uint.
func (notJS *NotJS) GetValueUint(element js.Value) uint {
	v := element.Get(valueAttributeName).String()
	n, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return uint(0)
	}
	return uint(n)
}

// GetValueUint64 gets an element's value as an uint64.
func (notJS *NotJS) GetValueUint64(element js.Value) uint64 {
	v := element.Get(valueAttributeName).String()
	n, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return uint64(0)
	}
	return uint64(n)
}

// GetValueFloat64 gets an element's value
func (notJS *NotJS) GetValueFloat64(element js.Value) float64 {
	v := element.Get(valueAttributeName).String()
	n, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return float64(0)
	}
	return n
}

// GetChecked gets an element's checked
func (notJS *NotJS) GetChecked(element js.Value) bool {
	return element.Get(checkedAttributeName).Bool()
}

// SetChecked gets an element's checked
func (notJS *NotJS) SetChecked(element js.Value, checked bool) {
	element.Set(checkedAttributeName, checked)
}

// Focus sets an element to focused.
func (notJS *NotJS) Focus(element js.Value) {
	element.Call("focus")
}

// Blur removes an element's focus.
func (notJS *NotJS) Blur(element js.Value) {
	element.Call("blur")
}
