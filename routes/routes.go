package routes

import (
	"github.com/enylvia/boilerplate-go/domain/controller"
	"github.com/gorilla/mux"
)

func RouterPost(r *mux.Router, post controller.Controller) {
	r.HandleFunc("/post", post.CreatePost).Methods("POST")
	r.HandleFunc("/post", post.GetPost).Methods("GET")
}
