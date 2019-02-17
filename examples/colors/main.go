package main

import (
	"fmt"
	"log"
	"net"
	"path/filepath"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"

	"github.com/josephbudd/kickwasm/examples/colors/domain/data/filepaths"
	"github.com/josephbudd/kickwasm/examples/colors/domain/data/settings"
	"github.com/josephbudd/kickwasm/examples/colors/domain/implementations/storing/boltstoring"
	"github.com/josephbudd/kickwasm/examples/colors/domain/interfaces/storer"
	"github.com/josephbudd/kickwasm/examples/colors/domain/types"
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
	var err error
	// build the stores and setup the close.
	if err = buildBoltStores(); err != nil {
		log.Println(err)
		return
	}
	// close the bolt store later.
	defer colorStore.Close()
	// get the application's host and port and then setup the listener.
	var appSettings *types.ApplicationSettings
	if appSettings, err = settings.NewApplicationSettings(); err != nil {
		log.Println(err)
		return
	}
	// initialize and start the listener.
	// the listener may have reset the address if "localhost:0".
	// use the listener's address.
	location := fmt.Sprintf("%s:%d", appSettings.Host, appSettings.Port)
	var listener net.Listener
	if listener, err = net.Listen("tcp", location); err != nil {
		log.Println(err)
		return
	}
	// build the callMap
	callMap := calls.GetCallMap(colorStore)
	// make the call server and start it.
	callServer := callserver.NewCallServer(listener, callMap)
	callServer.Run(serve)
}

// buildBoltStores makes bolt data stores.

// The store is an implementation of an interface defined in package storer.
// Close the bolt store later.
func buildBoltStores() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "buildBoltStores()")
		}
	}()

	var path string
	if path, err = filepaths.BuildUserSubFoldersPath("boltdb"); err != nil {
		err = errors.WithMessage(err, "filepaths.BuildUserSubFoldersPath(\"boltdb\")")
		return
	}
	path = filepath.Join(path, "allstores.nosql")
	var db *bolt.DB
	if db, err = bolt.Open(path, filepaths.GetFmode(), nil); err != nil {
		err = errors.WithMessage(err, "bolt.Open(path, filepaths.GetFmode(), nil)")
		return
	}
	colorStore = boltstoring.NewColorBoltDB(db, path, filepaths.GetFmode())
	return
}

