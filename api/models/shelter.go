package models

type Shelter struct {
	Name    string   `name`
	Contact string   `contact`
	Address *Address `address`
}
