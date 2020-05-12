package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"text/template"
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
			fatalOnError(err, "failed to listen")
		}
	}()

	notifyOnShutdown(server)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tree := r.URL.Query().Get("favoriteTree")
	log.Printf("favoriteTree query: %s", tree)

	var body string
	if tree == "" {
		body = "Please tell me your favorite tree"
	} else {
		body = "It's nice to know that your favorite tree is a " + tree
	}

	t, err := template.New("index").Parse(indexTpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	content := struct {
		Title string
		Body  string
	}{
		Title: "Favorite Tree",
		Body:  body,
	}

	if err := t.Execute(w, content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

func fatalOnError(err error, msg string) {
	if err != nil {
		log.Fatalf(msg, err)
	}
}
