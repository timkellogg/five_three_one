package controllers

import "net/http"

// InfoShow - prints out details about api
func InfoShow(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}
