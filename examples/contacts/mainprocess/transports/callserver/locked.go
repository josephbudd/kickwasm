package callserver

import (
	"sync"
	"time"
)

// these are vars that need to be locked to avoid data races.
var (
	connectionCount = -1

	lastDisconnectMutex = &sync.Mutex{}
	// when the last web socket connection was closed.
	lastDisconnect = time.Now()
)

func (callServer *CallServer) incConnectionCount() {
	callServer.connectionCountMutex.Lock()
	if connectionCount == -1 {
		connectionCount = 1
	} else {
		connectionCount++
	}
	callServer.connectionCountMutex.Unlock()
}

func (callServer *CallServer) decConnectionCount() {
	callServer.connectionCountMutex.Lock()
	if connectionCount > 0 {
		connectionCount--
	}
	callServer.connectionCountMutex.Unlock()
}

// GetConnectionCount returns the connection count.
func (callServer *CallServer) GetConnectionCount() int {
	callServer.connectionCountMutex.Lock()
	cc := connectionCount
	callServer.connectionCountMutex.Unlock()
	return cc
}

func (callServer *CallServer) setLastDisconnect(t time.Time) {
	lastDisconnectMutex.Lock()
	lastDisconnect = t
	lastDisconnectMutex.Unlock()
}

// GetLastDisconnect returns the time of the last disconnect.
func (callServer *CallServer) GetLastDisconnect() time.Time {
	lastDisconnectMutex.Lock()
	t := lastDisconnect
	lastDisconnectMutex.Unlock()
	return t
}
