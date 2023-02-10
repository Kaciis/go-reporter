package routes

import (
	"log"

	"github.com/go-chi/chi"
)

type Router struct{}

func New() *Router {
	return &Router{}
}

func (r *Router) Routes(router chi.Router) {
	router.Route("/report", reportRouter)
	log.Println("Routes loaded")
}
