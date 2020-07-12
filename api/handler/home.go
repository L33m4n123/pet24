package handler

import (
	"net/http"
)

// HomeHandler handler for the home route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
}
