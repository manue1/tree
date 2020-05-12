package handler

import (
	"log"
	"net/http"
	"text/template"
)

const (
	defaultBody = "Please tell me your favorite tree"
	treePrefix  = "It's nice to know that your favorite tree is a "
)

// Index is the "/" handler that renders the index HTML template
func Index(w http.ResponseWriter, r *http.Request) {
	tree := r.URL.Query().Get("favoriteTree")
	log.Printf("favoriteTree query: %s", tree)

	var body string
	if tree == "" {
		body = defaultBody
	} else {
		body = treePrefix + tree
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
