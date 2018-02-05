package models

import (
	"testing"

	"github.com/timkellogg/five_three_one/services/authentication"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestUsersCreate(t *testing.T) {
	var err error

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("UsersCreate failed to connect to db: %s", err)
	}
	defer db.Close()

	auth := &authentication.AuthService{}

	u := User{Email: "test@test.com"}
	password := "password"

	rows := sqlmock.NewRows([]string{"id"}).AddRow("1")

	mock.ExpectQuery("INSERT INTO").WithArgs(u.Email, sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(rows)

	err = u.CreateUser(db, auth, password)
	if err != nil {
		t.Error(err)
	}

	if u.EncryptedPassword == "" {
		t.Error("Did not set encrypted password")
	}

	if u.ObfuscatedID == "" {
		t.Error("Did not set obfuscated id")
	}
}
