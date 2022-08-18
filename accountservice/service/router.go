package service

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	// Create an instance of Gorilla router
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			HandlerFunc(route.HandlerFunc)
	}
	return router
}
