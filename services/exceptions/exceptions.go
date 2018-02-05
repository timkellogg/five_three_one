package exceptions

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// ApplicationException - application error messages and codes
type ApplicationException struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	HTTPStatus int    `json:"-"`
}

// JSONParseError - occurs when cannot marshal json object
var JSONParseError = ApplicationException{
	Code:       "JSONParseError",
	Message:    "Something went wrong",
	HTTPStatus: http.StatusInternalServerError,
}

// ResourceNotFoundError - could not find resource
var ResourceNotFoundError = ApplicationException{
	Code:       "ResourceNotFound",
	Message:    "There is no resource at that location",
	HTTPStatus: http.StatusNotFound,
}

// NotImplementedError - endpoint is not finished yet
var NotImplementedError = ApplicationException{
	Code:       "ResourceNotImplemented",
	Message:    "Endpoint is not implemented yet",
	HTTPStatus: http.StatusNotImplemented,
}

// TokenCreateError - something went wrong with creating JWT
var TokenCreateError = ApplicationException{
	Code:       "TokenNotCreated",
	Message:    "Token was not valid",
	HTTPStatus: http.StatusInternalServerError,
}

func handleError(e error, appError ApplicationException, w http.ResponseWriter) {
	log.Println(e)
	log.Println(appError.Code)

	ae, err := json.Marshal(appError)
	if err != nil {
		log.Fatal(e)
	}

	w.WriteHeader(appError.HTTPStatus)
	w.Write(ae)
}

// Errors404 - Not found handler
func Errors404(w http.ResponseWriter, r *http.Request) {
	err := errors.New("Resource Not Found")
	handleError(err, ResourceNotFoundError, w)
}
