package main

import (
	"instaphoto/views"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var homeTemplate *views.View
var contactTemplate *views.View

func main() {
	homeTemplate = views.NewView("master", "./views/home.gohtml")
	contactTemplate = views.NewView("master", "./views/contact.gohtml")

	mux := mux.NewRouter()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/contact", contact)
	log.Println("Running in port :4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeTemplate.Template.ExecuteTemplate(w, homeTemplate.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactTemplate.Template.ExecuteTemplate(w, contactTemplate.Layout, nil)
	if err != nil {
		panic(err)
	}
}
