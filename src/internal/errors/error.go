package errors

import "fmt"

type SrvErr interface {
	Code() int
	Message() string
}

type ServerError struct {
	code    int
	message string
}

func (e *ServerError) Code() int {
	return e.code
}

func (e *ServerError) Message() string {
	return e.message
}

func (e *ServerError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.code, e.message)
}

func NewServerError(code int, msg string) *ServerError {
	return &ServerError{
		code:    code,
		message: msg,
	}
}

const (
	COMMON_ERROR_CODE = 10000 + iota*100
	SESSION_ERROR_CODE
)
