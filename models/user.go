package models

import (
	"time"

	"github.com/lib/pq"

	"github.com/timkellogg/five_three_one/config"
)

// User - a consumer of the application
type User struct {
	ID                int64
	ObfuscatedID      string      `json:"obfuscated_id" db:"obfuscated_id"`
	Email             string      `json:"email" db:"email"`
	Password          string      `json:"password" db:"-"`
	EncryptedPassword string      `json:"-" db:"encrypted_password"`
	Active            bool        `json:"active" db:"active"`
	CreatedAt         pq.NullTime `json:"created_at" db:"created_at"`
	UpdatedAt         pq.NullTime `json:"updated_at" db:"updated_at"`
}

// CreateUser - saves a user in the db
func (u *User) CreateUser(c *config.ApplicationContext) (*User, error) {
	encryptedPassword, err := c.Auth.Encrypt(u.Password)
	if err != nil {
		return u, err
	}

	err = c.Database.
		QueryRow("INSERT INTO users (email, obfuscated_id, encrypted_password, created_at) VALUES($1,$2,$3,$4) RETURNING id, email, active, obfuscated_id, encrypted_password",
			u.Email, createObfuscatedID(), encryptedPassword, pq.NullTime{Valid: true, Time: time.Now()}).
		Scan(&u.ID, &u.Email, &u.Active, &u.ObfuscatedID, &u.EncryptedPassword)
	if err != nil {
		return u, err
	}

	return u, nil
}

// FindByObfuscatedID - returns user by obfuscatedID
func (u *User) FindByObfuscatedID(c *config.ApplicationContext) (*User, error) {
	err := c.Database.
		QueryRow("SELECT id, email, active, obfuscated_id, encrypted_password FROM users WHERE obfuscated_id = $1", u.ObfuscatedID).
		Scan(&u.ID, &u.Email, &u.Active, &u.ObfuscatedID, &u.EncryptedPassword)
	if err != nil {
		return u, err
	}

	return u, nil
}

// FindByEmail - look up by email
func (u *User) FindByEmail(c *config.ApplicationContext) (*User, error) {
	err := c.Database.
		QueryRow("SELECT id, email, active, obfuscated_id, encrypted_password FROM users WHERE email = $1", u.Email).
		Scan(&u.ID, &u.Email, &u.Active, &u.ObfuscatedID, &u.EncryptedPassword)
	if err != nil {
		return u, err
	}

	return u, nil
}

// UserActiveToken - find active refresh token for user
func (u *User) UserActiveToken(c *config.ApplicationContext) (*UserToken, error) {
	var userToken UserToken

	err := c.Database.
		QueryRow("SELECT id, token, user_id, active, created_at, updated_at FROM user_tokens WHERE user_id = $1 AND active = true LIMIT 1", u.ID).
		Scan(&userToken.ID, &userToken.Token, &userToken.UserID, &userToken.Active, &userToken.CreatedAt, &userToken.UpdatedAt)
	if err != nil {
		return &userToken, err
	}

	return &userToken, nil
}
