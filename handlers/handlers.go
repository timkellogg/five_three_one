package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/timkellogg/five_three_one/config"
	"github.com/timkellogg/five_three_one/models"
	"github.com/timkellogg/five_three_one/services/exceptions"
)

// userFromAuthorizationHeader - verifies user based upon access token
func userFromAuthorizationHeader(c *config.ApplicationContext, w http.ResponseWriter, r *http.Request) (*models.User, error) {
	var u models.User

	token := r.Header.Get("Authorization")

	if token == "" {
		return &u, errors.New("Authorization Token was empty")
	}

	// verify token
	obfuscatedID, valid := c.Auth.VerifyToken(token)
	if !valid {
		return &u, errors.New("Authorization Token was unable to be verified")
	}

	// find user
	u.ObfuscatedID = obfuscatedID
	user, err := u.FindByObfuscatedID(c)
	if err != nil {
		return &u, errors.New("Authorization Token was unable to be verified")
	}

	return user, nil
}

// setCSRFToken - creates a new CSRF token and passes that into the request header
func setCSRFToken(c *config.ApplicationContext, w http.ResponseWriter, r *http.Request) {
	r.Header.Add("X-CSRF-Token", c.Auth.UniqueString())
}

// setAuthorization - adds Authorization header with bearer token
func setAuthorizationCookie(c *config.ApplicationContext, w http.ResponseWriter, u *models.User) {
	token, err := c.Auth.CreateToken(u.ObfuscatedID)
	if err != nil {
		handleError(err, exceptions.UserCreateError, w)
		return
	}

	authorizationCookie := http.Cookie{
		Name:    "Authorization",
		Value:   token,
		Expires: time.Now().Add(24 * time.Hour),
	}

	http.SetCookie(w, &authorizationCookie)
}
