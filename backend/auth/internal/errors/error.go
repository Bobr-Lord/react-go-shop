package errors

import "net/http"

type HTTPError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e HTTPError) Error() string {
	return e.Message
}

func NewHTTPError(status int, message string) *HTTPError {
	return &HTTPError{
		Message: message,
		Code:    status,
	}
}

func ParseHTTPError(err error) *HTTPError {
	res, ok := err.(*HTTPError)
	if !ok {
		return NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return res
}
