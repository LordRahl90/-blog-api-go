package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/LordRahl90/blog-api-go/utility"
	"github.com/go-chi/chi"
)

//NewSite - this maps all the functions in this package to handlers
func NewSite() *chi.Mux {
	r := chi.NewMux()
	r.Get("/", home)
	r.Get("/status", status)

	return r
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func status(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"server":  "1234",
		"time":    time.Now(),
		"day":     time.Now().Day(),
		"month":   time.Now().Month(),
		"year":    time.Now().Year(),
		"Weekday": time.Now().Weekday(),
	}
	res, err := utility.APIResponse(200, true, "Server Running Conveniently", data)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
