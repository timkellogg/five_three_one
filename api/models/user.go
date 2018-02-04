package models

import (
	"time"
)

// User - a consumer of the application
type User struct {
	ID                int64
	ObfuscatedID      string `json:"obfuscated_id"`
	Email             string `json:"email"`
	EncryptedPassword string `json:"encrypted_password"`
	CreatedAt time.Duration `json:"created_at"`
}

// CreateUser - saves user to db
func (u *User) CreateUser() error {
	obfuscatedID := createObfuscatedID()

	// create token

	err := Database.QueryRow("INSERT INTO users(email) VALUES($1) RETURNING id", u.Email).Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}

CREATE TABLE users (
    id                  serial PRIMARY KEY,
    obfuscated_id       VARCHAR(128) NOT NULL UNIQUE,
    email               VARCHAR(128) NOT NULL UNIQUE,
    encrypted_password  VARCHAR(128) NOT NULL,
    created_at          TIMESTAMP,
    updated_at          TIMESTAMP,
    active              BOOLEAN DEFAULT TRUE
);
