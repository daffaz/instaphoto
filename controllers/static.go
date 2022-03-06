package controllers

import "instaphoto/views"

type Static struct {
	Home    *views.View
	Contact *views.View
}

func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("master", "static/home"),
		Contact: views.NewView("master", "static/contact"),
	}
}
