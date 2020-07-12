package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// ShelterCollection contains a reference to its named collection in the mongoDB
	ShelterCollection *mongo.Collection
	// PetCollection contains a reference to its named collection in the mongoDB
	PetCollection *mongo.Collection
	// UserCollection contains a reference to its named collection in the mongoDB
	UserCollection *mongo.Collection

	client   *mongo.Client
	database *mongo.Database
)

// SetupConnection makes sure the api connects sucessfully to the database
func SetupConnection(connectionString string) {
	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	database = client.Database("pet24")

	ShelterCollection = database.Collection("shelters")
	PetCollection = database.Collection("pets")
	UserCollection = database.Collection("users")
}

// CloseConnection Closes the connection before API Shutdown
func CloseConnection() {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to MongoDB closed.")
}
