package models

// User struct that represents the data stored in the db
type User struct {
	// Firstname of the User
	Firstname string `json:"firstname"`
	// Lastname of the User
	Lastname string `json:"lastname"`
	// Address of the User
	Address *Address `json:"address"`
}
