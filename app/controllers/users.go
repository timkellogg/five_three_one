package controllers

import (
	"net/http"

	"github.com/timkellogg/five_three_one/app/models"
)

// UsersCreate - create an application user
func UsersCreate(w http.ResponseWriter, r *http.Request) {
	handleError(nil, models.NotImplementedError, w)
}
