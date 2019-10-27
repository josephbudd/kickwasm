# kickwasm version 11.0.0

An application framework generator written in go for applictions written in go.

## Still experimental because syscall/js is still experimental

[![Go Report Card](https://goreportcard.com/badge/github.com/josephbudd/kickwasm)](https://goreportcard.com/report/github.com/josephbudd/kickwasm)

## November 1, 2019

* A barely significant, single, tiny, framework source code file, formatting fix to get a better score in go report card. I noticed it while rebuilding my linux morse code trainer [CWT](https://github.com/josephbudd/cwt).

## October 23, 2019

* Rebuilt all of the tests. In the past few weeks I've reengineered kickwasm and it's tools 3 times. Apparently I neglected to rebuild the kickwasm tests after the last round of changes. Now the tests are rebuilt.

## October 22, 2019

Version 11.0.0

* Got rid of unused and unwanted legacy left behinds.
* Added VSCode work space files in the framework source code.
* All of the tools are in the tools/ folder and are rebuilt.
* Rebuilt all examples.
* Rebuilt the CRUD and it's wiki.

## Kickwasm Summary

You write a kickwasm.yaml file defining the application's GUI. Kickwams generates the framework which is the source code for a go desktop applicaiton with the GUI built as defined in the kickwasm.yaml file.

The source code builds an application that runs as 2 separate processes. The main process ( business logic ) is an http server running on the client. The renderer process ( renderer logic) runs in the client's browser. The 2 processes communicate through channels.

The main process will probably need to store and retrieve data or communicate with a remote service. That is managed with the tool kickstore.

The main process and renderer process will need to communicate with messages through channels. That is managed with the tool kicklpc.

The renderer process may need to be refactored. That is handled with the tool rekickwasm.

## My step by step creation a CRUD with kickwasm

The [CRUD application](https://github.com/josephbudd/crud) was built with kickwasm. The CRUD WIKI is very important it is where I detail the steps I took when I built the CRUD.

In the CRUD wiki, I begin with a complete plan for the full application.

I write a kickwasm.yaml file to define a small part of the GUI. Then use kickwasm to create that small part of the framework. Then, one markup panel at a time

1. I add a small amount of my own go code for my markup panel package in the renderer process.
1. I add a small amount of my own go code for what is needed in the main process.
1. I debug.

I repeately use the kickwasm tools to add another missing part to the application, writing and debuging a little more of my own code to the added part. Eventually, I have added all of the parts, all of my own code, and the application is complete.

## For more information check out

* The videos below.
* The wiki linked to above.
* The [CRUD application](https://github.com/josephbudd/crud) and it's wiki.

## Installation

### Installing kickwasm and it's tools

``` shell

$ go get -u github.com/josephbudd/kickwasm
$ cd ~/go/src/github.com/josephbudd/kickwasm
$ make install
$ make test

```

### Installing the framework dependencies

* [the boltdb package.](https://github.com/boltdb/bolt)
* [the yaml package.](https://gopkg.in/yaml.v2)
* [the gorilla websocket package.](https://github.com/gorilla/websocket)

## The examples

The examples/ folder contains 3 examples. The new example is **spawnwidgets**. It demonstrates how to create and use widgets with spawned tab panels.

The 2 videos below make very clear how the framework functions.

1. The framework GUI has it's own look and behavior so you won't have to create one.
1. The application always begins with a button pad where the user gets a general idea of what the application does.
1. From there, the GUI behaves according how the framework was designed.

### Colors example

The colors example is 100% pure kickwasm generated framework. Nothing more. I designed the framework with button pads and markup panels but no tab bars.

The video demonstrates some of what the framework does on its own without any code from a developer. It also demonstrates that each home has it's own color.

[![building and running the colors example](https://i.vimeocdn.com/video/744492343_640.webp?mw=550&amp;mh=310&amp;q=70)](https://vimeo.com/305091395)

### Spawntabs example

The spawntabs example is a simple tab spawning and unspawning application.

The video shows tabs being spawned and unspawned.

[![building and running the spawntabs example](https://i.vimeocdn.com/video/803691454.webp?mw=550&amp;mh=310&amp;q=70)](https://vimeo.com/351948165)
