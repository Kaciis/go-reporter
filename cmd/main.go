package main

import (
	"log"
	"net/http"

	"github.com/Kaciis/go-reporter/pkg/controllers/app"
	"github.com/Kaciis/go-reporter/pkg/routes"
	"github.com/go-chi/chi"
)

func main() {

	app.Init()

	route := routes.New()

	r := chi.NewRouter()
	r.Route("/", route.Routes)
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
