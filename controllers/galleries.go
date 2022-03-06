package controllers

import (
	"instaphoto/views"
	"net/http"
)

func NewGalleries() *Galleries {
	return &Galleries{
		NewView: views.NewView("master", "galleries/new"),
	}
}

type Galleries struct {
	NewView *views.View
}

func (g *Galleries) New(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	g.NewView.Render(w, nil)
}
