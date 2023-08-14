package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/guilhermefill/cp-server-v2/pkg/routes"
)

func main() {

	r := mux.NewRouter()
	routes.IndexHandler(r)
	routes.AuthRoutes(r)
	routes.PostRoutes(r)
	http.Handle("/", r)

	srv := &http.Server{
		Addr:         ":5005",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	srv.ListenAndServe()
}
