package main

import (
	"instaphoto/controllers"
	"instaphoto/views"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	homeView    *views.View
	contactView *views.View
)

func main() {
	homeView = views.NewView("master", "./views/home.gohtml")
	contactView = views.NewView("master", "./views/contact.gohtml")
	usersC := controllers.NewUsers()

	mux := mux.NewRouter()
	mux.HandleFunc("/", home).Methods("GET")
	mux.HandleFunc("/contact", contact).Methods("GET")
	mux.HandleFunc("/register", usersC.New).Methods("GET")
	mux.HandleFunc("/register", usersC.Create).Methods("POST")

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

func must(err error) {
	if err != nil {
		panic(err)
	}
}
