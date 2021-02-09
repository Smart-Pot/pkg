package handlers

import (
	"net/http"

	"github.com/Smart-Pot/pkg/common"
)

func GetInfoHandler(serviceName string, version common.Version) http.Handler {
	var f http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		// Write Information
	}

	return f
}
