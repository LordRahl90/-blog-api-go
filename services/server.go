package services

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

//Server - Struct to contain the server's necessity
type Server struct {
	Router chi.Router
}

//NewServer Create a new Server by initializing a new chi route
func NewServer() *Server {
	return &Server{}
}

//MountRoutes - This function is used to setup different routes
func (s *Server) MountRoutes(path string, r *chi.Mux) {
	s.Router.Mount(path, r)
}

//Start - This function starts the server.
func (s *Server) Start() {
	if s.Router == nil {
		log.Fatal("Please setup the routes for the server")
	}
	fmt.Println("Server Starting on Port 5000")
	server := &http.Server{
		Addr:         ":5000",
		Handler:      s.Router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Could not listen on thee specified address")
	}

	log.Println("Server Stopped")
}
