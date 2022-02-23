package main

import (
	"html/template"
	"log"
	"net/http"
)

var homeTemplate *template.Template
var contactTemplate *template.Template

func main() {
	var err error
	homeTemplate, err = template.ParseFiles(
		"./views/home.gohtml",
		"./views/layouts/footer.gohtml",
	)
	if err != nil {
		panic(err)
	}

	contactTemplate, err = template.ParseFiles(
		"./views/contact.gohtml",
		"./views/layouts/footer.gohtml",
	)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/contact", contact)
	log.Println("Running in port :4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}
