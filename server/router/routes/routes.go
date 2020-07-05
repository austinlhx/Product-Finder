package routes

import (
	"net/http"

	"../../middleware"
	"github.com/gorilla/mux"
)

type Route struct {
	URI     string
	Method  string
	Handler func(w http.ResponseWriter, r *http.Request)
}

func Load() []Route {
	routes := dealFinderRoutes
	return routes
}

func SetupRoutesWithMiddlewares(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.URI,
			middleware.SetMiddlewareLogger(
				middleware.SetMiddlewareJSON(route.Handler),
			),
		).Methods(route.Method)
	}
	return r
}
