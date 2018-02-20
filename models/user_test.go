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
