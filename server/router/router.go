package router

import (
	"./routes"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutesWithMiddlewares(router)

}
