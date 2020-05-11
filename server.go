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

func indexHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Request with query: %s\n", req.URL.Query())

	t, err := template.New("index").Parse(indexTpl)
	fatalOnError(err, "failed to parse template")

	content := struct {
		Title string
		Body  string
	}{
		Title: "Favorite Tree",
		Body:  blank,
	}
	err = t.Execute(w, content)
	fatalOnError(err, "failed to apply data to template")
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
