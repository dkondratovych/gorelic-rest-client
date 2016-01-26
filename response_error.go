package newrelic

import (
	"net/http"
)

type ResponseError struct {
	Response     *http.Response
	ErrorCode    int
	ErrorMessage string
}

func (e ResponseError) errorCode() int {
	return e.ErrorCode
}

func (e *ResponseError) Error() string {
	return e.ErrorMessage
}
