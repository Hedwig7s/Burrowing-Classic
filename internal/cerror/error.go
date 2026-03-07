package cerror

import "fmt"

type CodedError struct {
	Code    int
	Message string
}

func (error CodedError) Error() string {
	return error.Message
}

func NewError(code int, error string) CodedError {
	return CodedError{Code: code, Message: error}
}

func NewErrorf(code int, format string, args ...any) CodedError {
	return NewError(code, fmt.Sprintf(format, args...))
}
