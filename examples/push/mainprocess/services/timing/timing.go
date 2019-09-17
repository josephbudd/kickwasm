package timing

import (
	"log"
	"time"

	"github.com/josephbudd/kickwasm/examples/push/domain/lpc/message"
	"github.com/josephbudd/kickwasm/examples/push/domain/store"
	"github.com/josephbudd/kickwasm/examples/push/mainprocess/lpc"
)

// Do starts a go routine that periodically sends time to the renderer.
// Param sending is the send channel to the renderer process.
// Param eojing lpc.EOJer is the interface implementation which will give me the eoj channel to stop the go routine.
// Param stores *store.Stores contains the storage interface implementations.
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
