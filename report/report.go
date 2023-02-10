package report

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/Kaciis/go-reporter/pkg/controllers/report"
)

type Reporter struct {
	Iniciator string
	Url       string
}

func New() *Reporter {
	return &Reporter{}
}

func (r *Reporter) Info(message string) {

	send(report.ReportBody{
		Type:      "info",
		Message:   message,
		Iniciator: r.Iniciator,
	}, r.Url)

}

func (r *Reporter) Warning(message string) {

	send(report.ReportBody{
		Type:      "warning",
		Message:   message,
		Iniciator: r.Iniciator,
	}, r.Url)
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
