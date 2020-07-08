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
	productsFound := make(chan models.ProductFound)

    go func() {
        var wg sync.WaitGroup
        wg.Add(2)
        go services.SearchBestBuy(productQuery, productsFound, &wg) //Fan out 2 go routines
        go services.SearchAmazon(productQuery, productsFound, &wg) //producer
        wg.Wait()
        close(productsFound)
    }()
//consumer down here
    for productFound := range productsFound { //fanin
        allProducts = append(allProducts, productFound)
	}
	
	log.Println(allProducts)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(allProducts)
}
