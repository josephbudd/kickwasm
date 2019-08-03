# kickwasm version 5.0.0

## Still experimental because syscall/js is still experimental

[![Go Report Card](https://goreportcard.com/badge/github.com/josephbudd/kickwasm)](https://goreportcard.com/report/github.com/josephbudd/kickwasm) That A+ is for both kickwasm and the framework source code in the examples folder.

## Summary

* Kickwasm is a rapid development, desktop application framework generator for linux, windows and apple.
* It builds a framework that has 2 processes.
  1. The renderer process.
  1. The main process.
* The framework provides the GUI's layout, styles and behavior with 4 basic interfaces.
  1. button pads,
  1. tab bars,
  1. a back button,
  1. your own markup panels. A markup panel has it's own HTML template and go package, both of which are under your control. The framework only provides scrolling, sizing and some editable styles for markup panels.
* The framework provides the messaging model for the renderer and main processes.
* The framework provides the database model for the main process.

## The CRUD application

The [CRUD application](https://github.com/josephbudd/crud) is very important because

1. I tested the kickwasm version 5 by building the CRUD.
1. The crud has a wiki were I detail the steps I took when I built the crud.

### The CRUD WIKI

In the CRUD wiki, I begin with a complete plan for the full application.

Then create a small part of the application which allows me to only need to write and debug

* a small amount of my own go code for my markup panel packages in the renderer process,
* a small amount of my own go code for what is needed in the main process.

I repeately use the tools to add another missing part to the application, writing and debuging a little more of my own code to the added part. Eventually, I have added all of the parts, all of my own code, and the application is complete.

### If you really want to get a feel for kickwasm read the CRUD wiki

## This is kickwasm version 5.0.0

Version 5 is not backwards compatible. Version 5 is an implementation of these 4 priciples.

1. **Simplicity**. Remove the complex and replace it with the simple.
1. **Function**. Add the ability to do new things.
1. **Tools**. Use tools to perform complex tasks.
1. **syscall/js**. Keep up with the changes in the EXPERIMENTAL go syscall/js package. Kickwasm is currently compatible with go version 1.12.

### Simplicity

1. **Everything has been refactored and everything is simpler**.
1. Source code is organized into folders for logical reasons.
   1. **renderer/** contains the go packages for your markup panels. Also some other purely renderer code.
   1. **site/** contains the HTML templates for your markup panels. Also the css files, your image files, etc.
   1. **mainprocess/** contains the main process code. Your main process message handlers.
   1. **domain/** contains code shared by the main process and the renderer process. The LPC message definitions, the Store record definitions.
1. The **names of files** in the framework source code that you are free to edit begin with an upper case letter. The **names of files** that you must not edit begin with a lower case letter. Files with package API instructions are named **instructions.txt**.
1. The old complicated Local Procedure Call model has been removed. It has been replaced with the simpler **Local Process Communications ( LPC )** model which is managed with the tool **kicklpc**.
1. The old complicated data base model has been removed. It has been replaced with the simpler stores model which is managed with the tool **kickstore**.
1. The **kickwasm.yaml** file is simpler.

### Function

1. Added spawned tabs. A spawned tab is like any other tab in a tab bar except that you spawn and unspawn it's clones. For example, in an IRC application, you can open ( spawn ) a new chat room tab as the user joins a new chat room and let the user close ( unspawn ) the tab to leave the room.

### Tools

1. **kicklpc** is a new **Local Process Communications ( LPC )** tool. It lets you add to or remove from your application, messages sent between the main process and the renderer process.
   * Example: **kicklpc -add UpdateCustomer** would add the empty message **UpdateCustomerRenderToMainProcess** and the other empty message **UpdateCustomerMainProcessToRenderer**.
     * You would finish defining the 2 messages defined in **domain/lpc/messages/UpdateCustomer.go** so that they can contain the information you want sent.
     * You would add message handlers in your markup panel Callers that will send and or receive the messages.
     * You would finish the main process's message handler at **mainprocess/lpc/dispatch/UpdateCustomer.go** so that it processes the message and does what you need done with it.
     * You would send and receive those messages through the send and receive channels.
1. **kickstore** is a new tool that lets you add or remove **data stores** in your application. A data store is an overly simple API to a table in the application's bolt database and it's record.
   * Example: **kickstore -add Customer** would add the **Customer** store with it's API and the empty **Customer** record.
   * You would complete the the store's record definition in **domain/store/record/Customer.go** so that it contains the information needed.
   * If you want to modify the stores API, then you would edit it's interface in **domain/store/storer/Customer.go** and it's implementation in **domain/store/storing/Customer.go**.
   * In your code you would call the store's methods to update, remove, get etc.
1. **rekickwasm** is an old tool that lets you refactor your application's GUI. It has been refactored to work with kickwasm version 5.
1. **kickpack** is another old tool that allows your application to be compiled into a single executable. I made some improvements regarding error handling. You won't use kickpack but the framework's build scripts in the renderer/ folder do.

## For more information check out

* The videos below.
* The wiki linked to above.
* The CRUD application and it's wiki.

### The framework imports

* [the boltdb package.](https://github.com/boltdb/bolt)
* [the yaml package.](https://gopkg.in/yaml.v2)
* [the gorilla websocket package.](https://github.com/gorilla/websocket)

## Installation

``` shell

$ go get -u github.com/josephbudd/kickwasm
$ cd ~/go/src/github.com/josephbudd/kickwasm
$ go install
$ go get -u github.com/josephbudd/kickpack
$ cd ~/go/src/github.com/josephbudd/kickpack
$ go install
$ go get -u github.com/josephbudd/rekickwasm
$ cd ~/go/src/github.com/josephbudd/rekickwasm
$ go install
$ go get -u github.com/josephbudd/kicklpc
$ cd ~/go/src/github.com/josephbudd/kicklpc
$ go install
$ go get -u github.com/josephbudd/kickstore
$ cd ~/go/src/github.com/josephbudd/kickstore
$ go install

```

## The examples

The examples/ folder contains 2 examples.

1. The colors example is just a plain untouched framework. It was built with kickwasm version 5.0.0.
1. The spawntabs example only implements spawned tabs in a tab bar. It was built with kickwasm version 5.0.0.

## The example videos

### Colors

The colors example is 100% pure kickwasm generated framework. Nothing more.

The video demonstrates some of what the framework does on its own without any code from a developer. It also demonstrates that each service has it's own color.

[![building and running the colors example](https://i.vimeocdn.com/video/744492343_640.webp)](https://vimeo.com/305091395)

### Spawntabs

The spawntabs example is a simple tab spawning and unspawning application.

The video shows tabs being spawned and unspawned.

## The WIKI

The WIKI contains important information not included in the CRUD wiki. While the CRUD wiki shows how I wrote an application, this kickwasm wiki gets more into the reasons why things are the way they are in kickwasm.

The WIKI is a work in progress.

The WIKI has been updated to kickwasm version 5.0.0.

## To Do

### The renderer build scripts

The renderer build scripts are currently written for bash on ubuntu. Versions for windows and apple will need to be written. Any help with that would be appreciated.

### The tool kickstore

The tool kickstore has the potential to work with other types of stores.
