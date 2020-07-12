package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/l33m4n123/pet24/api/v1/database"
	"github.com/l33m4n123/pet24/api/v1/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetShelters handler for the GET /shelters/ route
func GetShelters(w http.ResponseWriter, r *http.Request) {
	shelters := models.GetAllShelters()
	json.NewEncoder(w).Encode(shelters)
}

// GetShelterByID retrieves a single shelter by ID from the database
func GetShelterByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	objectID, _ := primitive.ObjectIDFromHex(id)
	shelter := models.GetShelterByID(objectID)
	json.NewEncoder(w).Encode(shelter)
}

// AddNewShelter adds a new shelter to the database
func AddNewShelter(w http.ResponseWriter, r *http.Request) {
	var data models.Shelter
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&data)

	models.AddShelter(data)
}

// UpdateShelter adds a new shelter to the database
func UpdateShelter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	objectID, _ := primitive.ObjectIDFromHex(id)
	shelter := models.GetShelterByID(objectID)

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&shelter)

	filter := bson.D{primitive.E{Key: "_id", Value: objectID}}
	update := bson.D{
		primitive.E{
			Key: "$set", Value: bson.D{
				primitive.E{
					Key:   "name",
					Value: shelter.Name,
				},
				primitive.E{
					Key:   "contact",
					Value: shelter.Contact,
				},
				primitive.E{
					Key:   "address",
					Value: shelter.Address,
				},
				primitive.E{
					Key:   "updated_at",
					Value: time.Now(),
				},
			},
		},
	}

	updateResult, err := database.ShelterCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

// DeleteShelter adds a new shelter to the database
func DeleteShelter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	objectID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{primitive.E{Key: "_id", Value: objectID}}
	update := bson.D{
		primitive.E{
			Key: "$set", Value: bson.D{
				primitive.E{
					Key:   "deleted_at",
					Value: time.Now(),
				},
				primitive.E{
					Key:   "deleted",
					Value: true,
				},
			},
		},
	}

	updateResult, err := database.ShelterCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Matched %v documents and deleted %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}
