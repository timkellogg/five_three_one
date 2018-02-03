package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/timkellogg/five_three_one/app/models"
)

// UsersCreate - create an application user
func UsersCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var u models.User
	err := decoder.Decode(&u)
	if err != nil {
		handleError(err, models.JSONParseError, w)
	}

	// token, err := authentication.CreateToken(u.Email, u.Password)
	// if err != nil {
	// 	handleError(err, models.TokenCreateError, w)
	// }

	// handleError(nil, models.NotImplementedError, w)
	w.WriteHeader(http.StatusUnprocessableEntity)
}
