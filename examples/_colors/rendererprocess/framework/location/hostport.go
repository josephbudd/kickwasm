// +build js, wasm

package location

import (
	"strconv"
	"syscall/js"
)

var (
	document js.Value
)

func init() {
	g := js.Global()
	document = g.Get("document")
}

// HostPort returns the location's host and port.
func HostPort() (host string, port uint64) {
	documentLocation := document.Get("location")
	host = documentLocation.Get("hostname").String()
	var err error
	port, err = strconv.ParseUint(documentLocation.Get("port").String(), 10, 64)
	if err != nil {
		port = uint64(0)
	}
	return
}
