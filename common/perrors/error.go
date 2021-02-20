package perrors

// Error :
type Error struct {
	code int
	message string
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) Code() int {
	return e.code
}

func New(msg string,code int) *Error {
	return &Error{
		message: msg,
		code: code,
	}
}



var _ error = (*Error)(nil)