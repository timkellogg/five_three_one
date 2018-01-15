package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/timkellogg/five_three_one/app/models"
)

// InfoShow - prints out details about api
func InfoShow(w http.ResponseWriter, r *http.Request) {
	info := models.Info{Major: "0", Minor: "0", Patch: "1"}

	i, err := json.Marshal(info)
	if err != nil {
		handleError(err, models.JSONParseError, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(i)
}
