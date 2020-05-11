package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	server := &http.Server{
		Handler: mux,
		Addr:    ":8000",
	}

	log.Printf("listening on :8000")
	log.Fatal(server.ListenAndServe())
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Request with query: %s\n", req.URL.Query())
	io.WriteString(w, "Please tell me your favorite tree\n")
}
