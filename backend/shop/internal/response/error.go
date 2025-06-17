package response

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

func ParseHttpError(err error) (int, string, error) {
	if httpErr, ok := err.(*HttpError); ok {
		return httpErr.Code, httpErr.Message, nil
	}
	return 0, "", err
}

func IsHTTPError(err error) bool {
	_, ok := err.(*HttpError)
	return ok
}
