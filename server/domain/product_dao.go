package domain

import (
	"log"
	"strconv"

	"../models"
)

//SearchProduct converts the POST info into a proper format
func SearchProduct(product models.Product) *models.Products {
	//TODO: Verify if info is correct
	tempUpper, _ := strconv.ParseFloat(product.UpperBound, 2)
	tempLower, _ := strconv.ParseFloat(product.LowerBound, 2)
	p := &models.Products{
		ProductName: product.ProductName,
		ProductType: product.ProductType,
		UpperBound:  tempUpper,
		LowerBound:  tempLower,
	}
	log.Println(p)
	return p
}
