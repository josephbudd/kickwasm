package lpc

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

func (server *Server) incConnectionCount() {
	server.connectionCountMutex.Lock()
	if connectionCount == -1 {
		connectionCount = 1
	} else {
		connectionCount++
	}
	server.connectionCountMutex.Unlock()
}

func (server *Server) decConnectionCount() {
	server.connectionCountMutex.Lock()
	if connectionCount > 0 {
		connectionCount--
	}
	server.connectionCountMutex.Unlock()
}

// GetConnectionCount returns the connection count.
func (server *Server) GetConnectionCount() int {
	server.connectionCountMutex.Lock()
	cc := connectionCount
	server.connectionCountMutex.Unlock()
	return cc
}

func (server *Server) setLastDisconnect(t time.Time) {
	lastDisconnectMutex.Lock()
	lastDisconnect = t
	lastDisconnectMutex.Unlock()
}

// GetLastDisconnect returns the time of the last disconnect.
func (server *Server) GetLastDisconnect() time.Time {
	lastDisconnectMutex.Lock()
	t := lastDisconnect
	lastDisconnectMutex.Unlock()
	return t
}
