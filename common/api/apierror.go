package api

// APIError - struct that wraps information about an error while processing
// an API call
type APIError struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

// NewAPIError creates a new APIError with the passed HTTP status and message.
func NewAPIError(status int, message string) *APIError {
	return &APIError{
		Status:  status,
		Message: message,
	}
}

// FromError creates an APIError based on a passed error
func FromError(err error) *APIError {
	if err == nil {
		return nil
	}

	apiErr := &APIError{
		Status:  400,
		Message: err.Error(),
	}

	return apiErr
}
