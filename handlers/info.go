package handlers

import (
	"fmt"
	"net/http"

	"github.com/Smart-Pot/pkg/common"
)


var _ http.Handler = (*infoHandler)(nil) 

type infoHandler struct {
	serviceName string
	version common.Version
}


func (h *infoHandler) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintln(w,h.info())
}

func (h *infoHandler) info() string {
	return fmt.Sprintf("%s v-%s", h.serviceName, h.version.String())
}



func NewInfoHandler(serviceName string, version common.Version) http.Handler {
	return &infoHandler{
		serviceName: serviceName,
		version : version,
	}
}