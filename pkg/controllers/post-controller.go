package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guilhermefill/cp-server-v2/pkg/models"
	"github.com/guilhermefill/cp-server-v2/pkg/utils"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog(r)
	posts := models.GetAllPosts()
	res, err := json.Marshal(posts)
	if err != nil {
		utils.ErrorLog(r)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog(r)

	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "this is the post by id route "+id)
}

func GetCarouselPosts(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog(r)
	fmt.Fprintf(w, "this is the carousel posts route")
}
