package handler

import "net/http"

var _ http.Handler = (*DoPanicHandler)(nil)

type DoPanicHandler struct{}

func (h *DoPanicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	panic("panic test")
}
