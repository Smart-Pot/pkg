package perrors

import "fmt"

// Error :
type Error interface {
	Error() string
	Code() int
}
type perror struct {
	code int
	message string
}

func (e *perror) Error() string {
	return e.message
}

func (e *perror) Code() int {
	return e.code
}

func New(msg string,code int) *perror {
	return &perror{
		message: msg,
		code: code,
	}
}

func FromError(msg string,code int,cause error) Error {
	return New(fmt.Sprintf("%s : %s",msg,cause.Error()),code)
}

var _ Error = (*perror)(nil)
var _ error = (Error)(nil)