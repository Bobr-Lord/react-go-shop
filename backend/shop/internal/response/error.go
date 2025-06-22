package response

import "net/http"

type HttpError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e *HttpError) Error() string {
	return e.Message
}

func NewCustomError(message string, code int) *HttpError {
	return &HttpError{
		Message: message,
		Code:    code,
	}
}

func ParseHttpError(err error) *HttpError {
	if httpErr, ok := err.(*HttpError); ok {
		return httpErr
	}
	return &HttpError{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	}
}

func IsHTTPError(err error) bool {
	_, ok := err.(*HttpError)
	return ok
}
