package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestInfo(t *testing.T) {
	res, err := http.Get(server.URL + "/api/info")
	if err != nil {
		t.Error(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	var response InfoShowResponse
	json.Unmarshal(body, &response)

	if res.StatusCode != 200 {
		t.Errorf("Info handler returned a non-200 status of %v", res.StatusCode)
	}

	if response.Version != "0.0.1" {
		t.Errorf("Expected version to be 0.0.1 but was %s", response.Version)
	}
}
