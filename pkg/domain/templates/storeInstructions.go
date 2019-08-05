package templates

// StoreInstructionsTXT is domain/store/instructions.txt
const StoreInstructionsTXT = `
ABOUT THE FILES IN domain/store

* domain/store/stores.go defines the type Stores struct.
The Stores has each one of your date stores as a member.
Package main's func buildStores() will build a pointer to a type Stores.
It will eventually end up being passed to your lpc message handlers
  so that they can use your data stores which are contained in the Stores.

ABOUT THE FILES IN domain/store/record/

{{ if gt (len .Stores) 0 }}{{ range .Stores}}* {{.}}.go contains the {{.}} record.
The file was created by kickstore. When it was created it, the record only had an ID.
You will want to edit the record definition so that it contains the members that you need.
{{ end }}{{ else }}* There are no files because you haven't added any data stores yet.{{ end }}

ABOUT THE FILES IN domain/store/storer/

{{ if gt (len .Stores) 0 }}{{ range .Stores}}* {{.}}.go contains the {{.}}Storer interface. The interface defines the behavior ( API ) of the {{.}} data store.
The file was created by kickstore.
You may want to edit the interface definition so that it more closely meets you needs.
If you do you may also need to edit it's implementation in domain/store/storing/{{.}}.go.
{{ end }}{{ else }}* There are no files because you haven't added any data stores yet.{{ end }}

ABOUT THE FILES IN domain/store/storing/

{{ if gt (len .Stores) 0 }}{{ range .Stores}}* {{.}}.go contains the {{.}}BoltDB. interface.
The {{.}}BoltDB is the implementation of the {{.}}Storer interface for the bolt database.
The file was created by kickstore.
You may want to edit the implementation so that it more closely meets you needs.
If you do you may also need to edit it's interface definition in domain/store/storer/{{.}}.go.
{{ end }}{{ else }}* There are no files because you haven't added any data stores yet.{{ end }}

ABOUT THE FILES IN package main.

* stores.go in func buildStores(), builds a store.Stores. A struct store.Stores is defined in domain/store/stores.go.

* main.go gets the store.Stores from func buildStores() and passes it on to your LPC message handlers so that they can get data from and put data into your data stores.

MANAGING DATA STORES WITH kickstore.

* Use kickstore in this application's root folder:
  $ cd github.com/josephbudd/mptabs/

* Listing all of the messages:
  $ kickstore -l
  1. kickstore would
    * Display the names of each data store.

* Adding a data store:
  $ kickstore -add Customer
  1. kickstore would
    * Update the file domain/store/instructions.txt.
    * Update the file domain/store/stores.go.
    * Add the file domain/store/record/Customer.go.
	  * Add the file domain/store/storer/Customer.go.
	  * Add the file domain/store/storing/CustomerBoltDB.go.
    * Update the file stores.go.
  2. You would need to
	  * Complete the store record definition in domain/store/record/Customer.go.
	  * You may want to modify the interface in domain/store/storer/Customer.go.
	    If you do you may also need to modify the implementation in domain/store/storing/CustomerBoltDB.go.

* Deleting a data store:
  $ kickstore -delete-forever Customer
  1. kickstore would
    * Update the file domain/store/instructions.txt.
    * Update the file domain/store/stores.go.
    * Delete the file domain/store/record/Customer.go.
    * Delete the file domain/store/storer/Customer.go.
    * Delete the file domain/store/storing/CustomerBoltDB.go.
    * Update the file stores.go.
`