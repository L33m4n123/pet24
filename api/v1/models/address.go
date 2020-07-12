package models

// Address struct to represent data stored in the db
type Address struct {
	// City of the address
	City string `json:"city"`
	// Postalcode of the Address
	PostalCode string `json:"postalCode"`
	// Street of the Address
	Street string `json:"street"`
	//SuplementOne of the Address
	SuplementOne string `json:"suplementOne"`
}
