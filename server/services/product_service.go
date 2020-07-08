package services

import (
	"sync"

	"../domain"
	"../models"
)

func SearchBestBuy(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	domain.SearchBestBuy(product, ch, wg)
}
func SearchAmazon(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	domain.SearchAmazon(product, ch, wg)
}

func SearchProduct(product models.Product) *models.Products {
	return domain.SearchProduct(product)
}
