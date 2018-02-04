package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func testSimpleRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func testNotFoundRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusPermanentRedirect)
}

func TestNewRouter(t *testing.T) {
	var res *http.Response
	var err error
	var routes = Routes{
		Route{"Info", "GET", "/", testSimpleRoute},
	}

	r := NewRouter(routes, testNotFoundRoute)
	ts := httptest.NewServer(r)

	res, err = http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Router failed to receive status 200 for active route. Instead got %v", res.StatusCode)
	}

	res, err = http.Get(ts.URL + "/invalid")
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusPermanentRedirect {
		t.Errorf("Router failed to redirect to not found route.")
	}
}
