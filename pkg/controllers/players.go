package controllers

import (
	"fmt"
	"net/http"
)

func PlayerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the players page!")
}