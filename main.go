package main

import (
	"log"

	"github.com/LordRahl90/blog-api-go/handlers"
	"github.com/LordRahl90/blog-api-go/services"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	s := services.NewServer()
	s.Router = registerRoutes()
	s.DB.DB = registerDB()
	s.Start()
}

func registerRoutes() chi.Router {
	router := chi.NewRouter()
	site := handlers.NewSite()
	user := handlers.NewUsers()

	router.Route("/api/v1", func(rapi chi.Router) {
		rapi.Mount("/", site)
		rapi.Mount("/users", user)
	})

	return router
}

func registerDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "blog_api.db")
	if err != nil {
		log.Fatal(err)
	}

	//Automigrate stuffs
	db.AutoMigrate(&services.User{})

	return db
}
