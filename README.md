# kickwasm

A rapid development, desktop application framework generator for go. Kickwasm creates

* a browser based GUI framework source code written in html, css and go.
* a main process framework source code written in go.

## The frame work

### The framework code has a TDD architecture

TDD ( test driven design ) is very simple to do in go because TDD involves the implementation of principles that already belong to go. Kickwasm attempts to create a framework with idiomatic go and a TDD friendly code architecture. You create the unit tests if you want.

The framework works as soon as you build it. The colors example is only a framework without anything added to it. See the colors example video.

On the contrary, the contacts example is a framework with code added. The contacts example is an actual program, a simple CRUD. See the contacts example video.

### The framework has a 2 step build

1. Build the renderer process with **GOARCH=wasm GOOS=js go build -o app.wasm main.go panels.go**
1. Build the main process with **go build**
1. The framework is run like any other simple program with **./my-app-name**

### The framework code is physically and logically organized into 3 main levels

1. The domain folder contains domain ( application ) level logic.
1. The mainprocess folder contains the main process level logic.
1. The renderer folder contains the renderer level logic.

### The framework imports

* the boltdb package at https://github.com/boltdb/bolt.
* the yaml package at https://gopkg.in/yaml.v2

## This is version 0.4.0

### Not yet stable and possibly buggy

#### Update: Dec 1-5, 2018

I'm still building rekickwasm and that means more big changes to kickwasm and the framework it generates.

I'll have the example videos done soon.

#### Update: Nov 23, 2018

Building the contacts CRUD example resulted in major changes and corrections to

* kickwasm,
* the framework kickwasm generates,

Following that, my focus has been on rekickwasm. Rekickwasm is the refactoring tool for kickwasm. Building rekickwasm resulted in other major changes to kickwasm and the framework it generates.

### Kickwasm imports

* The yaml package at https://gopkg.in/yaml.v2

## Installation

``` bash

  go get -u github.com/josephbudd/kickwasm
  go get -u gopkg.in/yaml.v2

  cd $GOPATH/src/github.com/josephbudd/kickwasm/
  go install

```

### To do

* The videos for both examples.
* The rest of the wiki.
* Start from scratch installing kickwasm source code and building the examples on linux and windows so that I can correctly define the procedures.

