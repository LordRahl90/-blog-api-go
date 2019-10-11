package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

//NewUsers - Return the routes and handlers for users
func NewUsers() *chi.Mux {
	r := chi.NewMux()
	r.Get("/", allUsers)
	return r
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	// db := services.Database{}
	w.Write([]byte("Welcome to the all users listings"))
}

func newUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the single user endpoint"))
}
