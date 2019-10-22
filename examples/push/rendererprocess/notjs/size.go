// +build js, wasm

package notjs

import (
	"strconv"
	"syscall/js"
)

// WindowInnerWidth returns the window's inner width.
func (notjs *NotJS) WindowInnerWidth() float64 {
	return notjs.global.Get("innerWidth").Float()
}

// WindowInnerHeight returns the window's inner height.
func (notjs *NotJS) WindowInnerHeight() float64 {
	return notjs.global.Get("innerHeight").Float()
}

// InnerWidth returns the innermost width.
func (notjs *NotJS) InnerWidth(el js.Value) float64 {
	// offset - left, right padding and border
	styles := notjs.getComputedStyle(el)
	px := styles.Get("width").String()
	f, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	return f
}

// InnerHeight returns the innermost height.
func (notjs *NotJS) InnerHeight(el js.Value) float64 {
	// offset - top, bottom padding and border
	styles := notjs.getComputedStyle(el)
	px := styles.Get("height").String()
	f, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	return f
}

// OuterWidth returns the total width.
func (notjs *NotJS) OuterWidth(el js.Value) float64 {
	// offset + left, right margin
	ofw := el.Get("offsetWidth").Float()
	styles := notjs.getComputedStyle(el)
	px := styles.Get("marginLeft").String()
	ml, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("marginRight").String()
	mr, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("outlineWidth").String()
	olw, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	return ofw + ml + mr + olw + olw
}

// OuterHeight returns the total height.
func (notjs *NotJS) OuterHeight(el js.Value) float64 {
	// offset + top, bottom margin
	ofh := el.Get("offsetHeight").Float()
	styles := notjs.getComputedStyle(el)
	px := styles.Get("marginTop").String()
	mt, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("marginBottom").String()
	mb, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("outlineWidth").String()
	olw, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	return ofh + mt + mb + olw + olw
}

// OutlineWidth return the total outline width.
func (notjs *NotJS) OutlineWidth(el js.Value) float64 {
	styles := notjs.getComputedStyle(el)
	px := styles.Get("outlineWidth").String()
	f, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	return f
}

// WidthExtras returns the total width that is not the innermost width.
func (notjs *NotJS) WidthExtras(el js.Value) float64 {
	b := notjs.BorderWidth(el)
	p := notjs.PaddingWidth(el)
	m := notjs.MarginWidth(el)
	o := notjs.OutlineWidth(el)
	return b + p + m + o + o
}

// HeightExtras returns the total height that is not the innermost height.
func (notjs *NotJS) HeightExtras(el js.Value) float64 {
	b := notjs.BorderHeight(el)
	p := notjs.PaddingHeight(el)
	m := notjs.MarginHeight(el)
	o := notjs.OutlineWidth(el)
	return b + p + m + o + o
}

// PaddingWidth returns the total padding width.
func (notjs *NotJS) PaddingWidth(el js.Value) float64 {
	styles := notjs.getComputedStyle(el)
	px := styles.Get("paddingLeft").String()
	pl, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("paddingRight").String()
	pr, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	return pl + pr
}

// PaddingHeight returns the total padding height.
func (notjs *NotJS) PaddingHeight(el js.Value) float64 {
	styles := notjs.getComputedStyle(el)
	px := styles.Get("paddingTop").String()
	pt, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("paddingBottom").String()
	pb, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	return pt + pb
}

// MarginWidth returns the total margin width.
func (notjs *NotJS) MarginWidth(el js.Value) float64 {
	styles := notjs.getComputedStyle(el)
	px := styles.Get("marginLeft").String()
	ml, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("marginRight").String()
	mr, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	return ml + mr
}

// MarginHeight returns the total margin height.
func (notjs *NotJS) MarginHeight(el js.Value) float64 {
	styles := notjs.getComputedStyle(el)
	px := styles.Get("marginTop").String()
	mt, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("marginBottom").String()
	mb, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	return mt + mb
}

// BorderWidth returns the total border width.
func (notjs *NotJS) BorderWidth(el js.Value) float64 {
	styles := notjs.getComputedStyle(el)
	px := styles.Get("borderLeftWidth").String()
	blw, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("borderRightWidth").String()
	brw, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	return blw + brw
}

// BorderHeight returns the total border height.
func (notjs *NotJS) BorderHeight(el js.Value) float64 {
	styles := notjs.getComputedStyle(el)
	px := styles.Get("borderTopWidth").String()
	btw, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("borderBottomWidth").String()
	bbw, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	return btw + bbw
}

func (notjs *NotJS) getComputedStyle(el js.Value) js.Value {
	return notjs.global.Call("getComputedStyle", el)
}
