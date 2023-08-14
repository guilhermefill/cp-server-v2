package controllers

import (
	"fmt"
	"net/http"

	"github.com/guilhermefill/cp-server-v2/pkg/utils"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog(r)
	fmt.Fprintf(w, "Hello World")
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog(r)
	fmt.Fprintf(w, "Health Check")
}
