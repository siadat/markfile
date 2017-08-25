package middleware

import "net/http"

type Router func(r *http.Request) http.Handler

func (h Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(r).ServeHTTP(w, r)
}
