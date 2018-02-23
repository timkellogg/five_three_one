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

// TokenRequest - parameters from request call
type TokenRequest struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
}

// Authorize - grants access
func Authorize(c *config.ApplicationContext, w http.ResponseWriter, r *http.Request) {
	var u models.User
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	if err != nil {
		handleError(err, exceptions.JSONParseError, w)
		return
	}

	user, err := u.FindByEmail(c)
	if err != nil {
		handleError(err, exceptions.ResourceNotFoundError, w)
		return
	}

	valid := c.Auth.Decrypt(u.Password, user.EncryptedPassword)

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
	ut.UserTokenInvalidate(c)

	_, err = ut.UserTokenCreate(c)
	if err != nil {
		handleError(err, exceptions.RefreshTokenCreateError, w)
		return
	}

	accessToken, err := c.Auth.CreateToken(u.ObfuscatedID)
	if err != nil {
		handleError(err, exceptions.TokenCreateError, w)
		return
	}

	setCSRFToken(c, w, r)
	setAuthorizationCookie(c, w, accessToken)

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
	defer r.Body.Close()

	var (
		u   models.User
		us  models.UserSecret
		req TokenRequest
	)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if req.GrantType != "refresh_token" {
		handleError(err, exceptions.UnknownTokenGrantType, w)
		return
	}

	// find user from auth header
	user, err := userFromAuthorizationHeader(c, w, r)
	if err != nil {
		handleError(err, exceptions.UserNotAuthorized, w)
		return
	}

	us.UserID = user.ID

	// find user secret
	userSecret, err := us.UserSecretFindByID(c)
	if err != nil {
		handleError(err, exceptions.ResourceNotFoundError, w)
		return
	}

	// validate credentials
	if userSecret.ClientID != req.ClientID || userSecret.ClientSecret != req.ClientSecret {
		handleError(err, exceptions.UserNotAuthorized, w)
		return
	}

	// find active refresh token
	refreshToken, err := user.UserActiveToken(c)
	if err != nil {
		handleError(err, exceptions.UserNotAuthorized, w)
		return
	}

	// create new access token
	accessToken, err := c.Auth.CreateToken(u.ObfuscatedID)
	if err != nil {
		handleError(err, exceptions.TokenCreateError, w)
		return
	}

	setCSRFToken(c, w, r)
	setAuthorizationCookie(c, w, accessToken)

	AuthorizeReponse := AuthorizeReponse{
		TokenType:    "bearer",
		AccessToken:  accessToken,
		ExpiresIn:    time.Now().Add(time.Hour * 72),
		RefreshToken: refreshToken.Token,
	}

	response, err := json.Marshal(AuthorizeReponse)
	if err != nil {
		handleError(err, exceptions.JSONParseError, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Confirm -
func Confirm(c *config.ApplicationContext, w http.ResponseWriter, r *http.Request) {

}
