package handlers

import (
	"encoding/json"
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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	var response UsersResponse
	json.Unmarshal(body, &response)

	if res.StatusCode != 201 {
		t.Errorf("Success status expected but instead got: %d", res.StatusCode)
	}

	if response.Active != true {
		t.Errorf("Expected active to be true but was: %v", response.Active)
	}

	if response.Email != "test@test.com" {
		t.Errorf("Expected email to be test@test.com but was: %v", response.Email)
	}

	if response.ObfuscatedID == "" {
		t.Error("Expected obfuscated_id to be set but was empty")
	}
}

func TestUsersShow(t *testing.T) {
	defer context.TruncateDBTables()

	url := server.URL + "/api/users/me"

	_, err := testUser.CreateUser(&context)
	if err != nil {
		t.Error(err)
	}

	res, err := http.Get(url)
	if err != nil {
		t.Error(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	var response UsersResponse
	json.Unmarshal(body, &response)

	if res.StatusCode != 200 {
		t.Errorf("Expected status 200 but received %d", res.StatusCode)
	}

	if response.Active != true {
		t.Errorf("Expected response to be active but was %v", response.Active)
	}

	if response.Email != "test@test.com" {
		t.Errorf("Expected email to be test@test.com but was: %v", response.Email)
	}
}
