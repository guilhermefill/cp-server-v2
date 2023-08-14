package routes

import (
	"github.com/gorilla/mux"
	"github.com/guilhermefill/cp-server-v2/pkg/controllers"
)

var AuthRoutes = func(router *mux.Router) {
	router.HandleFunc("/auth", controllers.AuthHandler).Methods(("GET"))
}
