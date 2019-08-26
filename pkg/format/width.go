package format

import "strings"

// SameWidth pads each string in src to the same width.
func SameWidth(src []string) (sized []string) {
	var l int
	l = len(src)
	sized = make([]string, l)
	lens := make([]int, l)
	// Get max width
	var maxw int
	var i int
	var s string
	for i, s := range src {
		l = len(s)
		lens[i] = l
		if l > maxw {
			maxw = l
		}
	}
	var spaces = make([]string, maxw+1)
	for i = 0; i <= maxw; i++ {
		spaces[i] = " "
	}
	var space = strings.Join(spaces, "")
	for i, s = range src {
		l = maxw - lens[i]
		sized[i] = s + space[:l]
	}
	return
}
