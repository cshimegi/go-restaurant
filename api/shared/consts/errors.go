package consts

import (
	"errors"
	"net/http"
)

type ApiError struct {
	Code    int
	Message error
}

func (e ApiError) Error() string {
	return e.Message.Error()
}

// 4xx errors
var (
	ErrNotFound     = ApiError{Code: http.StatusNotFound, Message: errors.New("not found")}
	ErrForbidden    = ApiError{Code: http.StatusForbidden, Message: errors.New("forbidden")}
	ErrConflict     = ApiError{Code: http.StatusConflict, Message: errors.New("conflict")}
	ErrBadRequest   = ApiError{Code: http.StatusBadRequest, Message: errors.New("bad request")}
	ErrUnauthorized = ApiError{Code: http.StatusUnauthorized, Message: errors.New("unauthorized")}
)

// 5xx errors
var (
	ErrInternal = ApiError{Code: http.StatusInternalServerError, Message: errors.New("internal server error")}
	ErrUnknown  = ApiError{Code: http.StatusInternalServerError, Message: errors.New("unknown error")}
)
