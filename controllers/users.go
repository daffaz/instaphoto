package controllers

import (
	"fmt"
	"instaphoto/models"
	"instaphoto/views"
	"net/http"
)

// NewUsers is used to create a new Users controller.
// This function will panic if the templates are not
// parse correctly, and should only be used during
// initial setup.
func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView:   views.NewView("master", "users/new"),
		LoginView: views.NewView("master", "users/login"),
		us:        us,
	}
}

type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

type LoginForm struct {
	Email    string
	Password string
}

type Users struct {
	NewView   *views.View
	LoginView *views.View
	us        *models.UserService
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

// Login is used to render the form where a user can login
//
// GET /signup
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	if err := u.LoginView.Render(w, nil); err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "text/html")
}

func (u *Users) Authenticate(w http.ResponseWriter, r *http.Request) {
	var form LoginForm
	err := ParseForm(r, &form)
	if err != nil {
		panic(err)
	}

	user, err := u.us.Authenticate(form.Email, form.Password)

	if err != nil {
		switch err {
		case models.ErrNotFound:
			fmt.Fprintln(w, "Invalid email address.")
		case models.ErrInvalidPassword:
			fmt.Fprintln(w, "Invalid password provided.")
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	cookie := http.Cookie{
		Name:  "email",
		Value: user.Email,
	}
	http.SetCookie(w, &cookie)
	fmt.Fprintln(w, user)
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
	user := models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}
	if err = u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, user)
}
