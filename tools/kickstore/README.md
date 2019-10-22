
## Sept 2, 2019:

Finished the changes started below and a minor error message correction.

## August 30, 2019:

Updated to kickwasm version 7.0.2.

1. **kickstore** is how you manage your application's data storage model.
   * You can add local bolt data stores. A local bolt store is an API and a record. It runs locally in the application's bolt database.
   * You can add remote APIs. A remote API can be for a remote database or maybe a remote server that has some home you want to use. You must complete a remote API's functionality.
   * You can add remote records. A remote record is just a struct representing a record in a remote database or data for sending or receiving data from a remote server.
   * Example 1: **kickstore -add Customer** would add the local bolt **Customer** store with it's API and the empty **Customer** record.
     1. You would complete the the store's record definition in **domain/store/record/Customer.go** so that it contains the information needed.
     1. If you want to modify the stores API, then you would edit it's interface in **domain/store/storer/Customer.go** and it's implementation in **domain/store/storing/Customer.go**.
     1. In your code you would use the store's record and call the store's methods to update, remove, get etc.
   * Example 2: **kickstore -add-remote-api Inventory** would add the remote API **Inventory**.
     1. You would complete the remote database's API interface in **domain/store/storer/Inventory.go** and the remote database's API implementation in **domain/store/storing/Inventory.go**.
     1. In your code you would call the remote database's API methods that you wrote.
   * Example 3: **kickstore -add-remote-record Item** would add the remote database record **Item**.
     1. You would complete the **Item** record definition in **domain/store/record/Item.go** so that it contains the information needed.
     1. In your code you could use the remote record **Item** and call the remote database **Inventory** API methods that you wrote.

The crud.wiki page titled (Get The Add Contact Panel To Work)[https://github.com/josephbudd/crud.wiki/crud/Get-The-Add-Contact-Panel-To-Work.md] details how I used kickstore while making the CRUD application.

Below is a screen dump of the command **kickstore** without any flags.

``` shell

$ kickstore

  Run kickstore in your application's source code root folder.
  The one that contains the folders ".kickwasm" and ".kickstore".

  
  -add string
      names the local bolt store to add
  -add-remote-api string
      names the remote API to add
  -add-remote-record string
      names the remote database record to add
  -delete-forever string
      names the local bolt store to delete
  -delete-forever-remote-api string
      names the remote API to delete
  -delete-forever-remote-record string
      names the remote database record to delete
  -l  Lists the current stores
  -v  version

```

## Installation

``` shell

$ go get -u github.com/josephbudd/kickstore
$ cd ~/go/src/github.com/josephbudd/kickstore
$ go install

```
