package app_error

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"key"`
}

func (e *AppError) RootError() error {
	var err *AppError
	if errors.As(e.RootErr, &err) {
		return err.RootError()
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func NewCompleteAppError(statusCode int, rootError error, message, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    rootError,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewErrorResponse(rootError error, message, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    rootError,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func ErrDB(err error) *AppError {
	return NewCompleteAppError(http.StatusInternalServerError, err, "Something went wrong with db", err.Error(), "DB_ERROR")
}

func ErrInternal(err error) *AppError {
	return NewCompleteAppError(http.StatusInternalServerError, err, "Something went wrong in server", err.Error(), "INTERNAL_ERROR")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "INVALID_REQ_ERROR")
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCompleteAppError(http.StatusNotFound, err, fmt.Sprintf("No %s Found", entity), err.Error(), fmt.Sprintf("%s_NOT_FOUND", strings.ToUpper(entity)))
}

func ErrConflict(field string, err error) *AppError {
	return NewCompleteAppError(http.StatusConflict, err, fmt.Sprintf("%s was exist", field), err.Error(), fmt.Sprintf("%s_EXIST", strings.ToUpper(field)))
}

var ErrRecordNotFound error = errors.New("record Not Found")
