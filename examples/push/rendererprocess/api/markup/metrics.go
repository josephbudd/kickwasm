// +build js, wasm

package markup

import (
	"github.com/josephbudd/kickwasm/examples/push/rendererprocess/api/window"
)

// Rectangle is the element's location and size metrics.
// Left, Right, Top, Bottom are the element's outer points.
// Width, Height are the element's outer measurements.
// InnerWidth, InnerHeight are the element's inner measurements.
type Rectangle struct {
	Left, Right, Top, Bottom float64
	Width, Height            float64
	InnerWidth, InnerHeight  float64
}

// ContainsPoint returns if the Rectangle contains the point.
func (m *Rectangle) ContainsPoint(x, y float64) (sorounds bool) {
	if x < m.Left || x > m.Right {
		return
	}
	if y < m.Top || y > m.Bottom {
		return
	}
	sorounds = true
	return
}

// Metrics returns the rectangle measurements of an element.
func (e *Element) Metrics() (rectangle *Rectangle) {
	boundingClient := e.element.Call("getBoundingClientRect")
	left := boundingClient.Get("left").Float()
	top := boundingClient.Get("top").Float()
	width := boundingClient.Get("width").Float()
	height := boundingClient.Get("height").Float()

	rectangle = &Rectangle{
		Left:        left,
		Right:       left + width,
		Top:         top,
		Bottom:      top + height,
		Width:       width,
		Height:      height,
		InnerWidth:  window.InnerWidth(e.element),
		InnerHeight: window.InnerHeight(e.element),
	}
	return
}
