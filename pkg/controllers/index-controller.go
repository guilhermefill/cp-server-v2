package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/guilhermefill/cp-server-v2/pkg/models"
	"github.com/guilhermefill/cp-server-v2/pkg/utils"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog(r)
	users := models.GetUsers()
	if users == nil {
		utils.ErrorLog(r)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	res, err := json.Marshal(users)
	if err != nil {
		utils.ErrorLog(r)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog(r)
	fmt.Fprintf(w, "Health Check")
}
