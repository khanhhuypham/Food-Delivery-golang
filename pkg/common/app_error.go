package common

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Causes  error  `json:"-"` // lỗi gốc của chương trình

	// Debug information
	//
	// This field is often not exposed to protect against leaking
	// sensitive information.
	//
	// example: SQL field "foo" is not a bool.
	DebugField string `json:"debug,omitempty"`
}

func (e *AppError) RootCauses() error {
	if err, ok := e.Causes.(*AppError); ok {
		return err.RootCauses()
	}
	return e.Causes
}

func (e *AppError) Error() string {
	return e.RootCauses().Error()
}

func ErrUnauthorized(cause error) *AppError {
	return &AppError{Causes: cause, Status: http.StatusUnauthorized, Message: "you have no authorization"}
}

func ErrForbidden(cause error) *AppError {
	return &AppError{Causes: cause, Status: http.StatusForbidden, Message: "you have no permission"}
}

func ErrBadRequest(cause error) *AppError {
	return &AppError{Causes: cause, Status: http.StatusBadRequest, Message: "invalid request"}
}

func ErrNotFound(cause error) *AppError {
	return &AppError{Causes: cause, Status: http.StatusNotFound, Message: "not found"}
}

func ErrDB(cause error) *AppError {
	return &AppError{Causes: cause, Status: http.StatusInternalServerError, Message: "something went wrong with Database"}
}

func ErrInternal(cause error) *AppError {
	return &AppError{Causes: cause, Status: http.StatusInternalServerError, Message: "Internal server error"}
}

func ErrEntityNotFound(entity string, cause error) *AppError {
	return &AppError{Causes: cause, Status: http.StatusBadRequest, Message: fmt.Sprintf("can't find %s", entity)}
}

func ErrConflict(cause error) *AppError {
	return &AppError{Causes: cause, Status: http.StatusConflict, Message: "The resource could not be created due to a conflict"}
}

func (e *AppError) WithDebug(debug string) *AppError {
	e.DebugField = debug
	return e
}
