package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Kaciis/go-reporter/pkg/controllers/app"
	"github.com/Kaciis/go-reporter/pkg/controllers/report"
	"github.com/Kaciis/go-reporter/pkg/routes"
	"github.com/go-chi/chi"
)

type Reporter struct {
	Iniciator string
	Url       string
}

func main() {

	app.Init()

	route := routes.New()

	r := chi.NewRouter()
	r.Route("/", route.Routes)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func (r *Reporter) Error(err error) {

	send(report.ReportBody{
		Type:      "error",
		Message:   err.Error(),
		Iniciator: r.Iniciator,
	}, r.Url)

}

func send(report report.ReportBody, url string) error {

	if url == "" {
		url = os.Getenv("REPORT_URL")
	}

	s_report, err := json.Marshal(report)

	if err != nil {
		return err
	}

	_, err = http.Post(url, "application/json", strings.NewReader(string(s_report)))

	if err != nil {
		return err
	}

	return nil
}
