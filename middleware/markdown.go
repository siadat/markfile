package middleware

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/russross/blackfriday"
)

type Markdown struct {
	Root http.Dir
}

func (m *Markdown) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, err := m.Root.Open(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = renderMarkdown(f, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderMarkdown(r io.Reader, w io.Writer) error {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	buf = blackfriday.MarkdownCommon(buf)
	w.Write([]byte(`<!DOCTYPE html> <html> <head> <meta http-equiv="Content-Type" content="text/html; charset=utf-8"> </head>`))
	w.Write(buf)
	w.Write([]byte(`</body> </html>`))
	return nil
}
