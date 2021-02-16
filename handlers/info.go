package handlers

import (
	"fmt"
	"net/http"

	"github.com/Smart-Pot/pkg/common/version"
)


var _ http.Handler = (*infoHandler)(nil) 

type infoHandler struct {
	serviceName string
	version version.Version
}


func (h *infoHandler) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintln(w,h.info())
}

func (h *infoHandler) info() string {
	return fmt.Sprintf("%s v-%s", h.serviceName, h.version.String())
}



func NewInfoHandler(serviceName string, version version.Version) http.Handler {
	return &infoHandler{
		serviceName: serviceName,
		version : version,
	}
}