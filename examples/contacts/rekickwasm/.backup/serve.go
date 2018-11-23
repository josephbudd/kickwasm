package main

import (
	"net/http"
	"strings"

	"github.com/josephbudd/kickwasm/examples/contacts/domain/data/filepaths"
)

/*

	TODO: Modify func serve for your special needs.

	If for example you want this main process to serve your own css files in /renderer/widgetcss/.

	  1. In func serve below add the following 2 lines:
	    case strings.HasPrefix(r.URL.Path, "/widgetcss"):
			withDefaultHeaders(w, r, serveURLPath)

	  2. In the /renderer/ folder add the /widgetcss/ folder
	     Add your css files to the /renderer/widgetcss/ folder.

	  3. In the /renderer/templates/ folder create a "head.tmpl" file if you haven't already.
	     In /renderer/templates/head.tmpl add the line:
		  <style> @import url(widgetcss/vlist.css); </style>

*/

// serve serves files from renderer folders.
func serve(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	switch {
	case r.URL.Path == "/favicon.ico":
		withDefaultHeaders(w, r, serveFavIconPath)
	case r.URL.Path == "/":
		withDefaultHeaders(w, r, serveMain)
	case strings.HasPrefix(r.URL.Path, "/css"):
		withDefaultHeaders(w, r, serveURLPath)
	case strings.HasPrefix(r.URL.Path, "/wasm"):
		withDefaultWASMHeaders(w, r, serveWASMURLPath)
	case strings.HasPrefix(r.URL.Path, "/images"):
		withDefaultHeaders(w, r, serveURLPath)
	case strings.HasPrefix(r.URL.Path, "/widgetcss"):
		withDefaultHeaders(w, r, serveURLPath)
	default:
		http.Error(w, "Not found", 404)
	}
}

func withDefaultHeaders(w http.ResponseWriter, r *http.Request, fn http.HandlerFunc) {
	header := w.Header()
	header.Set("Cache-Control", "no-cache")
	fn(w, r)
}

func withDefaultWASMHeaders(w http.ResponseWriter, r *http.Request, fn http.HandlerFunc) {
	header := w.Header()
	header.Set("Cache-Control", "no-cache")
	header.Set("Content-Type", "application/wasm")
	fn(w, r)
}

func serveFavIconPath(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepaths.GetFaviconPath())
}

func serveMain(w http.ResponseWriter, r *http.Request) {
	// func serveMainHTML is in panelMap.go
	serveMainHTML(w)
}

func serveURLPath(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepaths.BuildRendererPath(r.URL.Path))
}

func serveWASMURLPath(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepaths.BuildRendererPath(r.URL.Path[5:]))
}
