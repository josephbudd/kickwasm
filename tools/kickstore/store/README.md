**kickstore** is part of the kickwasm tool chain.

It is the **store**  tool for framework's created using kickwasm. It lets you add or remove stores. A store is a simple API to a table in the application's bolt database and it's record. By default, a store's API contains the following methods but you can add your own.

* Open
* Close
* Get
* GetAll
* Update
* Remove

Example: **kickstore -add Customer** will add the Customer store with it's functions and the empty Customer record.

* You finish defining the store's record so that it contains the information needed.
* You call the store's methods to update, remove, get etc. Add more methods if you want.

The crud.wiki page titled (Get The Add Contact Panel To Work)[https://github.com/josephbudd/crud.wiki/crud/Get-The-Add-Contact-Panel-To-Work.md] details how I used kickstore while making the CRUD application.

## To install

``` shell

$ go get -u github.com/josephbudd/kickstore
$ cd ~/go/src/github.com/josephbudd/kickstore
$ go install

```
