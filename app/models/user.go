package models

import "errors"

// User - a consumer of the application
type User struct {
	firstName string
}

// Validate - validate user fields
func (u User) Validate() error {
	return errors.New("Not Implemented")
}
