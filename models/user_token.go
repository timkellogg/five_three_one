package models

import "github.com/timkellogg/five_three_one/config"

// UserToken - user's refresh tokens
type UserToken struct {
	ID     int64
	Token  string `json:"token" db:"token"`
	UserID int    `json:"user_id" db:"user_id"`
	Active bool   `json:"active" db:"active"`
}

// Save - saves token to db
func (ut *UserToken) Save(c *config.ApplicationContext) (UserToken, error) {
	var returnedUserToken UserToken

	ut.Active = true

	err := c.Database.QueryRow("INSERT INTO user_tokens (user_id, token) VALUES($1,$2) RETURNING *", ut.UserID, ut.Token).Scan(&ut.ID, &ut.Token, &ut.UserID, &ut.Active)
	if err != nil {
		return returnedUserToken, err
	}

	return returnedUserToken, nil
}

// Invalidate - sets token to be not active
func (ut *UserToken) Invalidate(c *config.ApplicationContext) (UserToken, error) {
	var returnedUserToken UserToken

	ut.Active = false

	err := c.Database.QueryRow("UPDATE user_tokens SET active = false WHERE user_id = $1 AND token = $2",
		ut.UserID, ut.Token).Scan(&returnedUserToken)
	if err != nil {
		return returnedUserToken, err
	}

	return returnedUserToken, nil
}
