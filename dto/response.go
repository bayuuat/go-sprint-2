package dto

type SuccessResponse[T any] struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Data T `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Message: message,
	}
}

func NewSuccessCreateResponse[T any](message string, data T) SuccessResponse[T] {
	return SuccessResponse[T]{
		Code: "201",
		Message: message,
		Data: data,
	}
}