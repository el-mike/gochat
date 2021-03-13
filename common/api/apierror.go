package api

import (
	"fmt"
	"net/http"
)

// APIError - struct that wraps information about an error while processing
// an API call.

// ErrorType holds a general type for the error.
type ErrorType string

// ErrorCode holds an app-specific code for the error.
type ErrorCode string

// Map of valid error types (ErrorType).
const (
	AuthorizationError ErrorType = "AUTHORIZATION"
	NotFoundError      ErrorType = "NOT_FOUND"
	InternalError      ErrorType = "INTERNAL"
	BadRequestError    ErrorType = "BAD_REQUEST"
)

// APIError holds a custom error for the application,
// which is later consumed by clients.
type APIError struct {
	Status    int       `json:"status"`
	Type      ErrorType `json:"type"`
	ErrorCode ErrorCode `json:"errorCode"`
	Message   string    `json:"message"`
}

// Error() satisfies standard Error interface.
func (e *APIError) Error() string {
	return e.Message
}

func getHttpStatusCode(errorType ErrorType) int {
	switch errorType {
	case AuthorizationError:
		return http.StatusUnauthorized
	case NotFoundError:
		return http.StatusNotFound
	case InternalError:
		return http.StatusInternalServerError
	case BadRequestError:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}

}

// Returns APIError related to general authorization problem.
func NewAuthorizationError(source error) *APIError {
	return &APIError{
		Status:    getHttpStatusCode(AuthorizationError),
		Type:      AuthorizationError,
		ErrorCode: "auth",
		Message:   source.Error(),
	}
}

// Returns APIError related to expired token.
func NewTokenExpiredError() *APIError {
	return &APIError{
		Status:    getHttpStatusCode(AuthorizationError),
		Type:      AuthorizationError,
		ErrorCode: "auth/token-expired",
		Message:   "Token expired, please log in again.",
	}
}

// Returns APIError related to login credentials being incorrect.
func NewLoginCredentialsIncorrectError() *APIError {
	return &APIError{
		Status:    getHttpStatusCode(AuthorizationError),
		Type:      AuthorizationError,
		ErrorCode: "auth/login-credentials-incorrect",
		Message:   "Email or password is incorrect.",
	}
}

// Return APIError related to malformed token.
func NewTokenMalforedError() *APIError {
	return &APIError{
		Status:    getHttpStatusCode(AuthorizationError),
		Type:      AuthorizationError,
		ErrorCode: "auth/token-malformed",
		Message:   "Token malformed.",
	}
}

// Returns APIError related to an entity which cannot be found.
func NewNotFoundError(modelName string) *APIError {
	return &APIError{
		Status:    getHttpStatusCode(NotFoundError),
		Type:      NotFoundError,
		ErrorCode: "not-found",
		Message:   fmt.Sprintf("resource: %v not found", modelName),
	}
}

// Returns APIError related to internal, runtime problem.
func NewInternalError(source error) *APIError {
	return &APIError{
		Status:    getHttpStatusCode(InternalError),
		Type:      InternalError,
		ErrorCode: "internal",
		Message:   source.Error(),
	}
}

// Returns APIError related to clients's request being incorrect.
func NewBadRequestError(source error) *APIError {
	return &APIError{
		Status:    getHttpStatusCode(BadRequestError),
		Type:      BadRequestError,
		ErrorCode: "bad-request",
		Message:   source.Error(),
	}
}

// Returns pair of values (http status, error) that satisfies Gin's
// handler response.
func ResponseFromError(err *APIError) (int, *APIError) {
	return err.Status, err
}
