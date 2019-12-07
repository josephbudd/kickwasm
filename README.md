# kickwasm version 13.0.0

An single page application framework generator written in go for applictions written in go.

## Still experimental because syscall/js is still experimental

[![Go Report Card](https://goreportcard.com/badge/github.com/josephbudd/kickwasm)](https://goreportcard.com/report/github.com/josephbudd/kickwasm)

## December 07, 2019

### Renderer process

The renderer API has been greatly reduced to only what is needed. It has been separated from **syscall/js** so that any further changes to **syscall/js** should not effect the renderer API.

Added printing of markup panel's to the API. Added graceful closing of the application to the API.

Removed unnecessary HTML in the GUI.

### Main process

Added context to the LPC message handlers.

### Tools

Rekickwasm, the refactoring tool for kickwasm, now has automated testing. See **Installation** below.

## Kickwasm Summary

You write a kickwasm.yaml file defining the application's GUI. Kickwasm generates the framework. The framework is your application's source code, written in go, HTML, and CSS with the GUI built as defined in the kickwasm.yaml file.

The source code builds an application that runs as 2 separate processes. The main process ( business logic ) is an http server running on the client. The renderer process ( view logic) runs in the client's browser. The 2 processes communicate with messages through channels.

The main process will probably need to use local and remote data bases and communicate with remote services. That is simple and quick with the tool **kickstore**.

The main process and renderer process will need to communicate with each other by sending messages through channels. That is simple and quick with the tool **kicklpc**.

The GUI may need to be refactored. And by **refactored** I mean moving parts of your GUI and their respective go packages driving your markup panels all over the place as well as adding and removing parts of the GUI. That is safe, simple and quick with the tool **rekickwasm**.

### GUI design is simple: Buttons, Tabs, Panels

The GUI always begins with the buttons in the initial button pad. Those buttons show the user what actions the user can perform with the application. So in the kickwasm.yaml file you begin by defining the **buttons**.

#### Panels

Panels always exist in a group. Even if the group only has one panel. In a panel group, only one panel is displayed. In a panel group, showing one panel hides the other panels.

A panel can have buttons or tabs or markup.

* **A panel with buttons** is rendered in the GUI as a button pad. The GUI controls the button pad for you. When the user clicks on a button the button pad disapears and the user is shown the button's panel group.
* **A panel with tabs** is renderered in the GUI as a tab bar of tabs. The GUI controls the tab bar for you. When the user clicks on a tab the tab moves to the front of the other tabs. The front tab in the tab bar appears larger and infront of the other tabs and it's panel group is always displayed under it.
* **A panel with markup** renders it's own HTML in the GUI. You control the markup panel with it's own HTML template file and it's own go package. In the go package the **panelController** controls user input. The **panelPresenter** writes to the panel. The **panelMessenger** communicates with the main process. The **panelGroup** has a show func for each panel in the panel group. Each show func makes it's panel visible and hides the other panels in the panel group.

#### Buttons

Buttons always appear in a button pad.

Every button has 1 or more panels so in your kickwasm.yaml file you also define each button's panels. A button's panels form a panel group. In a panel group, showing one panel hides the other panels.

#### Tabs

Tabs always appear in a tab bar.

Every tab has 1 or more markup panels so in your kickwasm.yaml file you also define each tab's panels. A tab's panels form a panel group. In a panel group, showing one panel hides the other panels.

### A hello world kickwasm.yaml file

```yaml
title: Hello World
importPath: github.com/josephbudd/helloworld
buttons:
- name: HelloButton
  label: Hello
  heading: Hey There.
  panels:
  - name: HelloPanel
    note: A panel to display "hello world".
    markup: |
      <p>Hello world.</p>
```

## My step by step creation a CRUD with kickwasm

The [CRUD application](https://github.com/josephbudd/crud) was built with kickwasm. The CRUD WIKI is very important. It is where I detail the steps I took when I built the CRUD.

In the CRUD wiki, I begin with a complete plan for the full application.

I write a kickwasm.yaml file to define a small part of the GUI. Then use kickwasm to create that small part of the framework. Then, one markup panel at a time

1. I add a small amount of my own go code for my markup panel package in the renderer process.
1. I add a small amount of my own go code for what is needed in the main process.
1. I debug.

I repeately use all of the kickwasm tools to add another missing part to the application, writing and debuging a little more of my own code to the added part. Eventually, I have added all of the parts, all of my own code, and the application is complete.

## For more information check out

* The videos below.
* This kickwasm wiki linked to above.
* The [CRUD application](https://github.com/josephbudd/crud) and it's wiki.

## Installation

``` shell
$ go get -u github.com/josephbudd/kickwasm
$ cd ~/go/src/github.com/josephbudd/kickwasm
$ make install
$ make test
$ make dependencies
$ make proofs
```

**make install** will build kickwasm and put it and it's tools into the go/bin.

**make test** will run the kickwasm unit tests. They only take a few seconds.

**make dependencies** will install the 3 go packages that are required by the frameworks that you build with kickwasm. It will get them from github.com. So it might take a minute.

**make proofs** will build and run the proofs. It will take a few minutes. **The proofs make sure that rekickwasm refactors correctly.** The proofs will start and stop each of the 17 applications in kickwasm/proofs/ sub folders. Each application will be built, refactored and then run. While an application is running, internal tests are done to make sure that rekickwasm refactored the application exactly as it was supposed to. The application will popup up on your screen for less than a second while results are logged to the terminal. You can stop proofs at any time with ^c.

## Hello world example

### Hello world example STEP 1

Create a folder in your go path.
Cd into the folder and create a kickwasm.yaml file.

```shell
~/go/src/github.com/josephbudd$ mkdir helloworld
~/go/src/github.com/josephbudd$ cd helloworld
~/go/src/github.com/josephbudd/helloworld$ gedit kickwasm.yaml
```

Below is the kickwasm.yaml file contents.

```yaml
title: Hello World
importPath: github.com/josephbudd/helloworld
buttons:
- name: HelloButton
  label: Hello
  heading: Hey There.
  panels:
  - name: HelloPanel
    note: A panel to display "hello world".
    markup: |
      <p>Hello world.</p>
```

### Hello world example STEP 2

Create the framework, the application source code with kickwasm.

```shell
~/go/src/github.com/josephbudd/helloworld$ kickwasm
```

### Hello world example STEP 3

Build and run the application.
**-rp** signals to quick build the renderer process.
**-mp** signals to build the main process.
**-run** signals to run the executable.

```shell
~/go/src/github.com/josephbudd/helloworld$ kickbuild -rp -mp -run
```

## The source code examples in the examples/ folder

The examples/ folder contains 3 examples.

The 2 videos below make very clear how the framework functions.

1. The framework GUI has it's own look and behavior so you won't have to create one.
1. The application always begins with a button pad where the user gets a general idea of what the application does.
1. From there, the GUI behaves according how the framework was designed.

### Colors example

The colors example is 100% pure kickwasm generated framework. Nothing more. I designed the framework with button pads and markup panels but no tab bars. Actually I did add a style for wide panels that I want to scroll horizontally. In the style I only set the **min-width**.

The video demonstrates some of what the framework does on its own without any code from a developer. It also demonstrates that each home has it's own color. It also demonstrates markup panels with horizontal scrolling.

[![building and running the colors example](https://i.vimeocdn.com/video/744492343_640.webp?mw=550&amp;mh=310&amp;q=70)](https://vimeo.com/305091395)

### Spawntabs example

The spawntabs example is a simple tab spawning and unspawning application.

The video shows tabs being spawned and unspawned.

[![building and running the spawntabs example](https://i.vimeocdn.com/video/803691454.webp?mw=550&amp;mh=310&amp;q=70)](https://vimeo.com/351948165)
