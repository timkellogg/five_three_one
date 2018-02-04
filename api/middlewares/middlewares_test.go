package middlewares

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type message struct {
	body string
}

func testSimpleRoute(w http.ResponseWriter, r *http.Request) {
	msg := message{body: "Hello World"}
	m, _ := json.Marshal(msg)

	w.Write(m)
	w.WriteHeader(http.StatusOK)
}

func TestSetHeaders(t *testing.T) {
	ts := httptest.NewServer(SetHeaders(http.HandlerFunc(testSimpleRoute)))

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	contentType := res.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Error("Middlewares failed to set content type to application/json")
	}

	ts = httptest.NewServer(http.HandlerFunc(testSimpleRoute))
	res, err = http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	contentType = res.Header.Get("Content-Type")
	if contentType == "application/json" {
		t.Errorf("Content Type should not be header without middleware")
	}
}
