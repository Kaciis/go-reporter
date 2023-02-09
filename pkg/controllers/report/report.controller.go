package report

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Report(w http.ResponseWriter, r *http.Request) {
	var rb *ReportBody

	byte_rb, err := io.ReadAll(r.Body)

	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	if len(byte_rb) == 0 {
		http.Error(w, "empty body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(byte_rb, &rb)

	if err != nil {
		log.Printf("Error unmarshalling body: %v", err)
		http.Error(w, "can't unmarshal body", http.StatusBadRequest)
		return
	}

	msg := GenerateReport(*rb)
	go SendMessage(msg)
}
