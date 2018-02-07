package handlers

import (
	"net/http"
	"testing"
)

func TestInfo(t *testing.T) {
	res, err := http.Get(server.URL + "/api/info")
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Info handler returned a non-200 status of %v", res.StatusCode)
	}
}
