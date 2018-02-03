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

// ResourceNotFoundError - could not find resource
var ResourceNotFoundError = ApplicationError{
	Code:       "ResourceNotFound",
	Message:    "There is no resource at that location",
	HTTPStatus: http.StatusNotFound,
}

// NotImplementedError - endpoint is not finished yet
var NotImplementedError = ApplicationError{
	Code:       "ResourceNotImplemented",
	Message:    "Endpoint is not implemented yet",
	HTTPStatus: http.StatusNotImplemented,
}

// TokenCreateError - something went wrong with creating JWT
var TokenCreateError = ApplicationError{
	Code:       "TokenNotCreated",
	Message:    "Token was not valid",
	HTTPStatus: http.StatusInternalServerError,
}
