package models

import (
	"time"

	"github.com/lib/pq"

	"github.com/timkellogg/five_three_one/config"
)

// UserSecret - stores client_id and client_secret for auth
type UserSecret struct {
	ID           int64
	UserID       int64       `json:"user_id" db:"user_id"`
	ClientID     string      `json:"client_id" db:"client_id"`
	ClientSecret string      `json:"client_secret" db:"client_secret"`
	Active       bool        `json:"active" db:"active"`
	CreatedAt    pq.NullTime `json:"created_at" db:"created_at"`
	UpdatedAt    pq.NullTime `json:"updated_at" db:"updated_at"`
}

// CreateUserSecret - persists the user secret
func (us *UserSecret) CreateUserSecret(c *config.ApplicationContext) (*UserSecret, error) {
	us.ClientSecret = c.Auth.UniqueString()
	us.ClientID = c.Auth.UniqueString()
	us.Active = true
	us.CreatedAt = pq.NullTime{Valid: true, Time: time.Now()}

	err := c.Database.
		QueryRow("INSERT INTO user_secrets (user_id, client_id, client_secret, active, created_at) VALUES($1,$2,$3,$4,$5) RETURNING *",
			us.UserID, us.ClientSecret, us.ClientID, us.Active, us.CreatedAt).
		Scan(&us.ID, &us.UserID, &us.ClientID, &us.ClientSecret, &us.Active, &us.CreatedAt, &us.UpdatedAt)
	if err != nil {
		return us, err
	}

	return us, nil
}

// UserSecretUser - returns UserSecret user
func (us *UserSecret) UserSecretUser(c *config.ApplicationContext) (*User, error) {
	var user User

	err := c.Database.
		QueryRow("SELECT obfuscated_id, email, active, created_at FROM users WHERE id = $1", us.UserID).
		Scan(&user.ObfuscatedID, &user.Email, &user.Active, &user.CreatedAt)
	if err != nil {
		return &user, err
	}

	return &user, nil
}

// UserSecretFindByID - returns UserSecret by ID
func (us *UserSecret) UserSecretFindByID(c *config.ApplicationContext) (*UserSecret, error) {
	err := c.Database.
		QueryRow("SELECT client_id, client_secret FROM user_secrets WHERE user_id = $1 AND active = true", us.UserID).
		Scan(&us.ClientID, &us.ClientSecret)
	if err != nil {
		return us, err
	}

	return us, nil
}
