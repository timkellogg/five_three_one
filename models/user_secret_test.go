package models

import "testing"

func TestCreateUserSecret(t *testing.T) {
	defer context.TruncateDBTables()

	testUser.CreateUser(&context)

	userSecret := UserSecret{UserID: testUser.ID}

	returnedUserSecret, err := userSecret.CreateUserSecret(&context)
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

	returnedUserSecret, err := userSecret.CreateUserSecret(&context)
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

func TestUserSecretFindByID(t *testing.T) {
	defer context.TruncateDBTables()

	testUser.CreateUser(&context)

	userSecret := UserSecret{UserID: testUser.ID}

	_, err := userSecret.CreateUserSecret(&context)
	if err != nil {
		t.Error(err)
	}

	returnedUserSecret, err := userSecret.UserSecretFindByID(&context)
	if err != nil {
		t.Error(err)
	}

	if returnedUserSecret.ClientID == "" {
		t.Error("Failed to return client id")
	}

	if returnedUserSecret.ClientSecret == "" {
		t.Error("Failed to return client secret")
	}
}
