package models

import "testing"

func TestSaveUserSecret(t *testing.T) {
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
func TestUser(t *testing.T) {
}
