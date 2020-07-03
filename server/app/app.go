package app

import (
	"github.com/austinlhx/DealFinder/server/router"
	"log"
	"net/http"
)

func StartApp(){
	r := router.Router()
	log.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
