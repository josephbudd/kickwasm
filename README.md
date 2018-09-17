# kickwasm

A rapid development, desktop application generator for go. It creates a browser based GUI and most of the main process golang and renderer html, golang and css code for you.

To summerize, it's https://github.com/josephbudd/kick with wasm instead of javascript.

## Version alpha 0.2.1

Still not stable but working. There are just some minor changes that I want to make.

### The generated application code has a TDD architecture

TDD ( test driven design ) is very simple to do in go. It involves the implementation of principles that already belong to go. So, it's just idiomatic go and unit tests. Kickwasm creates the idiomatic go and the code architecture. You create the unit tests if you want.

#### The application is constructed with separate 2 processes

1. The renderer process with its main.go. It's compiled with **GOARCH=wasm GOOS=js go build -o app.wasm main.go panels.go**
2. The main process with it's main.go. It's compiled with **go build**

#### Code is physically and logically organized into 3 main levels

1. The domain folder contains domain ( application ) level logic.
2. The mainprocess folder contains the main process level logic.
3. The renderer folder contains the renderer level logic.

### Kickwasm imports

* The yaml package at https://gopkg.in/yaml.v2

### Generated application imports

* Imports the boltdb package at https://github.com/boltdb/bolt.
* Imports my notjs package at https://github.com/josephbudd/kicknotjs.
  * The kicknotjs package is only needed by the framework. You don't have to use it in your own renderer code. I use it now because I wrote it.

### To do

The wiki yet.
The example video.

### Examples

The contacts example is finished so it works. I'll get a video posted as soon as I make one.