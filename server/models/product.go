package models

type Product struct {
	ProductName string `json:"ProductName, omitempty"`
	ProductType string `json:"ProductType, omitempty"`
	UpperBound  string `json:"UpperBound, omitempty"`
	LowerBound  string `json:"LowerBound, omitempty"`
}

type Products struct {
	ProductName string
	ProductType string
	UpperBound  float64
	LowerBound  float64
}

type ProductFound struct {
	Name string 
	Price float64
	//Link string
	//Photo string
}