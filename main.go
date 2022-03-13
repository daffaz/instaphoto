package main

import (
	"instaphoto/controllers"
	"instaphoto/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	dsn := "host=CONFIDENTAL user=CONFIDENTAL password=CONFIDENTAL dbname=CONFIDENTAL port=CONFIDENTAL sslmode=disable"
	us, err := models.NewUserService(dsn)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.AutoMigrate()
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(us)

	mux := mux.NewRouter()
	mux.Handle("/", staticC.Home).Methods("GET")
	mux.Handle("/contact", staticC.Contact).Methods("GET")
	// Register
	mux.HandleFunc("/register", usersC.New).Methods("GET")
	mux.HandleFunc("/register", usersC.Create).Methods("POST")
	// Login
	mux.HandleFunc("/login", usersC.Login).Methods("GET")
	mux.HandleFunc("/login", usersC.Authenticate).Methods("POST")

	log.Println("Running in port :4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}
