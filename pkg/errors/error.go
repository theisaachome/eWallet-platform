package errors

import "net/http"

const (
	ErrValidation     = "VALIDATION_ERROR"
	ErrUnauthorized   = "UNAUTHORIZED"
	ErrNotFound       = "NOT_FOUND"
	ErrInternalServer = "INTERNAL_SERVER_ERROR"
	BadCredentials    = "BAD_CREDENTIALS"
)

type AppError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func NewNotFoundException(msg string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

func NewUnexpectedError(msg string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}

func NewValidationError(msg string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: msg,
	}
}
