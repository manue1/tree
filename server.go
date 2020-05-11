package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, I'm handling this\n")
	})

	log.Printf("listening on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
