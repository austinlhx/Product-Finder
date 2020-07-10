package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"../models"
	"../services"
	"../utils"
)

var allProducts = []models.ProductFound{}

//SearchProduct is a controller that holds all of the logic
func SearchProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "Error decoding product",
			StatusCode: http.StatusInternalServerError,
			Code:       "server_error",
		}
		jsonValue, err := json.Marshal(apiErr)
		if err != nil {
			w.WriteHeader(apiErr.StatusCode)
		}
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		log.Println("Error Decoding Product")
		log.Println(err)
		return
	}
	productQuery := services.SearchProduct(product)
	//log.Println(product.ProductName)
	//var wg sync.WaitGroup
	allProducts = []models.ProductFound{}
	productsFound := make(chan models.ProductFound)
	if productQuery.ProductType == "Electronics"{
		go func() {
			var wg sync.WaitGroup
			wg.Add(5)
			go services.SearchBestBuy(productQuery, productsFound, &wg) //Fan out 2 go routines
			go services.SearchAmazon(productQuery, productsFound, &wg)  //producer
			go services.SearchNewEgg(productQuery, productsFound, &wg)
			go services.SearchBHPhotoVideo(productQuery, productsFound, &wg)
			go services.SearchAdorama(productQuery, productsFound, &wg)
			wg.Wait()
			close(productsFound)
		}()
		 //reset products
		//consumer down here
		for productFound := range productsFound { //fanin
			allProducts = append(allProducts, productFound)
		}
	}

	if productQuery.ProductType == "Clothing"{
		go func() {
			var wg sync.WaitGroup
			wg.Add(3)
			go services.SearchBloomingdales(productQuery, productsFound, &wg)
			go services.SearchSaksFifth(productQuery, productsFound, &wg)
			go services.SearchNeimanMarcus(productQuery, productsFound, &wg)
			wg.Wait()
			close(productsFound)
		}()
		//consumer down here
		for productFound := range productsFound { //fanin
			allProducts = append(allProducts, productFound)
		}
	}
	

}
//GetProduct is a controller that GET the product info
func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(allProducts)
}
