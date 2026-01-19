package httpresponse

import (
	"encoding/json"
	"log"
	"net/http"
)

func HTTPResponseOK(w http.ResponseWriter, statusCode int, data interface{}) {

	// Source - https://stackoverflow.com/a/31622112
	// Posted by dm03514, modified by community. See post 'Timeline' for change history
	// Retrieved 2026-01-19, License - CC BY-SA 4.0

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatalf("error sending response: %v", err)
		}
	}

}

func HTTPResponseError(w http.ResponseWriter, statusCode int, err error) {

	// Source - https://stackoverflow.com/a/31622112
	// Posted by dm03514, modified by community. See post 'Timeline' for change history
	// Retrieved 2026-01-19, License - CC BY-SA 4.0

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
		log.Fatalf("error sending response: %v", err)
	}
}
