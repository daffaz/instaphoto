package main

import (
	"instaphoto/views"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var homeTemplate *views.View
var contactTemplate *views.View
var faqTemplate *views.View

func main() {
	homeTemplate = views.NewView("master", "./views/home.gohtml")
	contactTemplate = views.NewView("master", "./views/contact.gohtml")
	faqTemplate = views.NewView("master", "./views/faq.gohtml")

	mux := mux.NewRouter()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/contact", contact)
	mux.HandleFunc("/faq", faq)
	log.Println("Running in port :4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeTemplate.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactTemplate.Render(w, nil))
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(faqTemplate.Render(w, nil))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
