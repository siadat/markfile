package logger

import (
	"log"
	"net/http"
)

type Logger struct {
	Next http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	if l.Next == nil {
		return
	}
	l.Next.ServeHTTP(w, r)
}
