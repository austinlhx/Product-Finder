package services

import (
	"../models"
	"../domain"
)

func SearchBestBuy(product *models.Products) []models.ProductFound{
	return domain.SearchBestBuy(product)
}
func SearchAmazon(product *models.Products) []models.ProductFound{
	return domain.SearchAmazon(product)
}

func SearchProduct(product models.Product) *models.Products {
	return domain.SearchProduct(product)
}
