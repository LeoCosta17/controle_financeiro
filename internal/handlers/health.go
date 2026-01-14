package handlers

import (
	"net/http"
)

type HealtHanlder struct {
}

func (h *HealtHanlder) APIHealth(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("API on!"))
}
