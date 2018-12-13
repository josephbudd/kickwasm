package notjs

import (
	"strconv"
	"syscall/js"
)

// WindowInnerWidth returns the window's inner width.
func (notJS *NotJS) WindowInnerWidth() float64 {
	return notJS.global.Get("innerWidth").Float()
}

// WindowInnerHeight returns the window's inner height.
func (notJS *NotJS) WindowInnerHeight() float64 {
	return notJS.global.Get("innerHeight").Float()
}

// InnerWidth returns the innermost width.
func (notJS *NotJS) InnerWidth(el js.Value) float64 {
	// offset - left, right padding and border
	styles := notJS.getComputedStyle(el)
	px := styles.Get("width").String()
	f, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	return f
}

// InnerHeight returns the innermost height.
func (notJS *NotJS) InnerHeight(el js.Value) float64 {
	// offset - top, bottom padding and border
	styles := notJS.getComputedStyle(el)
	px := styles.Get("height").String()
	f, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	return f
}

// OuterWidth returns the total width.
func (notJS *NotJS) OuterWidth(el js.Value) float64 {
	// offset + left, right margin
	ofw := el.Get("offsetWidth").Float()
	styles := notJS.getComputedStyle(el)
	px := styles.Get("marginLeft").String()
	ml, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("marginRight").String()
	mr, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("outlineWidth").String()
	olw, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	return ofw + ml + mr + olw + olw
}

// OuterHeight returns the total height.
func (notJS *NotJS) OuterHeight(el js.Value) float64 {
	// offset + top, bottom margin
	ofh := el.Get("offsetHeight").Float()
	styles := notJS.getComputedStyle(el)
	px := styles.Get("marginTop").String()
	mt, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("marginBottom").String()
	mb, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("outlineWidth").String()
	olw, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	return ofh + mt + mb + olw + olw
}

// OutlineWidth return the total outline width.
func (notJS *NotJS) OutlineWidth(el js.Value) float64 {
	styles := notJS.getComputedStyle(el)
	px := styles.Get("outlineWidth").String()
	f, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	return f
}

// WidthExtras returns the total width that is not the innermost width.
func (notJS *NotJS) WidthExtras(el js.Value) float64 {
	b := notJS.BorderWidth(el)
	p := notJS.PaddingWidth(el)
	m := notJS.MarginWidth(el)
	o := notJS.OutlineWidth(el)
	return b + p + m + o + o
}

// HeightExtras returns the total height that is not the innermost height.
func (notJS *NotJS) HeightExtras(el js.Value) float64 {
	b := notJS.BorderHeight(el)
	p := notJS.PaddingHeight(el)
	m := notJS.MarginHeight(el)
	o := notJS.OutlineWidth(el)
	return b + p + m + o + o
}

// PaddingWidth returns the total padding width.
func (notJS *NotJS) PaddingWidth(el js.Value) float64 {
	styles := notJS.getComputedStyle(el)
	px := styles.Get("paddingLeft").String()
	pl, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("paddingRight").String()
	pr, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	return pl + pr
}

// PaddingHeight returns the total padding height.
func (notJS *NotJS) PaddingHeight(el js.Value) float64 {
	styles := notJS.getComputedStyle(el)
	px := styles.Get("paddingTop").String()
	pt, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("paddingBottom").String()
	pb, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	return pt + pb
}

// MarginWidth returns the total margin width.
func (notJS *NotJS) MarginWidth(el js.Value) float64 {
	styles := notJS.getComputedStyle(el)
	px := styles.Get("marginLeft").String()
	ml, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("marginRight").String()
	mr, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	return ml + mr
}

// MarginHeight returns the total margin height.
func (notJS *NotJS) MarginHeight(el js.Value) float64 {
	styles := notJS.getComputedStyle(el)
	px := styles.Get("marginTop").String()
	mt, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("marginBottom").String()
	mb, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	return mt + mb
}

// BorderWidth returns the total border width.
func (notJS *NotJS) BorderWidth(el js.Value) float64 {
	styles := notJS.getComputedStyle(el)
	px := styles.Get("borderLeftWidth").String()
	blw, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("borderRightWidth").String()
	brw, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	return blw + brw
}

// BorderHeight returns the total border height.
func (notJS *NotJS) BorderHeight(el js.Value) float64 {
	styles := notJS.getComputedStyle(el)
	px := styles.Get("borderTopWidth").String()
	btw, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("borderBottomWidth").String()
	bbw, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	return btw + bbw
}

func (notJS *NotJS) getComputedStyle(el js.Value) js.Value {
	return notJS.global.Call("getComputedStyle", el)
}
