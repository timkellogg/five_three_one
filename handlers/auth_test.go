package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	defer context.TruncateDBTables()

	_, err := testUser.CreateUser(&context)
	if err != nil {
		t.Error(err)
	}

	url := server.URL + "/api/auth/login"
	payload := `{"email": "test@test.com", "password": "password"}`
	reader := strings.NewReader(payload)

	res, err := http.Post(url, "application/json", reader)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success status expected but instead got %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	var response LoginResponse
	json.Unmarshal(body, &response)

	if response.AccessToken == "" {
		t.Error("Expected access token to be defined but was empty")
	}

	if response.RefreshToken == "" {
		t.Error("Expected refresh token to be defined but was empty")
	}
}
