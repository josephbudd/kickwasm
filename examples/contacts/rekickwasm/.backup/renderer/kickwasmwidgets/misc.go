package kickwasmwidgets

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// NowToHTMLValue returns the current date to an html input type = date value.
func NowToHTMLValue(t *time.Time) string {
	return time.Now().Format("01-02-2006")
}

// MoneyFormat formats a float into a money string.
func MoneyFormat(f float64) string {
	th := math.Floor(f / 1000)
	h := f - (th * 1000)
	if th > 0.0 {
		return fmt.Sprintf("$%4d,%03.2f", int64(th), h)
	}
	return fmt.Sprintf("$%11.2f", h)
}

// MoneyFormatHTML formats a float into a money string for HTML.
func MoneyFormatHTML(f float64) string {
	return strings.Replace(MoneyFormat(f), " ", "&nbsp;", -1)
}
