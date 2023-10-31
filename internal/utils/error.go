package utils

import "net/http"

type Error interface {
	GetMessage() string
	GetStatusCode() int
	GetErrorType() string
}

type ErrorResponse struct {
	Message string `json:"message"`
	StatusCode int `json:"status_code"`
	ErrorType string `json:"error_type"`
}

func (e *ErrorResponse) GetMessage() string {
	return e.Message
}

func (e *ErrorResponse) GetStatusCode() int {
	return e.StatusCode
}

func (e *ErrorResponse) GetErrorType() string {
	return e.ErrorType
}

func NewError(message string, statusCode int, errorType string) Error {
	return &ErrorResponse{
		Message: message,
		StatusCode: statusCode,
		ErrorType: errorType,
	}
}

func NotFound(message string) Error {
	return NewError(message, http.StatusNotFound, "Not Found")
}

func BadRequest(message string) Error{
	return NewError(message, http.StatusBadRequest, "Bad Request")
}

func InternalServerError(message string) Error {
	return NewError(message, http.StatusInternalServerError, "Server Error")
}

func UnprocessibleEntity(message string) Error {
	return NewError(message, http.StatusUnprocessableEntity, "Invalid Request")
}