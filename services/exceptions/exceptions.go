package exceptions

import (
	"encoding/json"
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

// UserCreateError - user could not be created from attributes
var UserCreateError = ApplicationException{
	Code:       "UserCouldNotBeCreated",
	Message:    "User attributes were not valid",
	HTTPStatus: http.StatusUnprocessableEntity,
}

// UserNotAuthorized - user does not have a valid Authorization header
var UserNotAuthorized = ApplicationException{
	Code:       "UserNotAuthorized",
	Message:    "User does not valid token",
	HTTPStatus: http.StatusUnauthorized,
}

// RefreshTokenCreateError - signifies when a user could have a refresh token generated
var RefreshTokenCreateError = ApplicationException{
	Code:       "RefreshTokenNotCreated",
	Message:    "Could not generate refresh token for user",
	HTTPStatus: http.StatusInternalServerError,
}

// UnknownTokenGrantType - grant type of oauth not recognized
var UnknownTokenGrantType = ApplicationException{
	Code:       "GrantTypeNotRecognized",
	Message:    "Grant type is not recognized",
	HTTPStatus: http.StatusUnprocessableEntity,
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
