package routes

import (
	"github.com/austinlhx/DealFinder/server/controllers"
	"net/http"
)

var dealFinderRoutes = []Route{
	Route{
		URI:     "/api",
		Method:  http.MethodPost,
		Handler: controllers.SearchProduct,
	},
}