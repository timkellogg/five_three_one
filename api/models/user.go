package models

// User - a consumer of the application
type User struct {
	ID       int64
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUser - saves user to db
func (u *User) CreateUser() error {
	err := Database.QueryRow("INSERT INTO users(email) VALUES($1) RETURNING id", u.Email).Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}
