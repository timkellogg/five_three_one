package routing

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/timkellogg/five_three_one/config"
	"github.com/timkellogg/five_three_one/services/authentication"
	"github.com/timkellogg/five_three_one/services/database"
	"github.com/timkellogg/five_three_one/services/session"
)

type message struct {
	body string
}

// TODO: figure out how to decouple
var context = config.ApplicationContext{
	Database: database.NewDatabase().Store,
	Session:  session.NewSession().Memcache,
	Auth:     authentication.AuthService{},
}

func testSimpleRoute(context *config.ApplicationContext, w http.ResponseWriter, r *http.Request) {
	msg := message{body: "Hello World"}
	m, _ := json.Marshal(msg)

	w.Write(m)
	w.WriteHeader(http.StatusOK)
}

func testNotFoundRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusPermanentRedirect)
}

func TestNewRouter(t *testing.T) {
	var res *http.Response
	var err error
	var routes = Routes{
		Route{"Info", "GET", "/info", testSimpleRoute},
	}

	r := NewRouter(&context, routes, testNotFoundRoute)
	ts := httptest.NewServer(r)

	// valid api routes
	res, err = http.Get(ts.URL + "/api/info")
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Router failed to recieve status 200 for active route. Instead got %v", res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Error("Middlewares failed to set content type to application/json")
	}

	// unknown routes
	res, err = http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusPermanentRedirect {
		t.Errorf("Router failed to recieve status 200 for active route. Instead got %v", res.StatusCode)
	}
}
