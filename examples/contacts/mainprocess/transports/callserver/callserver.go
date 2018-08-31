package callserver

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"github.com/josephbudd/kickwasm/examples/contacts/mainprocess/transports/calls"
)

const pongWait = 60 * time.Second

// CallServer is a main process local procedure call.
type CallServer struct {
	host          string
	port          uint
	callsMap      map[int]*calls.LPC
	callsStruct   *calls.Calls
	DisconnectMax time.Duration

	connectionCountMutex *sync.Mutex
	// the number of web socket connections
	connectionCount int

	lastDisconnectMutex *sync.Mutex
	// when the last web socket connection was closed.
	lastDisconnect time.Time

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod time.Duration

	upgrader websocket.Upgrader
}

// NewCallServer constructs a new CallServer.
func NewCallServer(host string, port uint, callsMap map[int]*calls.LPC, callsStruct *calls.Calls) *CallServer {
	return &CallServer{
		host:          host,
		port:          port,
		callsMap:      callsMap,
		callsStruct:   callsStruct,
		DisconnectMax: time.Millisecond * 500,

		connectionCountMutex: &sync.Mutex{},
		// the number of web socket connections
		connectionCount: -1,

		lastDisconnectMutex: &sync.Mutex{},
		// when the last web socket connection was closed.
		lastDisconnect: time.Now(),

		// Send pings to peer with this period. Must be less than pongWait.
		pingPeriod: (pongWait * 9) / 10,

		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				values, found := r.Header["Origin"]
				if !found {
					log.Println("required oringin header not found")
					return false
				}
				appHost := fmt.Sprintf("%s:%d", host, port)
				for _, value := range values {
					loc, err := url.Parse(value)
					if err == nil {
						if loc.Host == appHost {
							return true
						}
					}
				}
				return false
			},
		},
	}
}
