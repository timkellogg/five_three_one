package authentication

import (
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// AuthService - provides authentication
type AuthService struct{}

// CreateToken - signs and encrypts auth token
func (a *AuthService) CreateToken(email, password string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// claims
	claims := make(jwt.MapClaims)
	claims["email"] = email
	claims["password"] = password
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iat"] = time.Now().Unix()

	token.Claims = claims

	// should use different config instead of hardcoding to use DevConfig
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Println(err)
		return "", err
	}

	return tokenString, nil
}
