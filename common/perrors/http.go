package perrors

import (
	"context"
	"encoding/json"
	"net/http"
)

// EncodeHTTPError :
func EncodeHTTPError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	e := json.NewEncoder(w)
	// Parse error to perrors.Error
	perr, ok := err.(*Error)
	code := 400
	// if it's not parsed
	if ok {
		code = perr.code
	}
	w.WriteHeader(code)
	e.Encode(map[string]interface{}{
		"error" : err.Error(),
		"code" : code,
	})
}