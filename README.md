# kickwasm

kickwasm is a single page application framework generator written in GO for applictions written in GO. kickwasm has it's own tools to keep development with kickwasm simple.

I developed the kickwasm GUI for a client who never used a computer. The current kickwasm GUI is the final design developed while testing with that user.

**Still experimental because the GO package syscall/js is still experimental.**

## Goals

1. Find and fix bugs! 😖
1. Keep the API simple or even simpler so that I have the freedom to do what I want.
1. Not break the API any further.
1. Keep up with changes in syscall/js.

[![Go Report Card](https://goreportcard.com/badge/github.com/josephbudd/kickwasm)](https://goreportcard.com/report/github.com/josephbudd/kickwasm)

## June 16, 2020

### Version 16.0.0

All of the built in tests run with the make file work perfectly however, I always build new applications for the final tests.

1. I am currently building new applications.
1. I still have to update the wiki.

Thanks to bill gates' attempt to take over the world with his chinese virus, this version does contain API breaking changes. All of my usual wifi hot spots had been and still are closed. However, the park just opened up.

Version 16 also contains some improvements I needed, bug fixes and documentation fixes. Also I made the following additions.

1. Refactored for go version 1.14's major changes to syscall/js.
1. The tiny renderer process API package at rendererprocess/api/jsvalue allows the developer to use the the go package's "syscall/js" Value and still work seamlessly with the framework's renderer process. It's for developers who only want to use syscall/js.
1. The larger renderer process API package at renderprocess/api/markup wraps syscall/js while allowing the user to work seamlessly with the framework's renderer process. It has been expaneded.
1. Tabs. The user can slide tabs to different positions in the tab bar. Also tabs are no longer printed.
1. Improvements to how the rendererprocess/framework/viewtools package resizes markup panels.
1. Improvments to rekickwasm and added a test and proof to demonstrate that the change works.
1. I make a widgets demonstration program called kwwidgets which demonstrates how to use and build renderer process widgets for kickwasm. I have yet to put it up at github.

I used version 16 to totally refactor my cwt application and it was so easy to do. I am really pleased with kickasm and it's tool chain.

#### Tests

I tested version 16.0.0 by running

* make install
* make test
* make prove

I want to do more tests by rebuilding

* kickwasm/examples/colors
* kickwasm/examples/spawnwidget
* [CRUD application](https://github.com/josephbudd/crud)
* [CWT application](https://github.com/josephbudd/cwt)

I want to update

* the kickwasm.wiki.
* the CRUD wiki.

## Kickwasm Summary

You begin your project by writing a kickwasm.yaml file. The kickwasm.yaml file defines your application's GUI. You run kickwasm and it generates the framework. The framework is your application's source code, written in GO, HTML, and CSS with the GUI built as defined in the kickwasm.yaml file.

The source code builds an application that runs as 2 separate processes. The 2 processes communicate with with each other by sending and receiving messages through channels.

1. The main process ( business logic ) is an http server running on the client.
1. The renderer process ( view logic) runs in the client's browser.

### Application Development Tools

The kickwasm tools take complicated framework refactoring operations and turn them into simple little code changes for you.

The main process will probably need to use the default local **Bolt database** and/or remote data bases and/or communicate with remote services. That is simple and quick with the tool **kickstore**.

The main process and renderer process will need to communicate with each other by sending messages through channels. That is simple and quick with the tool **kicklpc**.

The GUI may need to be refactored. And by **refactored** I mean adding, deleting and moving buttons, tabs, and panels with their respective GO packages that you coded, all over the place. That is safe, simple and quick with the tool **rekickwasm**.

## GUI design is simple: Buttons, Tabs, Panels

The GUI always begins with the buttons in the initial button pad. Those buttons show the user what actions the user can perform with the application. So in the kickwasm.yaml file you begin by defining the **buttons**.

### Buttons

Buttons always appear in a button pad.

Every button has 1 or more panels so in your kickwasm.yaml file you also define each button's panels. A button's panels form a panel group. In a panel group, showing one panel hides the other panels.

### Tabs

Tabs always appear in a tab bar.

Every tab has 1 or more markup panels so in your kickwasm.yaml file you also define each tab's panels. A tab's panels form a panel group. In a panel group, showing one panel hides the other panels.

### Panels

Panels always exist in a group. Even if the group only has one panel. In a panel group, only one panel is displayed. In a panel group, showing one panel hides the other panels.

A panel can have buttons or tabs or markup.

* **A panel with buttons** is rendered in the GUI as a button pad. **The framework controls the button pad for you.** When the user clicks on a button the button pad disapears and the user is shown the button's panel group.
* **A panel with tabs** is renderered in the GUI as a tab bar of tabs. **The framework controls the tab bar for you.** When the user clicks on a tab, the tab moves to the front of the other tabs. The front tab in the tab bar appears larger and infront of the other tabs and it's panel group is always displayed under it.
* **A panel with markup** renders it's own HTML in the GUI. **You control the markup panel with it's own HTML template file and it's own GO package.** In the GO package,
  * the **panelController** controls user input,
  * the **panelPresenter** writes to the panel,
  * the **panelMessenger** communicates with the main process.
  * The **panelGroup** has a show func for each panel in the panel group. Each show func makes it's panel visible and hides the other panels in the panel group.

## A hello world kickwasm.yaml file

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
    markup: <p>Hello world.</p>
```

## My step by step creation a CRUD with kickwasm

The [CRUD application](https://github.com/josephbudd/crud) was built with kickwasm. The CRUD WIKI is very important. It is where I detail the steps I took when I built the CRUD.

In the CRUD wiki, I begin with a complete plan for the full application.

I write a kickwasm.yaml file to define a small part of the GUI. Then use kickwasm to create that small part of the framework. Then, one markup panel at a time

1. I add a small amount of my own GO code for my markup panel package in the renderer process.
1. I add a small amount of my own GO code for what is needed in the main process.
1. I debug.

I repeately use all of the kickwasm tools to add another missing part to the application, writing and debuging a little more of my own code to the added part. Eventually, I have added all of the parts, all of my own code, and the application is complete.

## For more information check out

* The videos below.
* This kickwasm wiki linked to above.
* The [CRUD application](https://github.com/josephbudd/crud) and it's wiki.

## Installation

``` shell
~$ go get -u github.com/josephbudd/kickwasm
~$ cd ~/go/src/github.com/josephbudd/kickwasm
~/go/src/github.com/josephbudd/kickwasm$ make install
~/go/src/github.com/josephbudd/kickwasm$ make test
~/go/src/github.com/josephbudd/kickwasm$ make dependencies
~/go/src/github.com/josephbudd/kickwasm$ make prove
```

**make install** will build kickwasm and put it and it's tools into the go/bin.

**make test** will run the kickwasm unit tests. They only take a few seconds.

**make dependencies** will install the 3 GO packages that are required by the frameworks that you build with kickwasm. It will get them from github.com. So it might take a minute.

**make prove** will build and run the proofs. It will take a few minutes. The proofs make sure that rekickwasm refactors correctly.

There are currently 17 proofs. Each proof tests one or more ways of refactoring the framework's GUI. Each proof is a go program that

* Creates a kickwasm.yaml file.
* Runs kickwasm to build the framework ( application source code ).
* Adds the test code to the renderer process' ProveButtonPanel go package.
* Refactors the application's GUI using rekickwasm and another rekickwasm.yaml file.
* Builds and runs the application so that the tests run inside the ProveButtonPanel's go package.
* Logs the status of the run.
* Removes the files and folders it created if the run was successful.

Each application will popup up on your screen for less than a second while results are logged to the terminal. You can stop proofs at any time with ^c.

## A kickwasm hello world example step by step

### Hello world example STEP 1

Create a folder in your GO path.
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
The **-rp** flag signals to quick build the renderer process.
The **-mp** flag signals to build the main process.
The **-run** flag signals to run the executable.

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

### Spawnwidgets example

The spawnwidgets example is a simple tab spawning and unspawning application. It also uses a simple widget button as an example of widgets in spawned tab panels. It's README explains what is happening and where it is happening.

The video shows tabs being spawned and unspawned.

[![building and running the spawntabs example](https://i.vimeocdn.com/video/803691454.webp?mw=550&amp;mh=310&amp;q=70)](https://vimeo.com/351948165)
