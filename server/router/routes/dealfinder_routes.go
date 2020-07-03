package routes

import (
	"github.com/austinlhx/DealFinder/server/controllers"
	"net/http"
)

var dealFinderRoutes = []Route{
	Route{
		URI:     "",
		Method:  http.MethodPost,
		Handler: controllers.GetEachCase,
	},
}