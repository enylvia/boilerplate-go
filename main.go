package main

import (
	"github.com/enylvia/boilerplate-go/routes"
	"net/http"

	"github.com/enylvia/boilerplate-go/database"
	"github.com/enylvia/boilerplate-go/domain/controller"
	"github.com/enylvia/boilerplate-go/domain/repository"
	"github.com/enylvia/boilerplate-go/domain/service"
	"github.com/gorilla/mux"
)

func main() {
	db := database.Connect()

	postRepository := repository.NewRepository(db)
	postService := service.NewService(postRepository)
	postController := controller.NewController(postService)

	router := mux.NewRouter()
	routes.RouterPost(router, *postController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	server.ListenAndServe()
}
