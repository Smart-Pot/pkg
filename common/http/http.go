package http

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// EnableCORS enable CORS for given router.
func EnableCORS(r *mux.Router) http.Handler {
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization","x-auth-token"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	return handlers.CORS(headers, methods, origins)(r)
}

// EnableCORSWithOption enable CORS for given router and given options.
func EnableCORSWithOption(r *mux.Router, opts handlers.CORSOption) http.Handler {
	return handlers.CORS(opts)(r)
}
