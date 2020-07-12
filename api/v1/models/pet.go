package models

// Pet struct to represent the data that is stored in the db
type Pet struct {
	// Name of the pet
	Name string `json:"name"`
	// Race of the pet - aka Dog, Cat, Camel, and so on
	Race string `json:"race"`
	// Breed of the pet - Weinerdog, German Sheppard, Bulldog
	Breed string `json:"breed"`
}
