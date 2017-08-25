package middleware

import (
	"net/http"
	"os"
	"path/filepath"
)

type Rewriter struct {
	Root string
}

func (h *Rewriter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		_, err := os.Stat(filepath.Join(h.Root, "README.md"))
		if err == nil {
			r.URL.Path = "/README.md"
		}
	}
}
