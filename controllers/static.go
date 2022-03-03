package controllers

import "instaphoto/views"

type Static struct {
	Home    *views.View
	Contact *views.View
}

func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("master", "./views/static/home.gohtml"),
		Contact: views.NewView("master", "./views/static/contact.gohtml"),
	}
}
