package notjs

import (
	"strconv"
)

func (notjs *NotJS) HostPort() (host string, port uint64) {
	documentLocation := notjs.document.Get("location")
	host = documentLocation.Get("hostname").String()
	var err error
	port, err = strconv.ParseUint(documentLocation.Get("port").String(), 10, 64)
	if err != nil {
		port = uint64(0)
	}
	return
}
