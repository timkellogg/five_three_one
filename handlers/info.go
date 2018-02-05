package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/timkellogg/five_three_one/config"
	"github.com/timkellogg/five_three_one/models"
)

// InfoShow - prints out details about api
func InfoShow(c *config.ApplicationContext, w http.ResponseWriter, r *http.Request) {
	info := models.Info{
		Major: models.Major,
		Minor: models.Minor,
		Patch: models.Patch,
	}

	i, err := json.Marshal(info)
	if err != nil {
		handleError(err, JSONParseError, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(i)
}
