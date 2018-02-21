package models

import (
	"testing"
)

func TestSave(t *testing.T) {
	defer context.TruncateDBTables()

	userToken := UserToken{
		Token:  "token",
		UserID: 1,
	}

	returnedUserToken, err := userToken.UserTokenCreate(&context)
	if err != nil {
		t.Error(err)
	}

	if returnedUserToken.Active != true {
		t.Errorf("Expexted user token to be true by default but was %v", returnedUserToken.Active)
	}

	userToken2 := UserToken{
		Token:  "token2",
		UserID: 1,
	}

	_, err = userToken2.UserTokenCreate(&context)
	if err == nil {
		t.Error("Token violated active uniqueness constraint")
	}
}

func TestInvalidate(t *testing.T) {
	defer context.TruncateDBTables()

	userToken := UserToken{
		Token:  "token",
		UserID: 1,
	}

	token, err := userToken.UserTokenCreate(&context)
	if err != nil {
		t.Error(err)
	}

	token, err = userToken.UserTokenInvalidate(&context)
	if err != nil {
		t.Error(err)
	}

	if token.Active != false {
		t.Errorf("Expected user token active to be false but was %v", token.Active)
	}
}

func TestUserTokenUser(t *testing.T) {
	defer context.TruncateDBTables()

	user, err := testUser.CreateUser(&context)
	if err != nil {
		t.Error(err)
	}

	userToken := UserToken{
		Token:  "token",
		UserID: user.ID,
	}

	tokenUser, err := userToken.UserTokenCreate(&context)
	if err != nil {
		t.Error(err)
	}

	if tokenUser.UserID != user.ID {
		t.Error("UserTokenUser was unable to return user")
	}
}
