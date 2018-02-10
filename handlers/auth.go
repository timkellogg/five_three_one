package handlers

import (
	"net/http"

	"github.com/timkellogg/five_three_one/config"
	"github.com/timkellogg/five_three_one/models"
	"github.com/timkellogg/five_three_one/services/exceptions"
)

// Login - refreshes tokens
func Login(c *config.ApplicationContext, w http.ResponseWriter, r *http.Request) {
	var u models.User

	obfuscatedID := requireAuthorization(c, w, r)

	// Find user
	returnedUser, err := u.FindByObfuscatedID(c, obfuscatedID)
	if err != nil {
		handleError(err, exceptions.ResourceNotFoundError, w)
	}

	// Add CSRF to header
	addCSRFToken(c, w, r)

	// Create refresh token
	refreshToken, err := c.Auth.CreateRefreshToken(obfuscatedID)
	if err != nil {
		handleError(err, exceptions.RefreshTokenCreateError, w)
		return
	}

	// Store refresh token in database

	// Create access token
	accessToken, err := c.Auth.CreateToken(returnedUser.Email, obfuscatedID)
	if err != nil {
		handleError(err, exceptions.TokenCreateError, w)
		return
	}

	// authorizations table stores JTI (unique string to identify this specific token)

	w.WriteHeader(http.StatusOK)
	// w.Write(i)
}

// expiration := time.Now().Add(24 * time.Hour)
// authorizationCookie := http.Cookie{
// 	Name:    "Authorization",
// 	Value:   token,
// 	Expires: expiration,
// }

// serializedUser, err := u.SerializedUser(c)
// if err != nil {
// 	handleError(err, exceptions.JSONParseError, w)
// 	return
// }

// http.SetCookie(w, &authorizationCookie)
