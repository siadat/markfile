package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/siadat/markfile/flags"
	"github.com/siadat/markfile/logger"
	"github.com/siadat/markfile/markdown"
	"github.com/siadat/markfile/router"
)

func main() {
	opts := flags.Parse()
	fmt.Print(opts)

	fileHandler := http.FileServer(http.Dir(opts.Root()))
	markHandler := &markdown.Handler{http.Dir(opts.Root())}

	err := http.ListenAndServe(opts.ListenAddr(), &logger.Logger{
		&router.Handler{
			func(r *http.Request) http.Handler {
				if strings.HasSuffix(r.URL.Path, ".md") {
					return markHandler
				}
				return fileHandler
			},
		},
	},
	)
	if err != nil {
		log.Fatal(err)
	}
}
