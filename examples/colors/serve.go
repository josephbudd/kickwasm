package main

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickwasm/examples/colorssitepack"
	"github.com/josephbudd/kickwasm/examples/colors/domain/data/filepaths"
)

const (
	wasmPrefix = "/wasm"
)

/*

	TODO: Modify func serve for your special needs.

	If for example you want this main process to serve your own css files in /site/widgetcss/.

	  1. In func serve below add the following 2 lines:
	    case strings.HasPrefix(r.URL.Path, "/widgetcss"):
			withDefaultHeaders(w, r, serveURLPath)

	  2. In the /site/ folder add the /widgetcss/ folder
	     Add your css files to the /site/widgetcss/ folder.

	  3. In the /site/templates/ folder create a "head.tmpl" file if you haven't already.
	     In /site/templates/head.tmpl add the line:
		  <style> @import url(widgetcss/vlist.css); </style>

	  4. Rebuild the renderer process.
		 $ cd renderer/
		 $ build.sh

	  5. Rebuild the main process.
		 $ cd ..
		 $ go build

*/

// serve serves files from renderer folders.
func serve(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	switch {
	case r.URL.Path == "/":
		withDefaultHeaders(w, r, serveMain)
	case strings.HasPrefix(r.URL.Path, "/css"):
		withDefaultHeaders(w, r, serveFileStore)
	case strings.HasPrefix(r.URL.Path, wasmPrefix):
		withDefaultWASMHeaders(w, r, serveFileStore)
	case r.URL.Path == "/favicon.ico":
		withDefaultHeaders(w, r, serveFileStore)
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

func serveFileStore(w http.ResponseWriter, r *http.Request) {
	var bb []byte
	var found bool
	var path string
	var urlPath string
	urlPath = r.URL.Path
	// fix url path
	if strings.HasPrefix(urlPath, wasmPrefix) {
		// the wasm prefix only flags to use wasm headers.
		// there is no wams folder.
		urlPath = urlPath[len(wasmPrefix):]
	}
	path = filepath.Join(filepaths.GetShortSitePath(), urlPath)
	if bb, found = colorssitepack.Contents(path); !found {
		log.Println("%q not found", path)
		http.Error(w, "Not found", 404)
		return
	}
	var err error
	if _, err = w.Write(bb); err != nil {
		http.Error(w, err.Error(), 300)
	}
}
