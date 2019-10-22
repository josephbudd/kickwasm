## Sept 15, 2019

Updated to kickwasm 9.0.0

### Sept 9, 2019

**kicklpc -l** now displays a sorted list of LPCs.

### Sept 2, 2019

Fixed an error message.

### August 26, 2019

Fixed a bug that wasn't testing the kickwasm version correctly.

### August 8, 2019

Updated to kickwasm version 6.0.0.

**kicklpc** is part of the kickwasm tool chain. It is the **Local Process Communications ( LPC )** tool for framework's created using kickwasm. It let's you add or remove LPC messages which are passed between the main process and the renderer process.

Example: **kicklpc -add UpdateCustomer** will add the empty message **UpdateCustomerRenderToMainProcess** and the other empty message **UpdateCustomerMainProcessToRenderer**.

* You complete the 2 messages defined in **domain/lpc/messages/UpdateCustomer.go** so that they can contain the dependencies you want.
* You add a message sender and a message receiver in your markup panel's Messenger.
* You finish the main process's message handler at **mainprocess/lpc/dispatch/UpdateCustomer.go** so that it processes the message and does what you need done with it.
* You send and receive those messages through the send and receive channels.

Example: **kicklpc -delete-forever UpdateCustomer**

* Will delete **domain/lpc/messages/UpdateCustomer.go**.
* Will delete **mainprocess/lpc/dispatch/UpdateCustomer.go**.

The crud.wiki page titled (Get The Add Contact Panel To Work)[https://github.com/josephbudd/crud.wiki/crud/Get-The-Add-Contact-Panel-To-Work.md] details how I used kicklpc while making the CRUD application.

## Installation

``` shell

$ go get -u github.com/josephbudd/kicklpc
$ cd ~/go/src/github.com/josephbudd/kicklpc
$ go install

```
