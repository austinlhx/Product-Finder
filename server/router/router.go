package router

import (
	"./routes"
	"github.com/gorilla/mux"
)

//Router sets up the router
func Router() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutesWithMiddlewares(router)

}
