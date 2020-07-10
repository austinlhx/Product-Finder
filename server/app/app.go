package app

import (
	"../router"
	"log"
	"net/http"
)

//StartApp starts the app
func StartApp() {
	r := router.Router()
	log.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
