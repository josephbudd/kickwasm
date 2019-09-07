package notjs

import (
	"net/url"
	"strconv"
	"strings"
)

// HostPort returns the document location host and port.
func (notjs *NotJS) HostPort() (host string, port uint64) {
	documentLocation := notjs.document.Get("location").String()
	// "http://127.0.0.1:9094/"
	URL, err := url.Parse(documentLocation)
	if err != nil {
		return
	}
	parts := strings.Split(URL.Host, ":")
	host = parts[0]
	if len(parts) == 2 {
		var err error
		port, err = strconv.ParseUint(parts[1], 10, 64)
		if err != nil {
			port = uint64(0)
		}
	}
	return
}
