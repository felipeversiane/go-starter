package response

import "net/http"

type ErrorResponse struct {
	Message string  `json:"message"`
	Err     string  `json:"error"`
	Code    int     `json:"code"`
	Causes  []Cause `json:"causes,omitempty"`
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *ErrorResponse) Error() string {
	return r.Message
}

func NewBadRequestError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewUnauthorizedError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func NewInternalServerError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}

func NewBadRequestValidationError(message string, causes []Cause) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewUnauthorizedRequestError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}
