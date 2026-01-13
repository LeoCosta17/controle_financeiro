package handlers

import "net/http"

func APIHealth(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("API on!"))
}
