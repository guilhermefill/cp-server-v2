package controllers

import (
	"fmt"
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is the auth route")
}
