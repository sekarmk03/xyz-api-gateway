package utils

type ErrorResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewErrorResponse(code int, status, message string) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Status: status,
		Message: message,
	}
}
