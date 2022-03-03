package main

import (
	"instaphoto/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()

	mux := mux.NewRouter()
	mux.Handle("/", staticC.Home).Methods("GET")
	mux.Handle("/contact", staticC.Contact).Methods("GET")
	mux.HandleFunc("/register", usersC.New).Methods("GET")
	mux.HandleFunc("/register", usersC.Create).Methods("POST")

	log.Println("Running in port :4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
