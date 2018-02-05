package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/timkellogg/five_three_one/models"
)

// UsersCreate - create an application user
func UsersCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var u models.User
	err := decoder.Decode(&u)
	if err != nil {
		handleError(err, JSONParseError, w)
	}

	if err != nil {
		handleError(err, TokenCreateError, w)
	}

	handleError(nil, NotImplementedError, w)
	w.WriteHeader(http.StatusUnprocessableEntity)
}
