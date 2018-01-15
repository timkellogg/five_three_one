package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/timkellogg/five_three_one/config"
)

func TestMain(m *testing.M) {
	app.Initialize(config.TestConfig)

	code := m.Run()

	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	app.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d", expected, actual)
	}
}

func clearTables(tables []string) {
	for table := range tables {
		app.DB.Exec("TRUNCATE %s;", table)
	}
}

func TestControllersInfo(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/info", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	var major = "0"
	var minor = "0"
	var patch = "1"

	if m["major"] != major {
		t.Errorf("Expected major to be '%s' but it was %s", major, m["major"])
	}
	if m["minor"] != minor {
		t.Errorf("Expected minor to be '%s' but it was %s", minor, m["minor"])
	}
	if m["patch"] != patch {
		t.Errorf("Expected patch to be '%s' but it was %s", patch, m["patch"])
	}
}

func TestControllersErrors404(t *testing.T) {
	var expectedCode = "ResourceNotFound"
	var expectedMessage = "There is no resource at that location"

	req, _ := http.NewRequest("GET", "/api/norouteexists", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["code"] != expectedCode {
		t.Errorf("Expected code to be '%s' but found %s", expectedCode, m["code"])
	}
	if m["message"] != expectedMessage {
		t.Errorf("Expected message to be '%s' but found %s", expectedMessage, m["message"])
	}
}
