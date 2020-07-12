package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/l33m4n123/pet24/api/database"
	"github.com/l33m4n123/pet24/api/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ShelterHandler handler for the home route
func ShelterHandler(w http.ResponseWriter, r *http.Request) {
	shelters := getShelters()
	json.NewEncoder(w).Encode(shelters)
}

func getShelters() []models.Shelter {
	var shelters []models.Shelter
	cursor, _ := database.ShelterCollection.Find(context.TODO(), bson.D{{}})
	cursor.All(context.TODO(), &shelters)

	return shelters
}
