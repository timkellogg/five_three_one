package models

import (
	"time"

	"github.com/timkellogg/five_three_one/config"
)

// UserSecret - stores client_id and client_secret for auth
type UserSecret struct {
	ID           int64
	UserID       int64     `json:"user_id" db:"user_id"`
	ClientID     string    `json:"client_id" db:"client_id"`
	ClientSecret string    `json:"client_secret" db:"client_secret"`
	Active       bool      `json:"active" db:"active"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// Save - persists the user secret
func (us *UserSecret) SaveUserSecret(c *config.ApplicationContext) (*UserSecret, error) {
	us.ClientSecret = c.Auth.UniqueString()
	us.ClientID = c.Auth.UniqueString()
	us.Active = true
	us.CreatedAt = time.Now()

	err := c.Database.
		QueryRow("INSERT INTO user_secrets (user_id, client_id, client_secret, active, created_at) VALUES($1,$2,$3,$4,$5) RETURNING *",
			us.UserID, us.ClientSecret, us.ClientID, us.Active, us.CreatedAt).
		Scan(&us.ID, &us.UserID, &us.ClientID, &us.ClientSecret, &us.Active, &us.CreatedAt, &us.UpdatedAt)
	if err != nil {
		return us, err
	}

	return us, nil
}

// User - returns UserSecret user
func (us *UserSecret) User(c *config.ApplicationContext) (User, error) {
	return User{}, nil
}
