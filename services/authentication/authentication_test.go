package authentication

import (
	"os"
	"testing"
)

func TestAuthService(t *testing.T) {
	os.Setenv("AUTH_SECRET", "secret")

	auth := &AuthService{}

	var (
		valid bool
		err   error
	)

	email, id := "test@test.com", "1"

	token, err := auth.CreateToken(email, id)
	if err != nil {
		t.Errorf("AuthService failed to create a token: %v", err)
	}

	valid = auth.VerifyToken(token)
	if !valid {
		t.Errorf("AuthService failed to verify token: %v", err)
	}

	valid = auth.VerifyToken("invalid")
	if valid {
		t.Error("AuthService verified a bad token")
	}
}

func TestEncryption(t *testing.T) {
	var valid bool
	password := "password"

	auth := &AuthService{}

	encrypted, err := auth.Encrypt(password)
	if err != nil || encrypted == "" {
		t.Error("AuthService failed to encrypt string")
	}

	valid = auth.Decrypt(password, encrypted)
	if !valid {
		t.Errorf("AuthService failed to decrypt password")
	}

	valid = auth.Decrypt(password, "invalid")
	if valid {
		t.Errorf("AuthService verified a bad token")
	}
}
