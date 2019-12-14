// +build js, wasm

package window

import (
	"fmt"
	"strconv"
	"syscall/js"
)

var (
	global js.Value
)

const (
	pxFormatter     = "%fpx"
	styleMemberName = "style"
)

func init() {
	global = js.Global()
}

// WindowInnerWidth returns the browser window's inner width.
func WindowInnerWidth() (width float64) {
	width = global.Get("innerWidth").Float()
	return
}

// WindowInnerHeight returns the browser window's inner height.
func WindowInnerHeight() (height float64) {
	height = global.Get("innerHeight").Float()
	return
}

// InnerWidth returns the block element's innermost width.
func InnerWidth(e js.Value) (width float64) {
	// offset - left, right padding and border
	styles := global.Call("getComputedStyle", e)
	px := styles.Get("width").String()
	width, _ = strconv.ParseFloat(px[:len(px)-2], 64)
	return
}

// InnerHeight returns the block element's innermost height.
func InnerHeight(e js.Value) (height float64) {
	// offset - top, bottom padding and border
	styles := global.Call("getComputedStyle", e)
	px := styles.Get("height").String()
	height, _ = strconv.ParseFloat(px[:len(px)-2], 64)
	return
}

// OuterWidth returns the block element's outermost width.
func OuterWidth(e js.Value) (width float64) {
	// offset + left, right margin
	ofw := e.Get("offsetWidth").Float()
	styles := global.Call("getComputedStyle", e)
	px := styles.Get("marginLeft").String()
	ml, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("marginRight").String()
	mr, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("outlineWidth").String()
	olw, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	width = ofw + ml + mr + olw + olw
	return
}

// OuterHeight returns the block element's outermost height.
func OuterHeight(e js.Value) (height float64) {
	// offset + top, bottom margin
	ofh := e.Get("offsetHeight").Float()
	styles := global.Call("getComputedStyle", e)
	px := styles.Get("marginTop").String()
	mt, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("marginBottom").String()
	mb, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("outlineWidth").String()
	olw, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	height = ofh + mt + mb + olw + olw
	return
}

// OutlineWidth return the block element's total outline width.
func OutlineWidth(e js.Value) (width float64) {
	styles := global.Call("getComputedStyle", e)
	px := styles.Get("outlineWidth").String()
	width, _ = strconv.ParseFloat(px[:len(px)-2], 64)
	return
}

// WidthExtras returns the block element's total distance between it's innermost and outermost widths.
func WidthExtras(e js.Value) (width float64) {
	b := BorderWidth(e)
	p := PaddingWidth(e)
	m := MarginWidth(e)
	o := OutlineWidth(e)
	width = b + p + m + o + o
	return
}

// HeightExtras returns the block element's total distance between it's innermost and the outermost heights.
func HeightExtras(e js.Value) (height float64) {
	b := BorderHeight(e)
	p := PaddingHeight(e)
	m := MarginHeight(e)
	o := OutlineWidth(e)
	height = b + p + m + o + o
	return
}

// PaddingWidth returns the block element's total padding width.
func PaddingWidth(e js.Value) (width float64) {
	styles := global.Call("getComputedStyle", e)
	px := styles.Get("paddingLeft").String()
	pl, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("paddingRight").String()
	pr, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	width = pl + pr
	return
}

// PaddingHeight returns the block element's total padding height.
func PaddingHeight(e js.Value) (height float64) {
	styles := global.Call("getComputedStyle", e)
	px := styles.Get("paddingTop").String()
	pt, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("paddingBottom").String()
	pb, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	height = pt + pb
	return
}

// MarginWidth returns the block element's total margin width.
func MarginWidth(e js.Value) (width float64) {
	styles := global.Call("getComputedStyle", e)
	px := styles.Get("marginLeft").String()
	ml, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("marginRight").String()
	mr, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	width = ml + mr
	return
}

// MarginHeight returns the block element's total margin height.
func MarginHeight(e js.Value) (height float64) {
	styles := global.Call("getComputedStyle", e)
	px := styles.Get("marginTop").String()
	mt, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("marginBottom").String()
	mb, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	height = mt + mb
	return
}

// BorderWidth returns the block element's total border width.
func BorderWidth(e js.Value) (width float64) {
	styles := global.Call("getComputedStyle", e)
	px := styles.Get("borderLeftWidth").String()
	blw, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("borderRightWidth").String()
	brw, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	width = blw + brw
	return
}

// BorderHeight returns the block element's total border height.
func BorderHeight(e js.Value) (height float64) {
	styles := global.Call("getComputedStyle", e)
	px := styles.Get("borderTopWidth").String()
	btw, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	px = styles.Get("borderBottomWidth").String()
	bbw, _ := strconv.ParseFloat(px[:len(px)-2], 64)
	height = btw + bbw
	return
}

// Style

// SetStyleHeight sets the block element's outermost height.
func SetStyleHeight(e js.Value, height float64) {
	style := e.Get(styleMemberName)
	style.Set("height", fmt.Sprintf(pxFormatter, height))
}

// SetStyleWidth sets the block element's outermost width.
func SetStyleWidth(e js.Value, width float64) {
	style := e.Get(styleMemberName)
	style.Set("width", fmt.Sprintf(pxFormatter, width))
}

// SetStyleMinWidth sets the block element's minimum outermost width.
func SetStyleMinWidth(e js.Value, width float64) {
	style := e.Get(styleMemberName)
	style.Set("min-width", fmt.Sprintf(pxFormatter, width))
}
