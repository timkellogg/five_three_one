package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/timkellogg/five_three_one/config"
	"github.com/timkellogg/five_three_one/models"
	"github.com/timkellogg/five_three_one/services/exceptions"
)

// UsersCreate - create an application user
func UsersCreate(c *config.ApplicationContext, w http.ResponseWriter, r *http.Request) {
	var (
		u     models.User
		err   error
		token string
	)

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err = decoder.Decode(&u)
	if err != nil {
		handleError(err, exceptions.JSONParseError, w)
		return
	}

	token, err = u.CreateUser(c)
	if err != nil {
		handleError(err, exceptions.UserCreateError, w)
		return
	}

	expiration := time.Now().Add(24 * time.Hour)
	authorizationCookie := http.Cookie{
		Name:    "Authorization",
		Value:   token,
		Expires: expiration,
	}

	serializedUser, err := u.SerializedUser(c)
	if err != nil {
		handleError(err, exceptions.JSONParseError, w)
		return
	}

	http.SetCookie(w, &authorizationCookie)

	w.WriteHeader(http.StatusCreated)
	w.Write(serializedUser)
}

// UsersShow - return customer details
func UsersShow(c *config.ApplicationContext, w http.ResponseWriter, r *http.Request) {
	var u models.User

	obfuscatedID := requireAuthorization(c, w, r)

	returnedUser, err := u.FindByObfuscatedID(c, obfuscatedID)
	if err != nil {
		handleError(err, exceptions.UserNotAuthorized, w)
	}

	serializedUser, err := returnedUser.SerializedUser(c)
	if err != nil {
		handleError(err, exceptions.JSONParseError, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(serializedUser)
}
