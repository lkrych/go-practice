package controllers

import (
	"flatphoto.com/views"
)

func NewStatic() *Static {
	return &Static{
		Home: views.NewView(
			"bootstrap", "views/static/home.gohtml"),
		Contact: views.NewView(
			"bootstrap", "views/static/contact.gohtml"),
		Faq: views.NewView(
			"bootstrap", "views/static/faq.gohtml"),
		FourOhfour: views.NewView(
			"bootstrap", "views/static/404.gohtml"),
	}
}

type Static struct {
	Home       *views.View
	Contact    *views.View
	Faq        *views.View
	FourOhfour *views.View
}
