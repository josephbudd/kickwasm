## Intro

push demonstrates how to push LPC messages from the main process to the renderer as soon as the application starts.

The Init LPC message is automatically sent from the renderer to the main process in renderer/Main.go.

## Summary

When the renderer is up and running, after all of your markup panels packages have made their initial calls to the main process, the renderer sends an Init LPC message to the main process.

The main process's Init message handler is at mainprocess/lpc/dispatch/Init.go. func handleInit is shown below. In the push example, I use it to start the timing package's func Do which pushes the time to the renderer.

``` go

func handleInit(rxmessage *message.InitRendererToMainProcess, sending lpc.Sending, eojing lpc.EOJer, stores *store.Stores) {
  log.Println("dispatch.handleInit")
  timing.Do(sending, eojing, stores)
  return
}

```

The **timing** package is at mainprocess/services/timing and it's code is in the file timing.go. func Do in timing.go is shown below. It periodically pushes the time to the renderer using the Time LPC message.

``` go

// Do starts a go routine that periodically sends time to the renderer.
// Param sending is the send channel to the renderer process.
// Param eojing lpc.EOJer is the interface implementation which will give me the eoj channel to stop the go routine.
// Param stores *store.Stores contains the storage interface implementations. Param stores is only used here as an example and is not really needed.
func Do(sending lpc.Sending, eojing lpc.EOJer, stores *store.Stores) {
  log.Println("Do")
  go func(sending lpc.Sending, eojing lpc.EOJer, stores *store.Stores) {
    eojCh := eojing.NewEOJ()
    timer := time.NewTimer(time.Second)
    for {
      select {
      case <-eojCh:
        log.Println("timing.Do go returning.")
        if !timer.Stop() {
          <-timer.C
        }
        return
      case t := <-timer.C:
        f := t.Format(time.UnixDate)
        log.Println("timing.Do sending ", f)
        msg := &message.TimeMainProcessToRenderer{
          Time: f,
        }
        sending <- msg
        timer.Reset(time.Second)
      }
    }
  }(sending, eojing, stores)
  /*
    rxmessage is ignored because it only sinals that the renderer has been loaded.
  */
}

```

Because func Do has a go routine it uses Param eojing to get an eoj channel so that the go routine knows when the application is ending. You can see how this work's by watching the log.