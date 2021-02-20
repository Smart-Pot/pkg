package perrors

import "net/http"

var (
	// ErrInternalServer :
	ErrInternalServer = FromStatusCode(http.StatusInternalServerError)
)