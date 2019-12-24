package timing

import (
	"context"
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
				msg := &message.TimeMainProcessToRenderer{
					Time: f,
				}
				sending <- msg
				timer.Reset(time.Second)
			}
		}
	}(ctx, sending)
}
