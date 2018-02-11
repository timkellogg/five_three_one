package models

import (
	"time"

	"github.com/timkellogg/five_three_one/config"
)

// UserToken - user's refresh tokens
type UserToken struct {
	ID        int64
	Token     string    `json:"token" db:"token"`
	UserID    int64     `json:"user_id" db:"user_id"`
	Active    bool      `json:"active" db:"active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Save - saves token to db
func (ut *UserToken) Save(c *config.ApplicationContext) (*UserToken, error) {
	ut.CreatedAt = time.Now()

	err := c.Database.
		QueryRow("INSERT INTO user_tokens (user_id, token, created_at) VALUES($1,$2,$3) RETURNING id, token, user_id, active", ut.UserID, ut.Token, ut.CreatedAt).
		Scan(&ut.ID, &ut.Token, &ut.UserID, &ut.Active)
	if err != nil {
		return ut, err
	}

	return ut, nil
}

// Invalidate - setstoken to be not active
func (ut *UserToken) Invalidate(c *config.ApplicationContext) (*UserToken, error) {
	ut.UpdatedAt = time.Now()

	c.Database.Exec("UPDATE user_tokens SET active = false, updated_at = $1 WHERE user_id = $2", ut.UpdatedAt, ut.UserID)
	// if err != nil {
	// 	return ut, err
	// }

	return ut, nil
}
