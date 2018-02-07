package models

import (
	"testing"
)

func TestUsersCreate(t *testing.T) {
	token, err := testUser.CreateUser(&context)
	if err != nil {
		t.Error(err)
	}

	if testUser.EncryptedPassword == "" {
		t.Error("Did not set encrypted password")
	}

	if testUser.ObfuscatedID == "" {
		t.Error("Did not set obfuscated id")
	}

	if token == "" {
		t.Error("Token was invalid")
	}

	r, err := context.Database.Exec("SELECT * FROM users;")
	if err != nil {
		t.Error(err)
	}

	rowsChanged, err := r.RowsAffected()
	if err != nil {
		t.Error(err)
	}

	if rowsChanged != 1 {
		t.Error("User was not persisted")
	}
}
