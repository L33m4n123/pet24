package models

import (
	"context"

	"github.com/l33m4n123/pet24/api/v1/database"
	"go.mongodb.org/mongo-driver/bson"
)

// Shelter struct to represent the data that is being saved in the database
type Shelter struct {
	// Name of the shelter
	Name string `json:"name"`
	// Contact informatioon of the shelter
	Contact string `json:"contact"`
	// Address of the shelter
	Address *Address `json:"address"`
}

// GetAllShelters returns a []Shelter containing all Shelters in the database
func GetAllShelters() []Shelter {
	var shelters []Shelter
	cursor, _ := database.ShelterCollection.Find(context.TODO(), bson.D{{}})
	cursor.All(context.TODO(), &shelters)

	return shelters
}
