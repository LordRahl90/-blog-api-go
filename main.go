package main

import (
	"github.com/LordRahl90/blog-api-go/handlers"
	"github.com/LordRahl90/blog-api-go/services"
	"github.com/go-chi/chi"
)

func main() {
	s := services.NewServer()
	s.Router = registerRoutes()
	s.Start()
}

func registerRoutes() chi.Router {
	router := chi.NewRouter()
	site := handlers.NewSite()

	router.Route("/api/v1", func(rapi chi.Router) {
		rapi.Mount("/", site)
	})

	return router
}
