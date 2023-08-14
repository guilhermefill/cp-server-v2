package routes

import (
	"github.com/gorilla/mux"
	"github.com/guilhermefill/cp-server-v2/pkg/controllers"
)

var PostRoutes = func(router *mux.Router) {
	router.HandleFunc("/posts", controllers.GetAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", controllers.GetPostByID).Methods("GET")
	router.HandleFunc("/posts/carousel", controllers.GetCarouselPosts).Methods("POST")
}
