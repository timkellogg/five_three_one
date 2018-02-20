package handlers

import (
	"encoding/json"
	"net/http"
	"time"

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
	var (
		u     models.User
		us    models.UserSecret
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

	// Needs to return user
	token, err = u.CreateUser(c)
	if err != nil {
		handleError(err, exceptions.UserCreateError, w)
		return
	}

	user, err := u.FindByEmail(c)
	if err != nil {
		handleError(err, exceptions.UserCreateError, w)
		return
	}

	us.UserID = user.ID

	userSecret, err := us.CreateUserSecret(c)
	if err != nil {
		handleError(err, exceptions.UserCreateError, w)
		return
	}

	// token, err = c.Auth.CreateToken(u.Email, u.ObfuscatedID)
	// if err != nil {
	// 	return u, err
	// }

	expiration := time.Now().Add(24 * time.Hour)
	authorizationCookie := http.Cookie{
		Name:    "Authorization",
		Value:   "Bearer " + token,
		Expires: expiration,
	}
	http.SetCookie(w, &authorizationCookie)

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
	var u models.User
	var us models.UserSecret

	obfuscatedID := requireAuthorization(c, w, r)

	returnedUser, err := u.FindByObfuscatedID(c, obfuscatedID)
	if err != nil {
		handleError(err, exceptions.UserNotAuthorized, w)
		return
	}

	us.UserID = returnedUser.ID
	userSecret, err := us.UserSecretFindByID(c)
	if err != nil {
		handleError(err, exceptions.ResourceNotFoundError, w)
		return
	}

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

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
