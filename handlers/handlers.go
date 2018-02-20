package handlers

import (
	"net/http"
	"time"

	"github.com/timkellogg/five_three_one/config"
	"github.com/timkellogg/five_three_one/services/exceptions"
)

// requireAuthentication - verifies user based upon access token
func requireAuthorization(c *config.ApplicationContext, w http.ResponseWriter, r *http.Request) string {
	token := r.Header.Get("Authorization")

	if token == "" {
		handleError(nil, exceptions.UserNotAuthorized, w)
		return ""
	}

	obfuscatedID, valid := c.Auth.VerifyToken(token)
	if !valid {
		handleError(nil, exceptions.UserNotAuthorized, w)
		return ""
	}

	return obfuscatedID
}

// setCSRFToken - creates a new CSRF token and passes that into the request header
func setCSRFToken(c *config.ApplicationContext, w http.ResponseWriter, r *http.Request) {
	r.Header.Add("X-CSRF-Token", c.Auth.UniqueString())
}

// setAuthorization - adds Authorization header with bearer token
func setAuthorizationCookie(w http.ResponseWriter, token string) {
	authorizationCookie := http.Cookie{
		Name:    "Authorization",
		Value:   token,
		Expires: time.Now().Add(24 * time.Hour),
	}

	http.SetCookie(w, &authorizationCookie)
}
