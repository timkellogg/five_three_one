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
}
