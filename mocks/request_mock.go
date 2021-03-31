package mocks

import "net/http"

func GetTestRequest() *http.Request {
	return &http.Request{
		Header: http.Header{},
	}
}
