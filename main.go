package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Hedwig7s/Burrowing-Classic/lib/networking"
)

func main() {
	var wg sync.WaitGroup
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	errCh := make(chan error, 1)

	server := networking.NewServer("0.0.0.0", 25564)

	defer server.Close()
	wg.Go(func() {
		errCh <- server.Start(ctx)
	})

	select {
	case <-ctx.Done():
		log.Println("shutdown signal received")

	case err := <-errCh:
		if err != nil {
			log.Printf("subsystem failed: %v", err)
		}
	}
	cancel()
	wg.Wait()
	log.Println("Shutdown complete")
}
