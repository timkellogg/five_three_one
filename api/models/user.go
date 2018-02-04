package models

import (
	"database/sql"
	"time"

	"github.com/timkellogg/five_three_one/services/authentication"
)

// User - a consumer of the application
type User struct {
	ID                int64
	ObfuscatedID      string        `json:"obfuscated_id"`
	Email             string        `json:"email"`
	EncryptedPassword string        `json:"encrypted_password"`
	CreatedAt         time.Duration `json:"created_at"`
}

// CreateUser - saves user to db
func (u *User) CreateUser(db *sql.DB, auth *authentication.AuthService, password string) error {
	var err error

	u.ObfuscatedID = createObfuscatedID()

	u.EncryptedPassword, err = auth.Encrypt(password)
	if err != nil {
		return err
	}

	err = db.QueryRow("INSERT INTO users(email, obfuscated_id, encrypted_password) VALUES($1, $2, $3) RETURNING id",
		u.Email, u.ObfuscatedID, u.EncryptedPassword).Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}
