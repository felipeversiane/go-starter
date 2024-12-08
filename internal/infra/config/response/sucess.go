package response

import (
	"net/http"
)

type SuccessResponse struct {
	Message string      `json:"message,omitempty"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(message string, code int, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Message: message,
		Code:    code,
		Data:    data,
	}
}

func NewCreatedResponse(message string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Message: message,
		Code:    http.StatusCreated,
		Data:    data,
	}
}

func NewNoContentResponse(message string) *SuccessResponse {
	return &SuccessResponse{
		Message: message,
		Code:    http.StatusNoContent,
	}
}

func NewOKResponse(message string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Message: message,
		Code:    http.StatusOK,
		Data:    data,
	}
}
