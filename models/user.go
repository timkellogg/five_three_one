package models

import (
	"time"

	"github.com/timkellogg/five_three_one/config"
)

// User - a consumer of the application
type User struct {
	ID                int64
	ObfuscatedID      string        `json:"obfuscated_id" db:"obfuscated_id"`
	Email             string        `json:"email" db:"email"`
	Password          string        `json:"password" db:""`
	EncryptedPassword string        `json:"encrypted_password" db:"encrypted_password"`
	CreatedAt         time.Duration `json:"created_at" db:"created_at"`
}

// CreateUser - saves user to db
func (u *User) CreateUser(c *config.ApplicationContext) (string, error) {
	var err error
	var token string

	u.ObfuscatedID = createObfuscatedID()

	u.EncryptedPassword, err = c.Auth.Encrypt(u.Password)
	if err != nil {
		return "", err
	}

	err = c.Database.QueryRow("INSERT INTO users (email, obfuscated_id, encrypted_password) VALUES($1, $2, $3) RETURNING id",
		u.Email, u.ObfuscatedID, u.EncryptedPassword).Scan(&u.ID)
	if err != nil {
		return "", err
	}

	token, err = c.Auth.CreateToken(u.Email, u.ObfuscatedID)
	if err != nil {
		return "", err
	}

	return token, nil
}
