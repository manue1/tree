package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/manue1/favorite-tree/pkg/handler"
)

func main() {
	const port = ":8000"

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Index)

	server := &http.Server{
		Handler: mux,
		Addr:    port,
	}

	go func() {
		log.Printf("listening on 127.0.0.1%s", port)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal("server failed to listen", err)
		}
	}()

	notifyOnShutdown(server)
}

// notifyOnShutdown listens for the listed system signals and shuts down the server gracefully
func notifyOnShutdown(s *http.Server) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s.SetKeepAlivesEnabled(false)
	s.Shutdown(ctx)

	log.Printf("shutting down gracefully")
	os.Exit(0)
}
