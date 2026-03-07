package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Hedwig7s/Burrowing-Classic/internal/networking/server"
	"github.com/Hedwig7s/Burrowing-Classic/internal/servercontext"
)

func main() {
	var wg sync.WaitGroup
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	errCh := make(chan error, 1)

	serverCtx := servercontext.DefaultServerContext()

	srv := server.NewServer("0.0.0.0", 25564, serverCtx)

	defer srv.Close()
	wg.Add(1)
	go func() {
		defer wg.Done()
		errCh <- srv.Start(ctx)
	}()

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
