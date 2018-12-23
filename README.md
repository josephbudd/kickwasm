# kickwasm

A rapid development, desktop application framework generator for go. Kickwasm reads your kickwasm yaml files and generates the source code of an application framework.

## The frame work

The framework works as soon as you build it. You can build the framework as soon as kickwasm generates the source code.

The colors example is only a framework without anything added to it. See the colors example video below.

On the contrary, the contacts example is a framework with code added. The contacts example is an actual program, a simple CRUD. See the contacts example video below.

### The framework code is physically and logically organized into 3 main levels

1. The domain folder contains domain ( application ) level logic.
1. The mainprocess folder contains the main process level logic.
1. The renderer folder contains the renderer level logic.

### The framework has 2 processes

1. The **main process** is a web server running through whatever port you indicate in your application's http.yaml file.
1. When you start the main process it opens a browser which loads and runs the **renderer process**.

### The framework has a 2 step build

So when you build the framework, you build both the renderer process and the main process.

1. Build the renderer process for browsers with **GOARCH=wasm GOOS=js go build -o app.wasm main.go panels.go**
1. Build the main process for your computer with **go build**
1. The framework is run like any other simple program with **./my-app-name**

### The framework imports

* [the boltdb package.](https://github.com/boltdb/bolt)
* [the yaml package.](https://gopkg.in/yaml.v2)
* [the gorilla websocket package.](https://github.com/gorilla/websocket)

### The framework code has a TDD architecture

TDD ( test driven design ) is very simple to do in go because TDD involves the implementation of principles that already belong to go. Kickwasm attempts to create a framework with idiomatic go and a TDD friendly code architecture. You create the unit tests.

## This is version 0.6.0

### Not yet stable and possibly buggy

#### Dec 19, 2018

Edited README.

#### Dec 12, 2018

Reviewed tests and fixed minor issues found while refactoring the heck out of the contacts example with rekickwasm. Rekickwasm is still not done.

##### Panel initialization errors

Your panel initialization errors are the errors that you make in your go code in each markup panel's initialization funcs. A typical mistake that I make is attempting to get an element using the wrong id. Another typical mistake of mine is attempting to setup a call to the main process before I have created the call.

Now the framework returns your errors and wraps those errors with meaning as they are returned. The framework also logs to the javascript console, the browser alert and the application's log. So you get to see the error and where the error came from right away. This makes finding and getting rid of those initial mistakes in your panel code a lot easier.

The documentation in each markup panel's package in the /renderer/panels/ folder has been updated.

The wiki has been updated.

## Installation

``` bash

  go get -u github.com/josephbudd/kickwasm
  cd $GOPATH/src/github.com/josephbudd/kickwasm/
  go install

```

### Example Videos

#### Contacts

The contacts example is a simple **C**reate **R**eview **U**pdate **D**elete application. The video shows the contacts example working and it shows how the GUI resizes with the browser size changes.

[![building and running the contacts example](https://i.vimeocdn.com/video/744492275_640.webp)](https://vimeo.com/305091300)

#### Colors

The colors example is 100% pure kickwasm generated framework. Nothing more. The video serves to show how each service has it's own color.

[![building and running the colors example](https://i.vimeocdn.com/video/744492343_640.webp)](https://vimeo.com/305091395)

### To do

* The rest of the wiki.
* Start from scratch installing kickwasm source code and building the examples on linux and windows so that I can correctly define the procedures.
* Test with **Snap**. I'm hoping to be able to build kickwasm apps with SNAP. It looks promising. We'll see.
