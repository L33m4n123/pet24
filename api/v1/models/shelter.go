package models

import (
	"context"
	"fmt"
	"time"

	"github.com/l33m4n123/pet24/api/v1/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Shelter struct to represent the data that is being saved in the database
type Shelter struct {
	// _ID represents the objectID in mongo
	ID primitive.ObjectID `bson:"_id,omitempty"`
	// Name of the shelter
	Name string `bson:"name,omitempty"`
	// Contact informatioon of the shelter
	Contact string `bson:"contact,omitempty"`
	// Address of the shelter
	Address *Address `bson:"address,omitempty"`
	// Deleted reflects wether or not this entry is/should be deleted
	Deleted bool `bson:"deleted"`
	// CreatedAt timestamp
	CreatedAt time.Time `bson:"created_at"`
	// DeletedAt - empty if it has not been deleted yet. Needed for soft-deletion
	DeletedAt time.Time `bson:"deleted_at,omitempty"`
	// UpdatedAt timestamp
	UpdatedAt time.Time `bson:"updated_at"`
}

// GetAllShelters returns a []Shelter containing all Shelters in the database
func GetAllShelters() []Shelter {
	var shelters []Shelter
	cursor, _ := database.ShelterCollection.Find(context.TODO(), bson.D{primitive.E{Key: "deleted", Value: false}})
	cursor.All(context.TODO(), &shelters)

	fmt.Println(time.Now())

	return shelters
}

// GetShelterByID returns a signle Shelter
func GetShelterByID(id primitive.ObjectID) Shelter {
	var shelter Shelter
	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	cursor := database.ShelterCollection.FindOne(context.TODO(), filter)
	cursor.Decode(&shelter)

	return shelter
}

// AddShelter adds a new entry to the database
func AddShelter(shelter Shelter) {
	shelter.CreatedAt = time.Now()
	shelter.UpdatedAt = shelter.CreatedAt

	database.ShelterCollection.InsertOne(context.TODO(), shelter)
}
