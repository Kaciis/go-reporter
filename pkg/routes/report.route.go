package routes

import (
	"github.com/Kaciis/go-reporter/pkg/controllers/report"
	"github.com/go-chi/chi"
)

func reportRouter(router chi.Router) {
	router.Post("/", report.Report)
}
