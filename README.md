# kickwasm version 4.0.1 experimental because syscall/js is experimental

I didn't realize that the go syscall/js package was experimental. With go version 1.12 I found that out the hard way.

There fore **Kickwasm is EXPERIMENTAL**. Its current primary scope is to keep up with the changes in the EXPERIMENTAL go syscall/js package. Kickwasm is currently compatible with go version 1.12.

## Framework changes in version 4.0.1

Fixed **renderer/build.sh** so that it removes the old sitepack package before creating a new one.

## Framework changes in version 4.0.0

Now the entire application is compiled into a single executable that you can distribute.

So the build process hasn't changed. However there is a tiny change to how you will edit serve.go to handle your added url paths. It's really the same but with a call to a different func name.

## Framework changes in version 3.1.0

Added funcs to detect changes for rekickwasm that were previously missed in version 3.0.1.

## Framework changes in version 3.0.1

Added some missing documentation.

## Framework changes in version 3.0.0

* func RegisterCallBack and func RegisterEventCallBack were moved from the NotJS package to the ViewTools package in the renderer.
* Framework code has been refactored to correctly use the new func RegisterEventCallBack for event handlers, not func RegisterCallBack.

## Previous changes

* Lowercased panel package names. I also moved the call back funcs from tools to viewtools.

## I love kickwasm

Kickwams is a rapid development, desktop application framework generator for linux, windows and apple.

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

#### Build Step 1

In the first build step you build the renderer process. In the **renderer/** folder you run the shell script **build.sh**. **build.sh** builds in 3 steps.

1. Build the renderer's wasm byte code into the **site/** folder at **site/app.wasm**.
1. Write the source code for a new package made from the **site/** folder and the **settings.yaml** file.
1. Build the new package.

#### Build Step 2

Then, you build the main process. You end up with your entire application in one executable file.

Here is a build example using the colors example.

``` text

$ cd ~/src/github.com/josephbudd/kickwasm/examples/colors
$ cd renderer
$ ./build.sh
Building your wasm into ../site/app.wasm

Great! Your wasm has been compiled.

Now its time to write the source code for your new colorssitepack package.
The colorssitepack package is your applications renderer process.
( The stuff that gets loaded into the browser. )
This could take a while.
cd /home/nil/go/src/github.com/josephbudd/kickwasm/examples/colors
kickpack -o /home/nil/go/src/github.com/josephbudd/kickwasm/examples/colorssitepack ./site ./http.yaml

Finally! Now its time to build your new colorssitepack package.
cd /home/nil/go/src/github.com/josephbudd/kickwasm/examples/colorssitepack
go build

You've done it!
The package at /home/nil/go/src/github.com/josephbudd/kickwasm/examples/colorssitepack contains the files from your renderer process.

$ cd ..
$ go build
$ ./colors

```

### The framework imports

* [the boltdb package.](https://github.com/boltdb/bolt)
* [the yaml package.](https://gopkg.in/yaml.v2)
* [the gorilla websocket package.](https://github.com/gorilla/websocket)

## Installation

``` text

$ go get -u github.com/josephbudd/kickwasm
$ cd ~/go/src/github.com/josephbudd/kickwasm
$ go install
$ go get -u github.com/josephbudd/kickpack
$ cd ~/go/src/github.com/josephbudd/kickpack
$ go install

```

## Distribution

Once you build your application you can distribute it. Your entire application is contained in the executable file.

### For the examples/colors/ application that would be

1. the **examples/colors/colors** file which is the executable.

### For the examples/contacts/ application that would be

1. the **examples/contacts/contacts** file which is the executable.

## The examples

The examples/ folder contains 2 examples.

1. The colors example which is just a plain untouched framework. It was built with kickwasm version 4.0.0.
1. The contacts example which is a simple **C**reate **R**ead **U**pdate **D**elete application. It wasm built with kickwasm version 3.0.0. I will update the contacts example to kickwasm version 4.0.0 when I get the chance.

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

**Kickpack** converts files into a go source code package where they can be accesses from memory using the file's path. The framework's script at renderer/build.sh uses kickpack to build the application's entire renderer process into the application's sitepack package.

## Applications built with kickwasm

### CWT

* [![Learning Morse Code with CWT.](https://i.vimeocdn.com/video/772644525.jpg)](https://vimeo.com/328175343)