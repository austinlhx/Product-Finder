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

func SearchNewEgg(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	domain.SearchNewEgg(product, ch, wg)
}

func SearchBHPhotoVideo(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	domain.SearchBHPhotoVideo(product, ch, wg)
}

func SearchAdorama(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	domain.SearchAdorama(product, ch, wg)
}

func SearchBloomingdales(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	domain.SearchBloomingdales(product, ch, wg)
}

func SearchSaksFifth(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	domain.SearchSaksFifth(product, ch, wg)
}

func SearchNeimanMarcus(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	domain.SearchNeimanMarcus(product, ch, wg)
}

func SearchProduct(product models.Product) *models.Products {
	return domain.SearchProduct(product)
}
