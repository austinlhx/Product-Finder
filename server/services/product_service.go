package services

import (
	"sync"

	"../domain"
	"../models"
)

//SearchBestBuy calls the domain SearchBestBuy
func SearchBestBuy(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	domain.SearchBestBuy(product, ch, wg)
}
//SearchAmazon calls the domain SearchAmazon
func SearchAmazon(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	domain.SearchAmazon(product, ch, wg)
}
//SearchNewEgg calls the domain SearchNewEgg
func SearchNewEgg(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	domain.SearchNewEgg(product, ch, wg)
}
//SearchBHPhotoVideo calls the domain SearchBHPhotoVideo
func SearchBHPhotoVideo(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	domain.SearchBHPhotoVideo(product, ch, wg)
}
//SearchAdorama calls the domain SearchAdorama
func SearchAdorama(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	domain.SearchAdorama(product, ch, wg)
}
//SearchBloomingdales calls the domain SearchBloomingdales
func SearchBloomingdales(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	domain.SearchBloomingdales(product, ch, wg)
}
//SearchSaksFifth calls the domain SearchSaksFifth
func SearchSaksFifth(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	domain.SearchSaksFifth(product, ch, wg)
}
//SearchNeimanMarcus calls the domain SearchNeimanMarcus
func SearchNeimanMarcus(product *models.Products, ch chan models.ProductFound, wg *sync.WaitGroup) {
	domain.SearchNeimanMarcus(product, ch, wg)
}
//SearchProduct calls the domain SearchProduct
func SearchProduct(product models.Product) *models.Products {
	return domain.SearchProduct(product)
}
