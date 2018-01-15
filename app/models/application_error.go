package models

import (
	"net/http"
)

// ApplicationError - application error messages and codes
type ApplicationError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	HTTPStatus int    `json:"-"`
}

// JSONParseError - occurs when cannot marshal json object
var JSONParseError = ApplicationError{
	Code:       "JSONParseError",
	Message:    "Something went wrong",
	HTTPStatus: http.StatusInternalServerError,
}

// ResourceNotFound - could not find resource
var ResourceNotFound = ApplicationError{
	Code:       "ResourceNotFound",
	Message:    "There is no resource at that location",
	HTTPStatus: http.StatusNotFound,
}
