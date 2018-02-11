package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/timkellogg/five_three_one/config"
)

// User - a consumer of the application
type User struct {
	ID                int64
	ObfuscatedID      string    `json:"obfuscated_id" db:"obfuscated_id"`
	Email             string    `json:"email" db:"email"`
	Password          string    `json:"password" db:"-"`
	EncryptedPassword string    `json:"-" db:"encrypted_password"`
	Active            bool      `json:"active" db:"active"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}

// CreateUser - saves user to db
func (u *User) CreateUser(c *config.ApplicationContext) (string, error) {
	var err error
	var token string

	u.ObfuscatedID = createObfuscatedID()
	u.EncryptedPassword, err = c.Auth.Encrypt(u.Password)
	u.Active = true

	if err != nil {
		return "", err
	}

	err = c.Database.
		QueryRow("INSERT INTO users (email, obfuscated_id, encrypted_password) VALUES($1,$2,$3) RETURNING email, active", u.Email, u.ObfuscatedID, u.EncryptedPassword).
		Scan(&u.Email, &u.Active)
	if err != nil {
		return "", err
	}

	token, err = c.Auth.CreateToken(u.Email, u.ObfuscatedID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// FindByObfuscatedID - returns user by obfuscatedID
func (u *User) FindByObfuscatedID(c *config.ApplicationContext, obfuscatedID string) (User, error) {
	var (
		email        string
		active       bool
		returnedUser User
	)

	q := fmt.Sprintf("SELECT email, active FROM users WHERE obfuscated_id='%s'", obfuscatedID)
	err := c.Database.QueryRow(q).Scan(&email, &active)
	if err != nil {
		return returnedUser, err
	}

	returnedUser.Email = email
	returnedUser.Active = active

	return returnedUser, nil
}

// SerializedUser - returns publicly accessible serialized user struct
func (u *User) SerializedUser(c *config.ApplicationContext) ([]byte, error) {
	s := make(map[string]string)
	s["obfuscated_id"] = u.ObfuscatedID
	s["email"] = u.Email
	s["active"] = strconv.FormatBool(u.Active)
	return json.Marshal(s)
}
