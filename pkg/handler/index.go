package handler

import (
	"log"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
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

	if err := t.Execute(w, struct {
		Title string
		Body  string
	}{
		Title: "Favorite Tree",
		Body:  body,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
