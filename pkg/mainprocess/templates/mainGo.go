package templates

// MainGo is the main.go template.
const MainGo = `{{$Dot := .}}{{$store0 := index .Stores 0}}package main

import (
	"log"
	"path/filepath"

	"github.com/boltdb/bolt"

	"{{.ApplicationGitPath}}{{.ImportDomainDataFilepaths}}"
	"{{.ApplicationGitPath}}{{.ImportDomainDataSettings}}"
	"{{.ApplicationGitPath}}{{.ImportDomainImplementationsStoringBolt}}"
	"{{.ApplicationGitPath}}{{.ImportDomainInterfacesStorers}}"
	"{{.ApplicationGitPath}}{{.ImportMainProcessCalls}}"
	"{{.ApplicationGitPath}}{{.ImportMainProcessCallServer}}"
)

/*
	YOU MAY EDIT THIS FILE.

	Rekickwasm will preserve this file for you.

	BUILD INSTRUCTIONS:

		cd renderer/
		GOARCH=wasm GOOS=js go build -o app.wasm main.go panels.go
		cd ..
		go build

*/

/*

	Data Storage:
	 * {{.ImportDomainInterfacesStorers}} is the storer interfaces.
	 * {{.ImportDomainImplementationsStoringBolt}} is the bolt implementations of the storer interfaces.
	 * {{.ImportDomainTypes}} is the record definitions.

*/

var ({{range .Stores}}
	{{call $Dot.LowerCamelCase .}}Store storer.{{.}}Storer{{end}}
)

func main() {
	buildBoltStores()
	defer {{call .LowerCamelCase $store0}}Store.Close()
	appSettings, err := settings.NewApplicationSettings()
	if err != nil {
		log.Println(err)
		return
	}
	callMap := calls.GetCallMap({{range $i, $store0 := .Stores}}{{if ne $i 0}}, {{end}}{{call $Dot.LowerCamelCase $store0}}Store{{end}})
	callServer := callserver.NewCallServer(appSettings.Host, appSettings.Port, callMap)
	callServer.Run(serve)
}

// buildBoltStores makes bolt data stores.
// Each store is the implementation of an interface defined in package repoi.
// Each store uses the same bolt database so closing one will close all.
func buildBoltStores() {
	path, err := filepaths.BuildUserSubFoldersPath("boltdb")
	if err != nil {
		log.Fatalf("filepaths.BuildFolderPath error is %q.", err.Error())
	}
	path = filepath.Join(path, "allstores.nosql")
	db, err := bolt.Open(path, filepaths.GetFmode(), nil)
	if err != nil {
		log.Fatalf("bolt.Open error is %q.", err.Error())
	}{{range .Stores}}
	{{call $Dot.LowerCamelCase .}}Store = boltstoring.New{{.}}BoltDB(db, path, filepaths.GetFmode()){{end}}
}

`

// PanelMapGo is the panelMap.go template for package main.
const PanelMapGo = `package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"{{.ApplicationGitPath}}{{.ImportDomainDataFilepaths}}"
)

/*

	DO NOT EDIT THIS FILE.

	This file is generated by kickasm and regenerated by rekickasm.

*/

const (
	mainTemplate = "main.tmpl"
	headTemplate = "{{.FileNames.HeadDotTMPL}}"
)

// serviceEmptyInsidePanelNamePathMap maps each markup panel template name to it's file path.
var serviceEmptyInsidePanelNamePathMap = {{.ServiceEmptyInsidePanelNamePathMap}}

// serveMainHTML only serves up main.tmpl with all of the templates for your markup panels.
func serveMainHTML(w http.ResponseWriter) {
	templateFolderPath := filepaths.GetTemplatePath()
	t := template.New(mainTemplate)
	t, err := t.ParseFiles(filepath.Join(templateFolderPath, mainTemplate))
	if err != nil {
		http.Error(w, err.Error(), 300)
		return
	}
	for _, namePathMap := range serviceEmptyInsidePanelNamePathMap {
		for name, folders := range namePathMap {
			folderPath := strings.Join(folders, string(os.PathSeparator))
			tpath := filepath.Join(templateFolderPath, folderPath, name+".tmpl")
			t, err = t.ParseFiles(tpath)
			if err != nil {
				http.Error(w, err.Error(), 300)
				return
			}
		}
	}
	// the head template which contains
	//  * any css imports
	//  * any javascript imports
	// needed for this applicaion
	tpath := filepath.Join(templateFolderPath, headTemplate)
	// it's ok if the template is not there
	// but if it's there use it.
	if _, err := os.Stat(tpath); os.IsNotExist(err) {
		// the template file does not exist so inform the developer.
		temp := fmt.Sprintf("%[1]s%[1]s define %[3]q %[2]s%[2]s<!-- You do not have a %[3]s file to import your css files. Feel free to add one in the render/template folder. -->%[1]s%[1]s end %[2]s%[2]s", "{", "}", headTemplate)
		t, err = t.Parse(temp)
		if err != nil {
			http.Error(w, err.Error(), 300)
			return
		}
	} else {
		// the file exists so parse it
		t, err = t.ParseFiles(tpath)
		if err != nil {
			http.Error(w, err.Error(), 300)
			return
		}
	}
	// do the template
	if err := t.ExecuteTemplate(w, mainTemplate, nil); err != nil {
		http.Error(w, err.Error(), 300)
	}
}

`

// ServeGo is the serve.go template which is the web server.
const ServeGo = `package main

import (
	"net/http"
	"strings"

	"{{.ApplicationGitPath}}{{.ImportDomainDataFilepaths}}"
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

`
