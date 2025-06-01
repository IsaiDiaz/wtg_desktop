package utils

import "net/http"

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewNotFound(msg string) *AppError {
	return &AppError{Code: http.StatusNotFound, Message: msg}
}

func NewInternalError(msg string) *AppError {
	return &AppError{Code: http.StatusInternalServerError, Message: msg}
}

func WrapAppError(err error) *AppError {
	switch e := err.(type) {
	case *AppError:
		return e
	default:
		return NewInternalError(err.Error())
	}
}
