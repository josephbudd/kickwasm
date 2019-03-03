# kickwasm version 1.1.0

## Description

A rapid development, desktop application framework generator for linux, windows and apple.

KickWasm lets you construct a working framework for a go program and then add your own code to the framework in order to turn the framework into a real application. You simply follow these steps.

1. Write a kickwams.yaml file(s)
    1. You begin by defining each button in the framework's opening button pad.
    1. You continue by defining each button's panel(s).
    1. Now a panel can only have buttons or tabs or markup so you continue by defining each new panel's buttons or tabs or markup.
    1. The cycle continues as you define each new button's panel(s) and each new tab's panel(s).
1. Generate the framework source code with the command **kickwasm -f path-to-your-kickwasm.yaml**. If you want you can immediately build and run the framework just to see it work.
1. Turn the framework into an application by adding your own code one panel at a time.

## The framework

The framework works as soon as you build it. You can build the framework as soon as kickwasm generates the framework source code.

The colors example in the examples/ folder, is only a framework without anything added to it. See the colors example video below.

### The framework code is physically and logically organized into 4 areas of logic

1. The go code in the **domain/** folder is the domain ( shared ) logic.
1. The go code in the **mainprocess/** folder is the main process logic.
1. The go code in the **renderer/** folder is the renderer logic. The go code in the **renderer/** folder is compiled into wasm into the **site/app.wasm** file.
1. The **site/** folder contains the compiled **app.wasm** file from the **renderer/** folder. It also contains the HTML templates as well as the css and any other files for the browser.

### The framework has 2 processes

1. The **main process** is a web server running through whatever port you indicate in your application's http.yaml file. When you start the main process it opens a browser which loads and runs the renderer process from the **site/** folder.
1. The **renderer process** is all of the wasm, html, css, images, etc contained in the **site/** folder.

### The framework has a 2 step build

So when you build the framework, you build both the main process and the renderer process.

You build the main process executable from the framework's root folder which has it's main.go.

You build the renderer process from the **renderer/** folder which has it's main.go. But you use the shell script **renderer/build.sh** to build the renderer's wasm byte code into the **site/** folder.

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

1. The colors example which is just a plain untouched framework. It wasm built with kickwasm version 1.2.0.
1. The contacts example which is a simple **C**reate **R**ead **U**pdate **D**elete application. It was built with kickwasm version 1.0.0. so the framework code is a little older.

## The example videos

### Colors

The colors example is 100% pure kickwasm generated framework. Nothing more.

The video demonstrates some of what the framework does on its own without any code from a developer. It also demonstrates that each service has it's own color.

[![building and running the colors example](https://i.vimeocdn.com/video/744492343_640.webp)](https://vimeo.com/305091395)

### Contacts

The contacts example is a simple CRUD application.

The video demonstrates some practical capabilities of the framework.

[![building and running the contacts example](https://i.vimeocdn.com/video/744492275_640.webp)](https://vimeo.com/305091300)

## The WIKI

Wow! The WIKI is where I attempt to demonstrate how

1. Kickwasm generates a framework by reading your **kickwasm.yaml** file.
1. A framework is turned into a real application.

I try to do so in very small meaningful steps. I mostly use the contacts example for reference.

The WIKI is a work in progress. I am still devoted to the WIKI.

## Tools

**Rekickwasm** is a refactoring tool for a framework generated with kickwasm. Rekickwasm only refactors the renderer part of the framework. I have been using it to refactor the contacts example renderer in all kinds of ways.

## Changelog for this latest version 1.2.0

### Feature added to the renderer part of the framework

Version 1.2.0 adds locking the back button and tab buttons so that they don't function and unlocking the back and tab buttons so that they do function.
