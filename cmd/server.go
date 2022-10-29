package cmd

import (
	"boilerplate-go/routes"

	"github.com/gorilla/mux"
)

func InitializeRestAPI() *mux.Router {
	// Initialize mux router
	router := mux.NewRouter()
	routes.NewAPI().Routes(router)
	return router
}
