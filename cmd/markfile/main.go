package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/siadat/markfile/middleware"
)

func main() {
	opts := parseOpts()
	fmt.Print(opts)

	fileHandler := http.FileServer(http.Dir(opts.Root()))
	markHandler := &middleware.Markdown{http.Dir(opts.Root())}
	urlRewriter := &middleware.Rewriter{Root: opts.Root()}
	logMiddleware := http.HandlerFunc(middleware.LoggerMiddleware)
	router := middleware.Router(func(r *http.Request) http.Handler {
		if strings.HasSuffix(r.URL.Path, ".md") {
			return markHandler
		}

		return fileHandler
	})

	handler := middleware.Combine(
		router,
		middleware.PrefixDecorator(urlRewriter),
		middleware.PrefixDecorator(logMiddleware),
	)

	err := http.ListenAndServe(opts.ListenAddr(), handler)
	if err != nil {
		log.Fatal(err)
	}
}
