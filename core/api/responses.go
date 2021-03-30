package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Returns pair of values (http status, error) that satisfies Gin's
// handler response.
func ResponseFromError(err *APIError) (int, *APIError) {
	return err.Status, err
}

// Returns pair of values (http status, data) that satisfies Gin's
// handler response.
func GetSuccessResponse(data interface{}) (int, interface{}) {
	if data == nil {
		return GetEmptySuccessResponse()
	}

	return http.StatusOK, data
}

// Returns pair of values (http status, data) that satisfies Gin's
// handler response.
func GetEmptySuccessResponse() (int, interface{}) {
	return http.StatusOK, gin.H{}
}
