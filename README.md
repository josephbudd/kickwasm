# kickwasm

A rapid development, desktop application generator for go. It creates a browser based GUI and most of the main process golang and renderer html, golang and css code for you.

To summerize, it's https://github.com/josephbudd/kick with wasm instead of javascript.

## Version alpha 0.2.0

### The generated application code has a TDD architecture

TDD ( test driven design ) is very simple. So simple you won't even notice it. This always begins in the main process's and render's **main.go** which intentionally abstracts away the production dependencies with interfaces defined at the domain level.

So, in order to test your code you will want to **inject dependencies**. That simply means that you will use the same interfaces defined at the domain level but replace the production dependencies with your own test dependencies.

#### The application is constructed with separate 2 processes

1. The renderer process with its main.go. It's compiled with **GOARCH=wasm GOOS=js go build -o app.wasm main.go panels.go**
2. The main process with it's main.go. It's compiled with **go build**

#### Code is physically and logically organized into 3 main levels

1. The domain folder contains domain ( application ) level logic.
2. The mainprocess folder contains the main process level logic.
3. The renderer folder contains the renderer level logic.

#### Inner Dependencies

Dependencies between the domain level and the main process level are abstracted away with the interfaces defined at the domain level.

Likewise, dependencies between the domain level and the renderer level are abstracted away with interfaces defined at the domain level.

### Kickwasm imports

* The yaml package at https://gopkg.in/yaml.v2

### Generated application imports

* Imports the boltdb package at https://github.com/boltdb/bolt.
* Imports my notjs package at https://github.com/josephbudd/kicknotjs.

### Wiki

I haven't begun the wiki yet.

### To do

* I'm rewriting the contact example.

* The wiki is not even started yet.
