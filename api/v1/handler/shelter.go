package handler

import (
	"encoding/json"
	"net/http"

	"github.com/l33m4n123/pet24/api/models"
)

// ShelterHandler handler for the home route
func ShelterHandler(w http.ResponseWriter, r *http.Request) {
	shelters := models.GetAllShelters()
	json.NewEncoder(w).Encode(shelters)
}
