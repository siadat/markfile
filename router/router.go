package router

import (
	"net/http"
)

type Handler struct {
	GetHandler func(r *http.Request) http.Handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.GetHandler == nil {
		http.Error(w, "no handler", http.StatusInternalServerError)
		return
	}
	h.GetHandler(r).ServeHTTP(w, r)
}
