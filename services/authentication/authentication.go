package authentication

import (
	"fmt"
	"os"
	"time"

	"github.com/satori/go.uuid"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// AuthService - provides authentication
type AuthService struct{}

// CreateToken - signs and encrypts auth token
func (a *AuthService) CreateToken(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = buildAccessTokenClaims(id)

	tokenString, err := token.SignedString([]byte(os.Getenv("AUTH_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken - checks that token is valid and returns user id if it is
func (a *AuthService) VerifyToken(tokenString string) (string, bool) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("AUTH_SECRET")), nil
	})

	if token == nil {
		return "", false
	}

	if _, ok := token.Claims.(jwt.MapClaims)["user_id"]; ok && token.Valid {
		userID := parseClaimsForValue(token, "user_id")
		return userID, true
	}

	return "", false
}

// CreateRefreshToken - creates a valid refresh token
func (a *AuthService) CreateRefreshToken(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = buildRefreshTokenClaims(id)

	tokenString, err := token.SignedString([]byte(os.Getenv("AUTH_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Encrypt string - encrypts string using bcrypt hashing algo
func (a *AuthService) Encrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 7)
	return string(bytes), err
}

// Decrypt - checks if hash matches decrypted string
func (a *AuthService) Decrypt(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// UniqueString - generate guuid string
func (a *AuthService) UniqueString() string {
	return UniqueString()
}

func getKey(t *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("AUTH_SECRET")), nil
}

// buildAccessTokenClaims - creates short-lived "access token" claims for full auth
func buildAccessTokenClaims(id string) jwt.Claims {
	var expireToken time.Duration

	expireToken, err := time.ParseDuration(os.Getenv("AUTH_EXP"))
	if err != nil {
		expireToken = 1
	}

	claims := make(jwt.MapClaims)

	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * expireToken).Unix()
	claims["iat"] = time.Now().Unix()

	return claims
}

// buildRefreshTokenClaims - creates long-lived "refresh token" claims for partial auth
func buildRefreshTokenClaims(id string) jwt.Claims {
	claims := make(jwt.MapClaims)

	claims["jti"] = UniqueString()
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["iat"] = time.Now().Unix()

	return claims
}

func parseClaimsForValue(token *jwt.Token, claimKey string) string {
	return token.Claims.(jwt.MapClaims)[claimKey].(string)
}

// UniqueString - generate a unique uuid
func UniqueString() string {
	return uuid.NewV4().String()
}
