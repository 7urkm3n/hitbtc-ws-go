package errors

import (
	"errors"
	"net/http"
)

// RestErr type struct
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

// NewErr returns new error
func NewErr(msg string) error {
	return errors.New(msg)
}

// NewBadRequestErr returns *RestErr
func NewBadRequestErr(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

// NewNotFoundErr returns *RestErr
func NewNotFoundErr(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

// NewInternalServerErr returns *RestErr
func NewInternalServerErr(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
