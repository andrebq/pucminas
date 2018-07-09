package ctrlc

import (
	"context"
	"os"
	"os/signal"
)

// Watch returns a context which is cancelled when SIGTERM/SIGKILL are received
func Watch(ctx context.Context) (context.Context, context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Kill, os.Interrupt)

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		<-ch
		signal.Stop(ch)
		signal.Reset(os.Kill, os.Interrupt)
		cancel()
	}()

	return ctx, cancel
}
