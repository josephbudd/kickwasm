# kickwasm

A rapid development, desktop application framework generator for go.

Kickwasm reads your kickwasm.yaml file and generates the source code of an application framework.

## The frame work

The framework works as soon as you build it. You can build the framework as soon as kickwasm generates the source code.

The colors example in the examples/ folder, is only a framework without anything added to it. See the colors example video below.

### The framework code is physically and logically organized into 4 main levels

1. The **domain/** folder contains domain ( application ) level logic.
1. The **mainprocess/** folder contains the main process level logic.
1. The **renderer/** folder contains the renderer level logic.
1. The **site/** folder contains the wasm, templates, styles etc for the browser.

### The framework has 2 processes

1. The **main process** is a web server running through whatever port you indicate in your application's **http.yaml** file.
1. When you start the main process it opens a browser which loads and runs the **renderer process** from the **site/** folder.

### The framework has a 2 step build

So when you build the framework, you build both the renderer process and the main process. The renderer process code is in the **renderer/** folder but it is built into the **site/** folder.

There is a shell script in the **renderer/** folder that builds the renderer process into the **site/** folder. It's **build.sh**.

Here is a build example using the colors example.

``` bash

cd $GOPATH
cd src/github.com/josephbudd/kickwasm/examples/colors
cd renderer
./build.sh
cd ..
go build
./colors

```

### The framework imports

* [the boltdb package.](https://github.com/boltdb/bolt)
* [the yaml package.](https://gopkg.in/yaml.v2)
* [the gorilla websocket package.](https://github.com/gorilla/websocket)

## This is version 1.0.2. January 31, 2019: Stable

Rewrote notJS.RemoveChildNodes in the generated source code so that it removes all text and html from an element. OK I'm done messing with notJS. If I want to make any more changes I'll just write an external package or I'll find one that exists.

## Previous version 1.0.1

Corrections to the documentation in the panelCaller.go files.

Kickwasm is stable. No further breaking changes forseen. That is because rekickwasm, the refactoring tool is now stable.

**Rekickwasm** is a refactoring tool for a framework generated with kickwasm. Rekickwasm only refactors the renderer part of the framework. I have been using it to refactor the contacts example renderer in all kinds of ways.

## Installation

``` bash

  go get -u github.com/josephbudd/kickwasm
  cd $GOPATH/src/github.com/josephbudd/kickwasm/
  go install

```

## Distribution

Once you build your application you can distribute it. You can distribute it as a folder. You only need to collect the following 3 items together in that folder.

1. the **executable** file.
1. the **http.yaml** file.
1. the **site/** folder.

### For the examples/colors/ application that would be

1. the **examples/colors/colors** file which is the executable.
1. the **examples/colors/http.yaml** file.
1. the **examples/colors/site/** folder.

### For the examples/contacts/ application that would be

1. the **examples/contacts/contacts** file which is the executable.
1. the **examples/contacts/http.yaml** file.
1. the **examples/contacts/site/** folder.

## The examples

The examples/ folder contains 2 examples.

1. The colors example which is just a plain untouched framework.
1. The contacts example which is a simple **C**reate **R**eview **U**pdate **D**elete application.

## The example videos

### Colors

The colors example is 100% pure kickwasm generated framework. Nothing more.

The video demonstrates some of what the framework does on its own without any code from a developer. It also demonstrates that each service has it's own color.

[![building and running the colors example](https://i.vimeocdn.com/video/744492343_640.webp)](https://vimeo.com/305091395)

### Contacts

The contacts example is a simple CRUD application.

The video demonstrates some practical capabilities of the framework.

[![building and running the contacts example](https://i.vimeocdn.com/video/744492275_640.webp)](https://vimeo.com/305091300)

### The WIKI

Wow! The wiki is where I attempt to demonstrate how a framework is turned into a real application.

I try to do so in very small meaningful steps. I mostly use the contacts example for reference.

### Yet to do

* The rest of the wiki.
* Start from scratch installing kickwasm source code and building the examples on linux and windows so that I can correctly define the procedures.
