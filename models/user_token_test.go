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

	returnedUserToken, err := userToken.Save(&context)
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

	_, err = userToken2.Save(&context)
	if err == nil {
		t.Error("Token violated active uniqueness constraint")
	}
}

// func TestInvalidate(t *testing.T) {
// 	defer context.TruncateDBTables()

// 	userToken := UserToken{
// 		Token:  "token",
// 		UserID: 1,
// 	}

// 	token, err := userToken.Save(&context)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	token, err = userToken.Invalidate(&context)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	// err = context.Database.QueryRow("SELECT * FROM user_tokens RETURNING *").Scan(&token.Active)

// 	if token.Active != false {
// 		t.Errorf("Expected user token active to be false but was %v", token.Active)
// 	}
// }
