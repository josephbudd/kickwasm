package lpc

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"time"
)

var (
	myServer *http.Server
)

// Run runs the application until the main process terminates.
// The main process detects when the browser window closes and then stops.
// The browser window also closes if the main process stops.
//
// Package callserver handles all local http requests where either of the following 2 conditions are true.
//  strings.HasPrefix(r.URL.Path, "/ws")
//  r.URL.Path == "/callserver.js"
// All other requests are passed to the message handlerFunc.
// Param: handlerFunc http.HandlerFunc
//        You will want your handlerFunc to do things like load your javascript, css and any other files.
func (server *Server) Run(handlerFunc http.HandlerFunc) error {
	appurl := "http://" + server.listener.Addr().String()
	log.Println("listen and serve: ", appurl)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		server.serve(w, r, handlerFunc)
	})
	// start the server
	go func() {
		if waitServer(appurl) && startBrowser(appurl) {
			log.Printf("A browser window should open. If not, please visit %s", appurl)
		} else {
			log.Printf("Please open your web browser and visit %s", appurl)
		}
	}()
	// start the still connected loop.
	// 2 func ending possibilities:
	//  1.a user closed browser window causing a connection error.
	//  1.b stillConnectedLoop detects the error, stops the server and itself.
	//  1.c myServer.Serve(server.listener) ends because the server is closed.
	//
	//  2.a the terminal user types ^c or ^\ and stopRunLoopCh gets the signal.
	//  2.b stillConnectedLoop stops the server and itself.
	//  2.c myServer.Serve(server.listener) ends because the server is closed.
	stopRunLoopCh := make(chan os.Signal, 1)
	signal.Notify(stopRunLoopCh, os.Interrupt)
	go server.stillConnectedLoop(stopRunLoopCh)
	// start the server
	myServer = &http.Server{}
	return myServer.Serve(server.listener)
}

// startBrowser tries to open the URL in a browser, and returns
// whether it succeed.
// This code is from google's godoc application.
func startBrowser(url string) bool {
	// try to start the browser
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	args2 := append(args[1:], url)
	log.Println(args[0], strings.Join(args2, ", "))
	cmd := exec.Command(args[0], args2...)
	return cmd.Start() == nil
}

// waitServer waits some time for the http Server to start
// serving url. The return value reports whether it starts.
// This code is from google's godoc application.
func waitServer(url string) bool {
	tries := 20
	for tries > 0 {
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
			return true
		}
		time.Sleep(100 * time.Millisecond)
		tries--
	}
	return false
}

// serve
//  * serves the web socket requests to func serveWebSocket.
//  * passes others to the handlerFunc func.
func (server *Server) serve(w http.ResponseWriter, r *http.Request, handlerFunc http.HandlerFunc) {
	// Only method GET allowed:
	// * GET + path "/ws" initializes the web sockets.
	// * Other GETs are not allowed.
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/ws") {
		// This is the initial web socket call from the renderer.
		// It will start the main process web socket server, opening the web socket connection.
		// The renderer and main process will send messages back and forth through this single connection.
		// See domain/renderer/client.go func NewClient.
		server.serveWebSocket(w, r)
		return
	}
	// Other requests.
	// See mainprocess/server.go
	handlerFunc(w, r)
}

// stillConnectedLoop periodically checks the number of web socket connections.
//  which are set in the func serveWebSocket.
// If there are no connections and the server has been disconnected for too long
//  * it closes myServer and returns
//  * the closed server causes myServer.ListenAndServe() in func Run to end.
//  * which causes func Run to end.
//  * An empty struct is sent on the quit chan signaling func main to stop the listening lpc go funcs.
// If there is an operationg system interrupt signal
//  * it closes myServer and returns
//  * the closed server causes myServer.ListenAndServe() in func Run to end.
//  * which causes func Run to end.
//  * An empty struct is sent on the quit chan signaling func main to stop the listening lpc go funcs.
func (server *Server) stillConnectedLoop(stopRunLoopCh chan os.Signal) {
	ticker := time.NewTicker(server.DisconnectMax / 2)
	defer ticker.Stop()
	for {
		select {
		case now := <-ticker.C:
			if server.GetConnectionCount() == 0 {
				if now.Sub(server.GetLastDisconnect()) > server.DisconnectMax {
					myServer.Close()
					log.Println("Ending server. Renderer disconnected.")
					server.QuitChan <- struct{}{}
					return
				}
			}
		case <-stopRunLoopCh:
			myServer.Close()
			log.Println("Ending server. Main process canceled.")
			server.QuitChan <- struct{}{}
			return
		}
	}
}
