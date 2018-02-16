package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/timkellogg/five_three_one/config"
	"github.com/timkellogg/five_three_one/models"
	"github.com/timkellogg/five_three_one/services/exceptions"
)

// AuthorizeReponse - structure of login response
type AuthorizeReponse struct {
	TokenType    string    `json:"token_type"`
	AccessToken  string    `json:"access_token"`
	ExpiresIn    time.Time `json:"expires_in"`
	RefreshToken string    `json:"refresh_token"`
}

// TokenResponse - structure of token response
type TokenResponse struct{}

// Authorize - grants access
func Authorize(c *config.ApplicationContext, w http.ResponseWriter, r *http.Request) {
	var (
		user  *models.User
		u     models.User
		err   error
		valid bool
	)

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err = decoder.Decode(&u)
	if err != nil {
		handleError(err, exceptions.JSONParseError, w)
		return
	}

	user, err = u.FindByEmail(c)
	if err != nil {
		handleError(err, exceptions.ResourceNotFoundError, w)
		return
	}

	valid = c.Auth.Decrypt(user.Password, user.EncryptedPassword)

	if !valid {
		handleError(err, exceptions.ResourceNotFoundError, w)
		return
	}

	refreshToken, err := c.Auth.CreateRefreshToken(user.ObfuscatedID)
	if err != nil {
		handleError(err, exceptions.RefreshTokenCreateError, w)
		return
	}

	ut := models.UserToken{Token: refreshToken, UserID: user.ID}

	// Look for previous access tokens and invalidate
	ut.Invalidate(c)

	_, err = ut.Save(c)
	if err != nil {
		handleError(err, exceptions.RefreshTokenCreateError, w)
		return
	}

	accessToken, err := c.Auth.CreateToken(u.Email, u.ObfuscatedID)
	if err != nil {
		handleError(err, exceptions.TokenCreateError, w)
		return
	}

	setCSRFToken(c, w, r)
	setAuthorizationCookie(w, accessToken)

	AuthorizeReponse := AuthorizeReponse{
		TokenType:    "bearer",
		AccessToken:  accessToken,
		ExpiresIn:    time.Now().Add(time.Hour * 72),
		RefreshToken: refreshToken,
	}

	response, err := json.Marshal(AuthorizeReponse)
	if err != nil {
		handleError(err, exceptions.JSONParseError, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Token - get new access token with refresh token
func Token(c *config.ApplicationContext, w http.ResponseWriter, r *http.Request) {

}

// Confirm -
func Confirm(c *config.ApplicationContext, w http.ResponseWriter, r *http.Request) {

}
