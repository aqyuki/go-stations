package handler

import (
	"fmt"
	"net/http"
)

var _ http.Handler = (*AuthHandler)(nil)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler { return &AuthHandler{} }

func (h *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is a secret message!\n")
}
