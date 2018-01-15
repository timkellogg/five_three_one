package controllers

import (
	"errors"
	"net/http"

	"github.com/timkellogg/five_three_one/app/models"
)

// Errors404 - Not found handler
func Errors404(w http.ResponseWriter, r *http.Request) {
	err := errors.New("Resource Not Found")
	handleError(err, models.ResourceNotFoundError, w)
}
