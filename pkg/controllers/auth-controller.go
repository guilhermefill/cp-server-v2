package controllers

import (
	"fmt"
	"net/http"

	"github.com/guilhermefill/cp-server-v2/pkg/utils"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog(r)
	fmt.Fprintf(w, "this is the auth route")
}
