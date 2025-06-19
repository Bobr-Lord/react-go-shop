package errors

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
