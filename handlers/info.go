package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/timkellogg/five_three_one/services/exceptions"

	"github.com/timkellogg/five_three_one/config"
	"github.com/timkellogg/five_three_one/models"
)

// InfoShowResponse - structure of InfoShow response
type InfoShowResponse struct {
	Version string
}

// InfoShow - prints out details about api
func InfoShow(c *config.ApplicationContext, w http.ResponseWriter, r *http.Request) {
	responseStructure := InfoShowResponse{
		Version: models.Major + "." + models.Minor + "." + models.Patch,
	}

	response, err := json.Marshal(responseStructure)
	if err != nil {
		handleError(err, exceptions.JSONParseError, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
