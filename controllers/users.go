package controllers

import (
	"github.com/bodecci/use_golang/lenslocked/views"
	"net/http"
)

type Users struct {
	NewView *views.View
}

// NewUsers is used to create a new User controller
// This function will panic if the templates are not
// parsed correctly, and should only be used during
// initial setup.
func NewUsers() *Users {
	return &Users{
		NewView:views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}