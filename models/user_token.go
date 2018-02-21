package models

import (
	"time"

	"github.com/lib/pq"
	"github.com/timkellogg/five_three_one/config"
)

// UserToken - user's refresh tokens
type UserToken struct {
	ID        int64
	Token     string      `json:"token" db:"token"`
	UserID    int64       `json:"user_id" db:"user_id"`
	Active    bool        `json:"active" db:"active"`
	CreatedAt pq.NullTime `json:"created_at" db:"created_at"`
	UpdatedAt pq.NullTime `json:"updated_at" db:"updated_at"`
}

// UserTokenCreate - saves token to db
func (ut *UserToken) UserTokenCreate(c *config.ApplicationContext) (*UserToken, error) {
	err := c.Database.
		QueryRow("INSERT INTO user_tokens (user_id, token, created_at) VALUES($1,$2,$3) RETURNING id, token, user_id, active",
			ut.UserID, ut.Token, pq.NullTime{Valid: true, Time: time.Now()}).
		Scan(&ut.ID, &ut.Token, &ut.UserID, &ut.Active)
	if err != nil {
		return ut, err
	}

	return ut, nil
}

// UserTokenInvalidate - set user token's active to false
func (ut *UserToken) UserTokenInvalidate(c *config.ApplicationContext) (*UserToken, error) {
	err := c.Database.
		QueryRow("UPDATE user_tokens SET active = false, updated_at = $1 WHERE user_id = $2 RETURNING id, token, user_id, active",
			pq.NullTime{Valid: true, Time: time.Now()}, ut.UserID).
		Scan(&ut.ID, &ut.Token, &ut.UserID, &ut.Active)
	if err != nil {
		return ut, err
	}

	return ut, nil
}

// UserTokenUser - returns user from token
func (ut *UserToken) UserTokenUser(c *config.ApplicationContext) (*User, error) {
	var user User

	err := c.Database.
		QueryRow("SELECT obfuscated_id, email, active, created_at FROM users WHERE id = $1", ut.UserID).
		Scan(&user.ObfuscatedID, &user.Email, &user.Active, &user.CreatedAt)
	if err != nil {
		return &user, err
	}

	return &user, nil
}
