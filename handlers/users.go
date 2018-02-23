package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/timkellogg/five_three_one/config"
	"github.com/timkellogg/five_three_one/models"
	"github.com/timkellogg/five_three_one/services/exceptions"
)

// UsersResponse - json returned from users handlers
type UsersResponse struct {
	ObfuscatedID string `json:"id"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Email        string `json:"email"`
	Active       bool   `json:"active"`
}

// UsersCreate - create an application user
func UsersCreate(c *config.ApplicationContext, w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var u models.User
	var us models.UserSecret

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&u)
	if err != nil {
		handleError(err, exceptions.JSONParseError, w)
		return
	}

	// create user
	user, err := u.CreateUser(c)
	if err != nil {
		handleError(err, exceptions.UserCreateError, w)
		return
	}

	us.UserID = user.ID

	// create user secret
	userSecret, err := us.CreateUserSecret(c)
	if err != nil {
		handleError(err, exceptions.UserCreateError, w)
		return
	}

	// create new access token
	accessToken, err := c.Auth.CreateToken(u.ObfuscatedID)
	if err != nil {
		handleError(err, exceptions.TokenCreateError, w)
		return
	}

	// set auth
	setAuthorizationCookie(c, w, accessToken)

	responseStructure := UsersResponse{
		Active:       u.Active,
		Email:        u.Email,
		ObfuscatedID: u.ObfuscatedID,
		ClientID:     userSecret.ClientID,
		ClientSecret: userSecret.ClientSecret,
	}

	response, err := json.Marshal(responseStructure)
	if err != nil {
		handleError(err, exceptions.JSONParseError, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// UsersShow - return customer details
func UsersShow(c *config.ApplicationContext, w http.ResponseWriter, r *http.Request) {
	var us models.UserSecret
	var err error

	user, err := userFromAuthorizationHeader(c, w, r)
	if err != nil {
		handleError(err, exceptions.UserNotAuthorized, w)
		return
	}

	us = models.UserSecret{UserID: user.ID}
	userSecret, err := us.UserSecretFindByID(c)
	if err != nil {
		handleError(err, exceptions.ResourceNotFoundError, w)
		return
	}

	responseStructure := UsersResponse{
		Active:       user.Active,
		Email:        user.Email,
		ObfuscatedID: user.ObfuscatedID,
		ClientID:     userSecret.ClientID,
		ClientSecret: userSecret.ClientSecret,
	}

	response, err := json.Marshal(responseStructure)
	if err != nil {
		handleError(err, exceptions.JSONParseError, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
