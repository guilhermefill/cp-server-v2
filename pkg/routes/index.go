package routes

import (
	"github.com/gorilla/mux"
	"github.com/guilhermefill/cp-server-v2/pkg/controllers"
)

var IndexHandler = func(router *mux.Router) {
	router.HandleFunc("/", controllers.IndexHandler).Methods("GET")
	router.HandleFunc("/healthCheck", controllers.HealthCheckHandler).Methods("GET")
}
