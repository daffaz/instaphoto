package controllers

import (
	"fmt"
	"instaphoto/views"
	"net/http"
)

// NewUsers is used to create a new Users controller.
// This function will panic if the templates are not
// parse correctly, and should only be used during
// initial setup.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("master", "users/new"),
	}
}

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

type Users struct {
	NewView *views.View
}

// New is used to render the form where a user can
// create a new user account.
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "text/html")
}

// Create is used to process the signup form when a user
// tries to create a new user account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	err := ParseForm(r, &form)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "Email is =>", form.Email)
	fmt.Fprintln(w, "Password is =>", form.Password)
}
