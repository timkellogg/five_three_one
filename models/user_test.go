package models

import (
	"testing"
)

func TestUsersCreate(t *testing.T) {
	defer context.TruncateDBTables()

	user, err := testUser.CreateUser(&context)
	if err != nil {
		t.Error(err)
	}

	if user.Email != "test@test.com" {
		t.Errorf("Expected email to be test@test.com but was %s", user.Email)
	}

	if user.Active != true {
		t.Errorf("Expected active to be true but was %v", user.Active)
	}
}

func TestUsersFindByEmail(t *testing.T) {
	defer context.TruncateDBTables()

	_, err := testUser.CreateUser(&context)
	if err != nil {
		t.Error(err)
	}

	user, err := testUser.FindByEmail(&context)
	if err != nil || user.Email != "test@test.com" {
		t.Error("Could not find user")
	}
}

func TestUsersFindByObfuscatedID(t *testing.T) {
	defer context.TruncateDBTables()

	user, err := testUser.CreateUser(&context)
	if err != nil {
		t.Error(err)
	}

	returnedUser, err := user.FindByObfuscatedID(&context)
	if err != nil || returnedUser.Email != "test@test.com" {
		t.Errorf("Could not find user: %v", returnedUser.ObfuscatedID)
	}
}

func TestUserActiveToken(t *testing.T) {
	defer context.TruncateDBTables()

	user, err := testUser.CreateUser(&context)
	if err != nil {
		t.Error(err)
	}

	_, err = user.UserActiveToken(&context)
	if err == nil {
		t.Errorf("Expected UserActiveToken to be empty. %v", err)
	}

	userToken := UserToken{
		Token:  "token",
		UserID: 1,
	}

	returnedUserToken, err := userToken.UserTokenCreate(&context)
	if err != nil {
		t.Error(err)
	}

	if returnedUserToken.Active != true {
		t.Error("Expected user token to be active but was inactive")
	}

	if returnedUserToken.Token == "" {
		t.Error("Expected user token to have an active token but was empty")
	}
}
