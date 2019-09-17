package main

import (
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickwasm/examples/push/domain/data/filepaths"
	"github.com/josephbudd/kickwasm/examples/pushsitepack"
)

const (
	wasmPrefix     = "/wasm"
	wasmExceDotJS  = "/wasm/wasm_exec.js"
	wasmAppDotWASM = "/wasm/app.wasm"
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
	case strings.HasPrefix(r.URL.Path, "/mycss"):
		withDefaultHeaders(w, r, serveFileStore)
	case r.URL.Path == wasmExceDotJS:
		withDefaultHeaders(w, r, serveFileStore)
	case r.URL.Path == wasmAppDotWASM:
		withDefaultHeaders(w, r, serveFileStore)
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
		// the wasm prefix is only a flag to use wasm headers.
		// there is no wasm folder.
		urlPath = urlPath[len(wasmPrefix):]
	}
	path = filepath.Join(filepaths.GetShortSitePath(), urlPath)
	if bb, found = pushsitepack.Contents(path); !found {
		log.Printf("404 Error: %q not found", path)
		http.Error(w, "Not found", 404)
		return
	}
	header := w.Header()
	contentType := mime.TypeByExtension(filepath.Ext(urlPath))
	header.Set("Content-Type", contentType)
	var err error
	if _, err = w.Write(bb); err != nil {
		http.Error(w, err.Error(), 300)
	}
}
