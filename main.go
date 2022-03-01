package main

import (
	"instaphoto/views"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View
var signupView *views.View

func main() {
	homeView = views.NewView("master", "./views/home.gohtml")
	contactView = views.NewView("master", "./views/contact.gohtml")
	signupView = views.NewView("master", "./views/signup.gohtml")

	mux := mux.NewRouter()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/contact", contact)
	mux.HandleFunc("/signup", signup)

	log.Println("Running in port :4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(signupView.Render(w, nil))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
