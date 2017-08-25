package middleware

import (
	"net/http"
)

type Decorator func(http.Handler) http.Handler

func PrefixDecorator(h http.Handler) Decorator {
	return func(innerHandler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
			innerHandler.ServeHTTP(w, r)
		})
	}
}

func Combine(core http.Handler, decorators ...Decorator) http.Handler {
	all := core
	for _, d := range decorators {
		all = d(all)
	}
	return all
}
