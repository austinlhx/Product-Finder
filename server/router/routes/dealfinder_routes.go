package routes

import (
	"net/http"

	"../../controllers"
)

var dealFinderRoutes = []Route{
	Route{
		URI:     "/api",
		Method:  http.MethodPost,
		Handler: controllers.SearchProduct,
	},
	Route{
		URI:     "/api",
		Method:  http.MethodGet,
		Handler: controllers.GetProduct,
	},
}
