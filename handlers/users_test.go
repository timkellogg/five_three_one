package handlers

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestUsersCreate(t *testing.T) {
	defer context.TruncateDBTables()

	url := server.URL + "/api/users/create"
	payload := `{"email": "test@test.com", "password": "password"}`
	reader := strings.NewReader(payload)

	res, err := http.Post(url, "application/json", reader)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 201 {
		t.Errorf("Success status expected but instead got: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	returnedResponse := string(body)
	if returnedResponse == "" {
		t.Error("Response was empty")
	}
}

func TestUsersShow(t *testing.T) {
	defer context.TruncateDBTables()

	url := server.URL + "/api/users/me"

	_, err := testUser.CreateUser(&context)
	if err != nil {
		t.Errorf("HERE: %v", err)
	}

	res, err := http.Get(url)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected status 200 but received %d", res.StatusCode)
	}
}
