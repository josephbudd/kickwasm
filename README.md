# kickwasm version 6.0.0

[![Go Report Card](https://goreportcard.com/badge/github.com/josephbudd/kickwasm)](https://goreportcard.com/report/github.com/josephbudd/kickwasm)

August 8, 2019:

Another backwards compatibility break. This time in order to better comply with the go report card. This change resulted in simpler framework source code and some spelling corrections.

I also made the same backwards compatibility breaking changes to the tools. This wiki is updated.

I rebuild the crud with version 6.0.0.

Sunday August 4, 2019:

I updated the crud wiki with corrections. Uploaded all the new videos. Uploaded files I didn't get uploaded so all the projects are fully uploaded.

Saturday August 3, 2019:

I just got the new kickwasm, its tools and the crud app uploaded to github. Now I'm working on getting the new videos uploaded, some minor details in the wikis and redoing the go report card stuff.

## Still experimental because syscall/js is still experimental

## Summary

* Kickwasm is a rapid development, desktop application framework generator for linux. In the future, windows and apple.
* Kickwasm reads your simple yaml file and builds a framework.

### The framework provides

1. Two processes.
   1. The renderer process.
   1. The main process.
1. The GUI's layout, styles and behavior with 4 basic interfaces.
   1. A button pad that contains your buttons.
   1. A tab-bar that contains your tabs.
   1. A markup panel that contains your markup.
   1. A back button.
1. The go package model and html template model for your markup panels.
1. The message model for passing messages between the renderer process and the main processes.
1. The database model.

## The CRUD application

The [CRUD application](https://github.com/josephbudd/crud) is very important because

1. I tested the kickwasm version 5 by building the CRUD.
1. The crud has a wiki were I detail the steps I took when I built the crud. A lot of what is mysteriously implied here in this short README is explicitly demonstrated in the crud wiki.

### The CRUD WIKI

In the CRUD wiki, I begin with a complete plan for the full application.

I write a kickwasm.yaml file to define a small part of the GUI. Then use kickwasm to create that small part of the framework. Then, one markup panel at a time

1. I add a small amount of my own go code for my markup panel package in the renderer process.
1. I add a small amount of my own go code for what is needed in the main process.
1. I debug.

I repeately use the kickwasm tools to add another missing part to the application, writing and debuging a little more of my own code to the added part. Eventually, I have added all of the parts, all of my own code, and the application is complete.

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
     * You would complete the definitions of the 2 messages in **domain/lpc/messages/UpdateCustomer.go** so that they can contain the information you want sent.
     * You would add message handlers in your markup panel callers. Each markup panel has a caller which communicates with the main process. In a markup panel's caller you could send an **UpdateCustomerRenderToMainProcess** message to the main process through the caller's send channel and receive an **UpdateCustomerMainProcessToRenderer** message from the main process through the caller's receive channel.
     * You would add the functionality to the main process's message handler at **mainprocess/lpc/dispatch/UpdateCustomer.go** so that it processes the **UpdateCustomerRenderToMainProcess** message received from the renderer process and does what you need done with it. The handler could send an **UpdateCustomerMainProcessToRenderer** message back to the renderer through the handler's send channel.
1. **kickstore** is a new tool that lets you add or remove **data stores** in your application. A data store is an overly simple API to a table in the application's bolt database and it's record.
   * Example: **kickstore -add Customer** would add the **Customer** store with it's API and the empty **Customer** record.
   * You would complete the the store's record definition in **domain/store/record/Customer.go** so that it contains the information needed.
   * If you want to modify behavior of the store's API, then you would
     1. Edit the API interface in **domain/store/storer/Customer.go**.
     1. Edit the API implementation in **domain/store/storing/Customer.go**.
   * In your main process code you would call the store's methods to update, remove, get etc.
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

1. The colors example is just a plain untouched framework. It was built with kickwasm version 6.0.0.
1. The spawntabs example only implements spawned tabs in a tab bar. It was built with kickwasm version 6.0.0.

## The example videos

The videos make very clear how the framework functions.

1. A framework always begins with a button pad where the user gets a general idea of what the application does.
1. From there, the GUI behaves according how the framework was designed.

### Colors

The colors example is 100% pure kickwasm generated framework. Nothing more. I designed the framework with button pads and markup panels but no tab bars.

The video demonstrates some of what the framework does on its own without any code from a developer. It also demonstrates that each service has it's own color.

[![building and running the colors example](https://i.vimeocdn.com/video/744492343_640.webp?mw=550&amp;mh=310&amp;q=70)](https://vimeo.com/305091395)

### Spawntabs

The spawntabs example is a simple tab spawning and unspawning application.

The video shows tabs being spawned and unspawned.

[![building and running the spawntabs example](https://i.vimeocdn.com/video/803691454.webp?mw=550&amp;mh=310&amp;q=70)](https://vimeo.com/351948165)

## The kickwasm WIKI

The kickwasm WIKI contains important information not included in the CRUD wiki. It is a work in progress.

The WIKI has been updated to kickwasm version 6.0.0.

## To Do

### The renderer build scripts

The renderer build scripts are currently written for bash on ubuntu. Versions for windows and apple will need to be written. Any help with that would be appreciated.

### The tool kickstore

The tool kickstore, like kickwasm is a work in progress. Kickstore needs to work with more types of databases. That will require some more refactoring with kickwasm as the bolt database is built into the framework.
