package response

type SuccessResponse struct {
	Message *string     `json:"message,omitempty"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(code int, data interface{}, message ...string) *SuccessResponse {
	var msg *string

	if len(message) > 0 {
		msg = &message[0]
	}

	return &SuccessResponse{
		Message: msg,
		Code:    code,
		Data:    data,
	}
}
