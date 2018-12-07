# kickwasm

A rapid development, desktop application framework generator for go. Kickwasm reads your kickwasm yaml files and generates the source code of an application framework.

## The frame work

The framework works as soon as you build it. You can build it as soon as kickwasm generates the source code.

The colors example is only a framework without anything added to it. See the colors example video below.

On the contrary, the contacts example is a framework with code added. The contacts example is an actual program, a simple CRUD. See the contacts example video below.

### The framework code is physically and logically organized into 3 main levels

1. The domain folder contains domain ( application ) level logic.
1. The mainprocess folder contains the main process level logic.
1. The renderer folder contains the renderer level logic.

### The framework has a 2 step build

The main process is a web browser running through whatever port you indicate in the application's http.yaml file. When you start the main process it opens a browser and runs the renderer process in the browser.

So when you build the framework, you build both the renderer process and the main process.

1. Build the renderer process with **GOARCH=wasm GOOS=js go build -o app.wasm main.go panels.go**
1. Build the main process with **go build**
1. The framework is run like any other simple program with **./my-app-name**

### The framework imports

* [the boltdb package.](https://github.com/boltdb/bolt)
* [the yaml package.](https://gopkg.in/yaml.v2)
* [the gorilla websocket package.](https://github.com/gorilla/websocket)

### The framework code has a TDD architecture

TDD ( test driven design ) is very simple to do in go because TDD involves the implementation of principles that already belong to go. Kickwasm attempts to create a framework with idiomatic go and a TDD friendly code architecture. You create the unit tests.

## This is version 0.5.0

### Not yet stable and possibly buggy

#### Update: Dec 7, 2018

More significant changes for various reasons as I use rekickwasm.

I also used rekickwasm to add a simple about section to the contacts example. The about section shows how tabs look and work.

#### Update: Dec 1-5, 2018

I'm still building and modifying rekickwasm and that means significant changes to

* kickwasm,
* the framework kickwasm generates,

I'll have the example videos done soon.

#### Update: Nov 23, 2018

Building the contacts CRUD example resulted in major changes and corrections to

* kickwasm,
* the framework kickwasm generates,

Following that, my focus has been on rekickwasm. Rekickwasm is the refactoring tool for kickwasm. Building rekickwasm resulted in other major changes to kickwasm and the framework it generates.

### Kickwasm imports

* [The yaml package.](https://gopkg.in/yaml.v2)

## Installation

``` bash

  go get -u github.com/josephbudd/kickwasm
  go get -u gopkg.in/yaml.v2

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
