package middleware

import (
	"log"
	"net/http"
)

func LoggerMiddleware(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
}
