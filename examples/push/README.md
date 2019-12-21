# What is there to learn from the push example

## How to start a go routine in the main process once the renderer process is up and running

### The Init message from the renderer process

The init message is defined in **examples/push/domain/lpc/message/Init.go**. When the main process receives the **Init** message it knows that the renderer process is up and running. No information is sent in the message because the meaning is implied.

## The main process

The init message is processed by the main process in **examples/push/mainprocess/lpc/dispatch/Init.go** which is shown below. **func handleInit**, ignores the message contents and simply calls **timing.Do(ctx, sending)** to start the timing service. Notice that timing.Do is passed the context and the channel used to send messages to the renderer process.

```go

package dispatch

import (
  "context"

  "github.com/josephbudd/kickwasm/examples/push/domain/lpc/message"
  "github.com/josephbudd/kickwasm/examples/push/domain/store"
  "github.com/josephbudd/kickwasm/examples/push/mainprocess/lpc"
  "github.com/josephbudd/kickwasm/examples/push/mainprocess/services/timing"
)

/*
  YOU MAY EDIT THIS FILE.

  Rekickwasm will preserve this file for you.
  Kicklpc will not edit this file.

*/

// handleInit is the *message.InitRendererToMainProcess handler.
//   The InitRendererToMainProcess message signals that
//   * the renderer process is up and running,
//   * the main process may push messages to the renderer process.
//   The message is sent from renderer/Main.go which you can edit.
// handleInit's response back to the renderer is the *message.InitMainProcessToRenderer.
// Param ctx is the context. if <-ctx.Done() then the main process is shutting down.
// Param rxmessage *message.InitRendererToMainProcess is the message received from the renderer.
// Param sending is the channel to use to send a *message.InitMainProcessToRenderer message back to the renderer.
// Param stores is a struct the contains each of your stores.
// Param errChan is the channel to send the handler's error through since the handler does not return it's error.
func handleInit(ctx context.Context, rxmessage *message.InitRendererToMainProcess, sending lpc.Sending, stores *store.Stores, errChan chan error) {
  timing.Do(ctx, sending)
  return
}


```

### The timing service in the main process

The main process has a service that sends the time to the renderer process every second. The service is at **mainprocess/services/timing.go** shown below. **func Do** starts the service.

The service is simply a go routine which

1. sends the time to the renderer process through the sending channel every second.
1. returns if the context is done.

```go

package timing

import (
  "context"
  "log"
  "time"

  "github.com/josephbudd/kickwasm/examples/push/domain/lpc/message"
  "github.com/josephbudd/kickwasm/examples/push/mainprocess/lpc"
)

// Do starts a go routine that periodically sends time to the renderer.
// Param ctx is the context. if <-ctx.Done() then the main process is shutting down.
// Param sending is the send channel to the renderer process.
func Do(ctx context.Context, sending lpc.Sending) {
  go func(ctx context.Context, sending lpc.Sending) {
    timer := time.NewTimer(time.Second)
    for {
      select {
      case <-ctx.Done():
        // The application has ended.
        // Stop everything and return.
        timer.Stop()
        return
      case t := <-timer.C:
        f := t.Format(time.UnixDate)
        log.Printf("timing.Do's go func is sending %q", f)
        msg := &message.TimeMainProcessToRenderer{
          Time: f,
        }
        sending <- msg
        timer.Reset(time.Second)
      }
    }
  }(ctx, sending)
}

```

## The renderer process

In this simple application there is only 1 markup panel. The panel is named **PushPanel**. It's template file is at **examples/push/site/templates/PushButton/PushPanel.tmpl** It's go package is at **examples/push/rendererprocess/panels/PushButton/PushPanel/**.

The **PushPanel** receives and displays the time. The panel messenger receives the message and the panel presenter displays it.

### The renderer process receiving the Time message

Below is the PushPanel's go package's messenger. A panel's messenger communicates with the main process.

1. In **func dispatchMessages**, the messenger receives the **TimeMainProcessToRenderer** message and distributes it to **func timeRX**.
1. **timeRX** passes the message time string to the presenter's **displayTimeSpan** func which displays the time string.

```go

// +build js, wasm

package pushpanel

import (
  "github.com/josephbudd/kickwasm/examples/push/domain/lpc/message"
  "github.com/josephbudd/kickwasm/examples/push/rendererprocess/api/display"
)

/*

  Panel name: PushPanel

*/

// panelMessenger communicates with the main process via an asynchrounous connection.
type panelMessenger struct {
  group      *panelGroup
  presenter  *panelPresenter
  controller *panelController

  /* NOTE TO DEVELOPER. Step 1 of 4.

  // 1.1: Declare your panelMessenger members.

  // example:

  state uint64

  */
}

/* NOTE TO DEVELOPER. Step 2 of 4.

// 2.1: Define your funcs which send a message to the main process.
// 2.2: Define your funcs which receive a message from the main process.

// example:

import "github.com/josephbudd/kickwasm/examples/push/domain/store/record"
import "github.com/josephbudd/kickwasm/examples/push/domain/lpc/message"
import "github.com/josephbudd/kickwasm/examples/push/rendererprocess/api/display"

// Add Customer.

func (messenger *panelMessenger) addCustomer(r *record.Customer) {
  msg := &message.AddCustomerRendererToMainProcess{
    UniqueID: messenger.uniqueID,
    Record:   record,
  }
  sendCh <- msg
}

func (messenger *panelMessenger) addCustomerRX(msg *message.AddCustomerMainProcessToRenderer) {
  if msg.UniqueID == messenger.uniqueID {
    if msg.Error {
      display.Error(msg.ErrorMessage)
      return
    }
    // no errors
    display.Success("Customer Added.")
  }
}

*/

func (messenger *panelMessenger) timeRX(msg *message.TimeMainProcessToRenderer) {
  if msg.Error {
    display.Error(msg.ErrorMessage)
    return
  }
  messenger.presenter.displayTimeSpan(msg.Time)
}

// dispatchMessages dispatches LPC messages from the main process.
// It stops when it receives on the eoj channel.
func (messenger *panelMessenger) dispatchMessages() {
  go func() {
    for {
      select {
      case <-eojCh:
        return
      case msg := <-receiveCh:
        // A message sent from the main process to the renderer.
        switch msg := msg.(type) {

        /* NOTE TO DEVELOPER. Step 3 of 4.

        // 3.1:   Remove the default clause below.
        // 3.2.a: Add a case for each of the messages
        //          that you are expecting from the main process.
        // 3.2.b: In that case statement, pass the message to your message receiver func.

        // example:

        import "github.com/josephbudd/kickwasm/examples/push/domain/lpc/message"

        case *message.AddCustomerMainProcessToRenderer:
          messenger.addCustomerRX(msg)

        */

        case *message.TimeMainProcessToRenderer:
          messenger.timeRX(msg)
        }
      }
    }
  }()

  return
}

// initialSends sends the first messages to the main process.
func (messenger *panelMessenger) initialSends() {

  /* NOTE TO DEVELOPER. Step 4 of 4.

  //4.1: Send messages to the main process right when the app starts.

  // example:

  import "github.com/josephbudd/kickwasm/examples/push/domain/data/loglevels"
  import "github.com/josephbudd/kickwasm/examples/push/domain/lpc/message"

  msg := &message.LogRendererToMainProcess{
    Level:   loglevels.LogLevelInfo,
    Message: "Started",
  }
  sendCh <- msg

  */
}

```

### The renderer process displaying the time

Below is the PushPanel's go package's presenter. A panel's presenter outputs to the GUI.

1. func **displayTimeSpan** displays the time.

```go

// +build js, wasm

package pushpanel

import (
  "errors"
  "fmt"

  "github.com/josephbudd/kickwasm/examples/push/rendererprocess/api/markup"
)

/*

  Panel name: PushPanel

*/

// panelPresenter writes to the panel
type panelPresenter struct {
  group      *panelGroup
  controller *panelController
  messenger  *panelMessenger

  /* NOTE TO DEVELOPER: Step 1 of 3.

  // Declare your panelPresenter members here.

  // example:

  import "github.com/josephbudd/kickwasm/examples/push/rendererprocess/api/markup"

  editCustomerName *markup.Element

  */

  timeSpan *markup.Element
}

// defineMembers defines the panelPresenter members by their html elements.
// Returns the error.
func (presenter *panelPresenter) defineMembers() (err error) {

  defer func() {
    if err != nil {
      err = fmt.Errorf("(presenter *panelPresenter) defineMembers(): %w", err)
    }
  }()

  /* NOTE TO DEVELOPER. Step 2 of 3.

  // Define your panelPresenter members.

  // example:

  // Define the edit form's customer name input field.
  if presenter.editCustomerName = document.ElementByID("editCustomerName"); presenter.editCustomerName == nil {
    err = fmt.Errorf("unable to find #editCustomerName")
    return
  }

  */

  // Define timeSpan output field.
  if presenter.timeSpan = document.ElementByID("timeSpan"); presenter.timeSpan == nil {
    err = errors.New("unable to find #timeSpan")
    return
  }

  return
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your panelPresenter functions.

// example:

// displayCustomer displays the customer in the edit customer form panel.
func (presenter *panelPresenter) displayCustomer(record *types.CustomerRecord) {
  presenter.editCustomerName.SetValue(record.Name)
}

*/

// displayTimeSpan displays time.
func (presenter *panelPresenter) displayTimeSpan(s string) {
  presenter.timeSpan.SetInnerText(s)
}

```
