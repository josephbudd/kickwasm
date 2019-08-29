# kickwasm version 7.0.0

[![Go Report Card](https://goreportcard.com/badge/github.com/josephbudd/kickwasm)](https://goreportcard.com/report/github.com/josephbudd/kickwasm)

August 26, 2019:

Version 7.0.0 brings another backwards compatibility break. Now the tool **kickstore** let one add or remove APIs and records for remote databases.

## Still experimental because syscall/js is still experimental

## Summary

Kickwasm is a rapid development, desktop application framework generator for linux. In the future, windows and apple.

Kickwasm reads your simple yaml file and builds a framework. Below is an example kickwasm.yaml file where the application contains only one button and it's labeled "Hello". The button has only one markup panel which contains the static text "Hello World". When the user clicks the "Hello" button, the person sees the button's panel which reads "Hello World.". I don't need to write any code for the markup panel because it's content will not change. So I can build and run this and it will work on it's own.

``` yaml

title: Hello World App
importPath: github.com/josephbudd/helloworld
services:
 - name: SayHello
   button:
    - name: HelloButton
      label: Hello
      heading: The hello page.
      panels:
       - name: HelloWorldPanel
         note: Just static text for the user to see.
         markup: <p>Hello World</p>

```

### The framework provides

1. Two processes.
   1. The renderer process which runs inside the browser.
   1. The main process.
1. The GUI's layout, styles and behavior with 4 basic interfaces.
   1. An auto functioning button pad made from your buttons.
   1. An auto functioning tab-bar made from your tabs.
   1. A markup panel that contains
      * your markup in it's own template file
      * your functionality in it's own go package.
   1. An auto functioning back button.
1. The go package model and html template model for your markup panels.
1. The message model for passing messages between the renderer process and the main processes.
1. The database model for a default local bolt database and or remote databases.

### Source code folder structure

1. Source code is organized into folders for logical reasons.
   1. **renderer/** contains the go packages for your markup panels. Also some other purely renderer code.
   1. **site/** contains the HTML templates for your markup panels. Also the css files, your image files, etc.
   1. **mainprocess/** contains the main process code. Your main process LPC message handlers. Your business models.
   1. **domain/** contains code shared by the main process and the renderer process. The LPC message definitions, the data storeage API's and record definitions.
1. The **names of files** in the framework source code that you are free to edit begin with an upper case letter. The **names of files** that you must not edit begin with a lower case letter. Files with package API instructions are named **instructions.txt**.

## Tools

1. **kicklpc** is how you manage your application's **LPC** message model. **Local Process Communications ( LPCs )** are the messages are sent between the main process and the renderer process.
   * Example: **kicklpc -add UpdateCustomer** would add the empty message **UpdateCustomerRenderToMainProcess** and the other empty message **UpdateCustomerMainProcessToRenderer**.
     * You would complete the definitions of the 2 messages in **domain/lpc/messages/UpdateCustomer.go** so that they can contain the information you want sent.
     * You would add message handlers in your markup panel callers. Each markup panel has a caller which communicates with the main process. In a markup panel's caller you could send an **UpdateCustomerRenderToMainProcess** message to the main process through the caller's send channel and receive an **UpdateCustomerMainProcessToRenderer** message from the main process through the caller's receive channel.
     * You would add the functionality to the main process's message handler at **mainprocess/lpc/dispatch/UpdateCustomer.go** so that it processes the **UpdateCustomerRenderToMainProcess** message received from the renderer process and does what you need done with it. The handler could send an **UpdateCustomerMainProcessToRenderer** message back to the renderer through the handler's send channel.
1. **kickstore** is how you manage your application's data storage model. You can add local bolt data stores. A local bolt store is an API and a record. You can add remote databases and remote records. A remote database is an API that you must complete. A remote record a struct representing a record in a remote database. You will use it for your remote database APIs.
   * Example: **kickstore -add Customer** would add the local bolt **Customer** store API and the empty **Customer** record.
   * You would complete the the store's record definition in **domain/store/record/Customer.go** so that it contains the information needed.
   * If you want to modify behavior of the store's API, then you would
     1. Edit the API interface in **domain/store/storer/Customer.go**.
     1. Edit the API implementation in **domain/store/storing/Customer.go**.
   * In your main process code you would call the store's methods to update, remove, get etc.
1. **rekickwasm** is a tool that lets you refactor your application's GUI. It let's you edit the application's kickwasm.yaml file and then makes the required changes to your framework's GUI.
1. **kickpack** is another tool. You won't use kickpack but the framework's build scripts in the renderer/ folder do.

## The CRUD application

The [CRUD application](https://github.com/josephbudd/crud) is very important because the crud has a wiki were I detail the steps I took when I built the crud.

### The CRUD WIKI

In the CRUD wiki, I begin with a complete plan for the full application.

I write a kickwasm.yaml file to define a small part of the GUI. Then use kickwasm to create that small part of the framework. Then, one markup panel at a time

1. I add a small amount of my own go code for my markup panel package in the renderer process.
1. I add a small amount of my own go code for what is needed in the main process.
1. I debug.

I repeately use the kickwasm tools to add another missing part to the application, writing and debuging a little more of my own code to the added part. Eventually, I have added all of the parts, all of my own code, and the application is complete.

### If you really want to get a feel for kickwasm read the CRUD wiki

## For more information check out

* The videos below.
* The wiki linked to above.
* The [CRUD application](https://github.com/josephbudd/crud) and it's wiki.

## Installation

### Installing kickwasm and it's tools

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

### Installing the framework dependencies

* [the boltdb package.](https://github.com/boltdb/bolt)
* [the yaml package.](https://gopkg.in/yaml.v2)
* [the gorilla websocket package.](https://github.com/gorilla/websocket)


## The examples

The examples/ folder contains 2 examples. Their videos make very clear how the framework functions.

1. A framework always begins with a button pad where the user gets a general idea of what the application does.
1. From there, the GUI behaves according how the framework was designed.

### Colors example

The colors example is 100% pure kickwasm generated framework. Nothing more. I designed the framework with button pads and markup panels but no tab bars.

The video demonstrates some of what the framework does on its own without any code from a developer. It also demonstrates that each service has it's own color.

[![building and running the colors example](https://i.vimeocdn.com/video/744492343_640.webp?mw=550&amp;mh=310&amp;q=70)](https://vimeo.com/305091395)

### Spawntabs example

The spawntabs example is a simple tab spawning and unspawning application.

The video shows tabs being spawned and unspawned.

[![building and running the spawntabs example](https://i.vimeocdn.com/video/803691454.webp?mw=550&amp;mh=310&amp;q=70)](https://vimeo.com/351948165)

## The kickwasm WIKI

The kickwasm WIKI contains important information not included in the CRUD wiki. It is a work in progress.

The WIKI has been updated to kickwasm version 7.0.0.

## To Do

### The renderer build scripts for windows and macOS

The renderer build scripts are currently written for bash on ubuntu. Versions for windows and macOS need to be written. Any help with that would be appreciated.
