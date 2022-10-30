package main

import (
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

	router.HandleFunc("/post", postController.CreatePost).Methods("POST")
	router.HandleFunc("/post", postController.GetPost).Methods("GET")

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	server.ListenAndServe()
}
