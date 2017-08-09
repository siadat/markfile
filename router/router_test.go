package router_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/siadat/markfile/router"
)

type Handler struct {
	F func(w http.ResponseWriter, r *http.Request)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.F(w, r)
}

func newHandler(response string) *Handler {
	return &Handler{
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(response))
		},
	}
}

func check(url, expected string) error {
	var err error
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	buf, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return err
	}
	if string(buf) != expected {
		return fmt.Errorf("expected %q got %q", expected, string(buf))
	}
	return nil
}

func TestRouter(t *testing.T) {
	h := &router.Handler{
		func(r *http.Request) http.Handler {
			if r.URL.Path == "/1" {
				return newHandler("expected1")
			}
			return newHandler("expected2")
		},
	}

	ts := httptest.NewServer(h)
	defer ts.Close()

	if err := check(ts.URL+"/1", "expected1"); err != nil {
		t.Error(err)
	}

	if err := check(ts.URL+"/2", "expected2"); err != nil {
		t.Error(err)
	}
}
