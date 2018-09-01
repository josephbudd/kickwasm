# kickwasm

A rapid development, desktop application generator for go. It creates a browser based GUI and most of the main process golang and renderer html, golang wasm and css code for you.

To summerize, it's https://github.com/josephbudd/kick with wasm instead of javascript

## Version alpha

kickwasm is working and still in a state of development.

I'm working on the contacts example right now. That will help me make any changes that I need to make to kickwasm. The contacts example does work but it doesn't do much of anything yet. I am still refactoring the javascript in kick/examples/contacts into golang in kickwasm/examples/contacts.

Also I haven't begun the wiki yet.

### The generated application.

* Imports the boltdb package.
* Imports my notjs package.
* Does not import my lpc package because I rewrote it and embedded it into the generated code. The renderer and the main process still call each other's funcs in a way similar to how they are done in kick.
* Each service now has it's own color. No more of the confusing color levels.
