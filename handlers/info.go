package handlers

import (
	"fmt"
	"net/http"

	"github.com/Smart-Pot/pkg/common"
	"github.com/gorilla/mux"
)

func SetInfoHandler(r *mux.Router, serviceName string, version common.Version) {
	r.Methods("GET").Path("info").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s v-%s", serviceName, version.String())
	})
}
