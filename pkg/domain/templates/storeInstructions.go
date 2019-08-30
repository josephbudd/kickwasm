package templates

// StoreInstructionsTXT is domain/store/instructions.txt
const StoreInstructionsTXT = `
ABOUT THE FILES IN domain/store/

* domain/store/stores.go defines the type Stores struct.
The type Stores has each one of your data stores as a member.
  * Each local bolt data store.
  * Each remote data store.
Package main's func buildStores() will build a pointer to a type Stores.
It will eventually end up being passed to your lpc message handlers
  so that they can use your data stores which are contained in the Stores.

ABOUT THE FILES IN domain/store/record/

{{ if (gt (len .BoltStores) 0) or (gt (len .RemoteRecords) 0) }}{{ range .BoltStores}}* {{.}}.go contains the {{.}} struct which is the {{.}} record.
The file was created by kickstore for the application's local bolt database. When it was created, the struct's only member was "ID uint64".
You will want to edit the record definition so that it contains the members that you need.
{{ end }}{{ range .RemoteRecords}}* {{.}}.go contains the {{.}} struct which is the {{.}} record.
The file was created by kickstore for one of the application's remote databases. When it was created, the struct had no members.
You will want to edit the record definition so that it contains the members that you need.
{{ end }}{{ else }}* There are no files because you haven't added any local bolt data stores or remote database records yet.{{ end }}

ABOUT THE FILES IN domain/store/storer/

{{ if (gt (len .BoltStores) 0) or (gt (len .RemoteDBs) 0) }}{{ range .BoltStores}}* {{.}}.go contains the {{.}}Storer interface which defines the local bolt {{.}} data store's ( API ).
The file was created by kickstore with a complete interface definition.
You may need to edit the interface definition so that it more closely meets your needs.
Likewise may also need to edit it's implementation in domain/store/storing/{{.}}.go.
{{ end }}{{ range .RemoteDBs}}* {{.}}.go contains the {{.}}Storer interface which defines the remote {{.}} database's ( API ).
The file was created by kickstore with a minimal interface definition.
You will need to edit the interface definition so that it meets your needs.
Likewise, you will also need to edit it's implementation in domain/store/storing/{{.}}.go.
{{ end }}{{ else }}* There are no files because you haven't added any local bolt data stores or remote databases yet.{{ end }}

ABOUT THE FILES IN domain/store/storing/

{{ if (gt (len .BoltStores) 0) or (gt (len .RemoteDBs) 0) }}{{ range .BoltStores}}* {{.}}.go contains the {{.}}BoltDB struct and it's implementation of the local bolt {{.}} data store's ( API ).
The file was created by kickstore with a complete interface implementation.
The {{.}}BoltDB struct implements the {{.}}Storer interface defined in domain/store/storer/{{.}}.go.
{{ end }}{{ range .RemoteDBs}}* {{.}}.go contains the {{.}}DB struct and it's implementation of the remote {{.}} database's ( API ).
The file was created by kickstore with a minimal interface implementation.
You must edit the file because {{.}}DB struct must implement the {{.}}Storer interface defined in domain/store/storer/{{.}}.go.
{{ end }}{{ else }}* There are no files because you haven't added any local bolt data stores or remote databases yet.{{ end }}

ABOUT THE FILES IN package main.

* stores.go in func buildStores(), builds a store.Stores. A struct store.Stores is defined in domain/store/stores.go.

* main.go gets the store.Stores from func buildStores() and passes it on to your LPC message handlers so that they can get data from and put data into your data stores.

MANAGING DATA STORES WITH kickstore.

* Use kickstore in this application's root folder:
  $ cd github.com/josephbudd/mptabs/

* Listing all of the messages:
  $ kickstore -l
  1. kickstore would
    * Display the names of each bolt data store.
    * Display the names of each remote database.
    * Display the names of each remote database record.

* Adding a local bolt data store:
  $ kickstore -add Customer
  1. kickstore would
    * Update the file domain/store/instructions.txt.
    * Update the file domain/store/stores.go.
	  * Add the file domain/store/storer/Customer.go.
	  * Add the file domain/store/storing/CustomerBoltDB.go.
    * Update the file stores.go.
  2. You would need to
	  * Complete the store record definition in domain/store/record/Customer.go.
	  * You may want to modify the interface in domain/store/storer/Customer.go.
	    If you do you may also need to modify the implementation in domain/store/storing/CustomerBoltDB.go.

* Deleting a local bolt data store:
  $ kickstore -delete-forever Customer
  1. kickstore would
    * Update the file domain/store/instructions.txt.
    * Update the file domain/store/stores.go.
    * Delete the file domain/store/storer/Customer.go.
    * Delete the file domain/store/storing/CustomerBoltDB.go.
    * Update the file stores.go.

* Adding a remote api:
  $ kickstore -add-remote-api Inventory
  1. kickstore would
    * Update the file domain/store/instructions.txt.
    * Update the file domain/store/stores.go.
	  * Add the file domain/store/storer/Inventory.go.
	  * Add the file domain/store/storing/Inventory.go.
    * Update the file stores.go.
  2. You would need to
	  * Complete the API definition in domain/store/storer/Inventory.go.
    * Complete the API implementation in domain/store/storing/InventoryDB.go.

* Deleting a remote data base:
  $ kickstore -delete-forever-remote-db Inventory
  1. kickstore would
    * Update the file domain/store/instructions.txt.
    * Update the file domain/store/stores.go.
    * Delete the file domain/store/storer/Inventory.go.
    * Delete the file domain/store/storing/InventoryDB.go.

* Adding a remote data base record:
  $ kickstore -add-remote-record Product
  1. kickstore would
    * Update the file domain/store/instructions.txt.
    * Add the file domain/store/record/Product.go.
  2. You would need to
    * Complete the remote record definition in domain/store/record/Product.go.
  
* Deleting a remote data base record:
  $ kickstore -delete-forever-remote-record Product
  1. kickstore would
    * Update the file domain/store/instructions.txt.
    * Delete the file domain/store/record/Product.go.
`
