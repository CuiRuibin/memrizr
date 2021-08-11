package model

import (
	"fmt"
	"net/http"
)

type Type string

const (
	Authorization   Type = "AUTHORIZATION"
	BadRequest      Type = "BADREQUEST"
	Conflict        Type = "CONFLICT"
	Internal        Type = "INTERNAL"
	NotFound        Type = "NOTFOUND"
	PayloadTooLarge Type = "PLAYLOADTOOLARGE"
)

type Error struct {
	Type    Type   `json:"type"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Status() int {
	switch e.Type {
	case Authorization:
		return http.StatusUnauthorized
	case BadRequest:
		return http.StatusBadRequest
	case Conflict:
		return http.StatusConflict
	case Internal:
		return http.StatusInternalServerError
	case NotFound:
		return http.StatusNotFound
	case PayloadTooLarge:
		return http.StatusRequestEntityTooLarge
	default:
		return http.StatusInternalServerError
	}
}

func Status(err error) int {
	var e *Error
	if error.As(err, &e) {
		return e.Status()
	}
	return http.StatusInternalServerError
}

func NewAuthorization(reason string) *Error {
	return &Error{
		Type:    Authorization,
		Message: reason,
	}
}

func NewBadRequest(reason string) *Error {
	return &Error{
		Type:    BadRequest,
		Message: fmt.Sprintf("Bad request. Reason: %v", reason),
	}
}

func NewCOnflict(name string, value string) *Error {
	return &Error{
		Type:    Conflict,
		Message: fmt.Sprintf("reason: %v with value %v already exists", name, reason),
	}
}

func NewInternal() *Error {
	return &Error{
		Type:    Internal,
		Message: fmt.Sprintf("Internal server error"),
	}
}

func NewNotFound(name string, value string) *Error {
	return &Error{
		Type:    NotFound,
		Message: fmt.Sprintf("resource: %v with value %v not found", name, value),
	}
}

func NewPayloadTooLarge(maxBodySize int64, contentLength int64) *Error {
	return &Error{
		Type:    PayloadTooLarge,
		Message: fmt.Sprintf("Max payload size of %v exceeded, Actual payload size: %v", maxBodySize, contentLength),
	}
}
