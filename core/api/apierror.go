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
	AuthorizationError  ErrorType = "AUTHORIZATION"
	AuthenticationError ErrorType = "AUTHENTICATION"
	NotFoundError       ErrorType = "NOT_FOUND"
	InternalError       ErrorType = "INTERNAL"
	BadRequestError     ErrorType = "BAD_REQUEST"
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
	case AuthenticationError:
		return http.StatusForbidden
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

// NewAuthorizationError - returns APIError related to general authorization problem.
func NewAuthorizationError(source error) *APIError {
	return &APIError{
		Status:    getHttpStatusCode(AuthorizationError),
		Type:      AuthorizationError,
		ErrorCode: "auth",
		Message:   source.Error(),
	}
}

// NewTokenExpiredError - returns APIError related to expired token.
func NewTokenExpiredError() *APIError {
	return &APIError{
		Status:    getHttpStatusCode(AuthorizationError),
		Type:      AuthorizationError,
		ErrorCode: "auth/token-expired",
		Message:   "Token expired, please log in again.",
	}
}

// NewLoginCredentialsIncorrectError - returns APIError related to
// login credentials being incorrect.
func NewLoginCredentialsIncorrectError() *APIError {
	return &APIError{
		Status:    getHttpStatusCode(AuthorizationError),
		Type:      AuthorizationError,
		ErrorCode: "auth/login-credentials-incorrect",
		Message:   "Email or password is incorrect.",
	}
}

// NewTokenMalforedError - returns APIError related to malformed token.
func NewTokenMalforedError() *APIError {
	return &APIError{
		Status:    getHttpStatusCode(AuthorizationError),
		Type:      AuthorizationError,
		ErrorCode: "auth/token-malformed",
		Message:   "Token malformed.",
	}
}

// NewAccessDeniedError - returns APIError related to missing permissions.
func NewAccessDeniedError(resource string, action string) *APIError {
	return &APIError{
		Status:    getHttpStatusCode(AuthenticationError),
		Type:      AuthenticationError,
		ErrorCode: "auth/access-denied",
		Message:   fmt.Sprintf("Acces denied - You can't perform %v on resource %v", action, resource),
	}
}

// NewNotFoundError - returns APIError related to an entity which cannot be found.
func NewNotFoundError(modelName string) *APIError {
	return &APIError{
		Status:    getHttpStatusCode(NotFoundError),
		Type:      NotFoundError,
		ErrorCode: "not-found",
		Message:   fmt.Sprintf("resource: %v not found", modelName),
	}
}

// NewInternalError - returns APIError related to internal, runtime problem.
func NewInternalError(source error) *APIError {
	return &APIError{
		Status:    getHttpStatusCode(InternalError),
		Type:      InternalError,
		ErrorCode: "internal",
		Message:   source.Error(),
	}
}

// NewBadRequestError - returns APIError related to clients's request being incorrect.
func NewBadRequestError(source error) *APIError {
	return &APIError{
		Status:    getHttpStatusCode(BadRequestError),
		Type:      BadRequestError,
		ErrorCode: "bad-request",
		Message:   source.Error(),
	}
}
