package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/timkellogg/five_three_one/config"
	"github.com/timkellogg/five_three_one/services/exceptions"
)

func handleError(e error, appError exceptions.ApplicationException, w http.ResponseWriter) {
	log.Println(e)
	log.Println(appError.Code)

	ae, err := json.Marshal(appError)
	if err != nil {
		log.Fatal(e)
	}

	w.WriteHeader(appError.HTTPStatus)
	w.Write(ae)
}

// Errors404 - Not found handler
func Errors404(c *config.ApplicationContext, w http.ResponseWriter, r *http.Request) {
	err := errors.New("Resource Not Found")
	handleError(err, exceptions.ResourceNotFoundError, w)
}
