package models

type Product struct {
	ProductName     string `json:"product_name, omitempty"` 
	ProductType     string `json:"product_type, omitempty"` 
	UpperBoundPrice float64 `json:"upper_bound, omitempty"` 
	LowerBoundPrice float64 `json:"lower_bound, omitempty"` 
} 
