package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/timkellogg/five_three_one/app/models"
)

func handleError(e error, appError models.ApplicationError, w http.ResponseWriter) {
	log.Println(e)
	log.Println(appError.Code)

	ae, err := json.Marshal(appError)
	if err != nil {
		log.Fatal(e)
	}

	w.WriteHeader(appError.HTTPStatus)
	w.Write(ae)
}
