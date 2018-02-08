package handlers

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestUsersCreate(t *testing.T) {
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
