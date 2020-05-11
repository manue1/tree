package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	server := &http.Server{
		Handler: mux,
		Addr:    ":8000",
	}

	go func() {
		log.Printf("listening on :8000")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("failed to listen", err)
		}
	}()

	notifyOnShutdown(server)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Request with query: %s\n", req.URL.Query())
	io.WriteString(w, "Please tell me your favorite tree\n")
}

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
