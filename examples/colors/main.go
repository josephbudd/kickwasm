package main

import (
	"log"
	"path/filepath"

	"github.com/boltdb/bolt"

	"github.com/josephbudd/kickwasm/examples/colors/domain/data/filepaths"
	"github.com/josephbudd/kickwasm/examples/colors/domain/data/settings"
	"github.com/josephbudd/kickwasm/examples/colors/domain/implementations/storing/boltstoring"
	"github.com/josephbudd/kickwasm/examples/colors/domain/interfaces/storer"
	"github.com/josephbudd/kickwasm/examples/colors/mainprocess/calls"
	"github.com/josephbudd/kickwasm/examples/colors/mainprocess/callserver"
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
	 * /domain/interfaces/storer is the storer interfaces.
	 * /domain/implementations/storing/boltstoring is the bolt implementations of the storer interfaces.
	 * /domain/types is the record definitions.

*/

var (
	colorStore storer.ColorStorer
)

func main() {
	buildBoltStores()
	defer colorStore.Close()
	appSettings, err := settings.NewApplicationSettings()
	if err != nil {
		log.Println(err)
		return
	}
	callMap := calls.GetCallMap(colorStore)
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
	}
	colorStore = boltstoring.NewColorBoltDB(db, path, filepaths.GetFmode())
}

