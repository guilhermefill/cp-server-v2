package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is the posts route")
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "this is the post by id route"+id)
}

func GetCarouselPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is the carousel posts route")
}
