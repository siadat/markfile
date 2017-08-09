package markdown

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/russross/blackfriday"
)

type Handler struct {
	Root http.Dir
}

func (m *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, err := m.Root.Open(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = convert(f, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func convert(r io.Reader, w io.Writer) error {
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
