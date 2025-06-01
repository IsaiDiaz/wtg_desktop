package errors

import "net/http"

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func (e *AppError) Error() string {
	return e.Message
}

func New(code int, msg string, detail ...string) *AppError {
	d := ""
	if len(detail) > 0 {
		d = detail[0]
	}
	return &AppError{Code: code, Message: msg, Detail: d}
}

func Wrap(err error) *AppError {
	if appErr, ok := err.(*AppError); ok {
		return appErr
	}
	return New(http.StatusInternalServerError, "Internal server error", err.Error())
}
