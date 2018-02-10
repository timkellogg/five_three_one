package models

import (
	"testing"
	"time"
)

func TestUsersCreate(t *testing.T) {
	defer context.TruncateDBTables()

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

func TestSerializedUser(t *testing.T) {
	defer context.TruncateDBTables()

	u := User{
		ID:                1,
		Email:             "test@test.com",
		ObfuscatedID:      "some-string",
		Password:          "password",
		EncryptedPassword: "a90ahind",
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		Active:            true,
	}

	s, err := u.SerializedUser(&context)
	if err != nil {
		t.Error(err)
	}

	expectedResponse := `{"active":"true","email":"test@test.com","obfuscated_id":"some-string"}`

	if string(s) != expectedResponse {
		t.Errorf("User serialized %s, but should have returned: %s", s, expectedResponse)
	}
}

func TestUsersFindByObfuscatedID(t *testing.T) {
	defer context.TruncateDBTables()

	var u User

	_, err := testUser.CreateUser(&context)
	if err != nil {
		t.Error(err)
	}

	returnedUser, err := u.FindByObfuscatedID(&context, testUser.ObfuscatedID)
	if err != nil {
		t.Errorf("Could not find user: %s", err)
	}

	if returnedUser.Email != testUser.Email {
		t.Errorf("FindByObfuscatedID returned %s when it should have returned: %s", returnedUser.Email, testUser.Email)
	}
}
