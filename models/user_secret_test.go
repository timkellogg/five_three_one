package models

import "testing"

func TestSaveUserSecret(t *testing.T) {
	defer context.TruncateDBTables()

	testUser.CreateUser(&context)

	userSecret := UserSecret{UserID: testUser.ID}

	returnedUserSecret, err := userSecret.SaveUserSecret(&context)
	if err != nil {
		t.Error(err)
	}

	if returnedUserSecret.ClientID == "" {
		t.Error("UserSecret did not set client id")
	}

	if returnedUserSecret.ClientSecret == "" {
		t.Error("UserSecret did not set client secret")
	}
}
func TestUserSecretUser(t *testing.T) {
	defer context.TruncateDBTables()

	testUser.CreateUser(&context)

	userSecret := UserSecret{UserID: testUser.ID}

	returnedUserSecret, err := userSecret.SaveUserSecret(&context)
	if err != nil {
		t.Error(err)
	}

	returnedUser, err := returnedUserSecret.UserSecretUser(&context)
	if err != nil {
		t.Error(err)
	}

	if returnedUser.Email != testUser.Email {
		t.Error("Could not retrieve user from user secret")
	}
}
