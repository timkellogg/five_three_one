package handlers

import (
	"net/http"

	"github.com/timkellogg/five_three_one/config"
	"github.com/timkellogg/five_three_one/services/exceptions"
)

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
