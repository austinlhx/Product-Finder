package models

//Product sets the product being received from a post request
type Product struct {
	ProductName string `json:"ProductName, omitempty"`
	ProductType string `json:"ProductType, omitempty"`
	UpperBound  string `json:"UpperBound, omitempty"`
	LowerBound  string `json:"LowerBound, omitempty"`
}
//Products sets the product being received from a post request
type Products struct {
	ProductName string
	ProductType string
	UpperBound  float64
	LowerBound  float64
}

//ProductFound sets the product being received from a post request
type ProductFound struct {
	Name string 
	Price float64
	Link string
	Image string
}	