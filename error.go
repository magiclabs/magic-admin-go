package magic

import (
	"encoding/json"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type ErrorCode string

type Error struct {
	Response
}

// Error serializes the error object to JSON and returns it as a string.
func (e *Error) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

// RateLimitError occurs when in case if API is hit with too many requests.
type RateLimitingError struct {
	Err *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *RateLimitingError) Error() string {
	return e.Err.Error()
}

// BadRequestError occurs with not well formed request.
type BadRequestError struct {
	Err *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *BadRequestError) Error() string {
	return e.Err.Error()
}

// AuthenticationError occurs if request is not authorized to proceed.
type AuthenticationError struct {
	Err *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *AuthenticationError) Error() string {
	return e.Err.Error()
}

// ForbiddenError occurs if request is not permitted to be executed.
type ForbiddenError struct {
	Err *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *ForbiddenError) Error() string {
	return e.Err.Error()
}

// APIError default unrecognized by any other errors.
type APIError struct {
	Err *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *APIError) Error() string {
	return e.Err.Error()
}

// APIConnectionError occurs if request is not permitted to be executed.
type APIConnectionError struct {
	Err error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *APIConnectionError) Error() string {
	return e.Err.Error()
}

// Wraps error into appropriate type.
func WrapError(r *resty.Response, err *Error) error {
	switch r.StatusCode() {
	case http.StatusForbidden:
		return &ForbiddenError{Err: err}
	case http.StatusBadRequest:
		return &BadRequestError{Err: err}
	case http.StatusTooManyRequests:
		return &RateLimitingError{Err: err}
	case http.StatusUnauthorized:
		return &AuthenticationError{Err: err}
	default:
		return &APIError{Err: err}
	}
}
