package hello

import (
	"github.com/matryer/is"
	"github.com/rafaelzig/go-rest/internal/app/hello"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutesUnauthorized(t *testing.T) {
	is := is.New(t)
	srv := hello.NewServer(nil)
	routes := map[string]string{
		"/shutdown": http.MethodDelete,
	}
	for route, method := range routes {
		r := httptest.NewRequest(method, route, nil)
		w := httptest.NewRecorder()
		srv.ServeHTTP(w, r)
		is.Equal(w.Code, http.StatusUnauthorized)
		r.SetBasicAuth("invalid", "invalid")
		srv.ServeHTTP(w, r)
		is.Equal(w.Code, http.StatusUnauthorized)
	}
}

func TestRoutesAllowedMethods(t *testing.T) {
	is := is.New(t)
	srv := hello.NewServer(nil)

	routes := map[string]string{
		"/":         http.MethodGet,
		"/health":   http.MethodGet,
		"/shutdown": http.MethodDelete,
	}

	for route, method := range routes {
		r := httptest.NewRequest(method, route, nil)
		w := httptest.NewRecorder()
		srv.ServeHTTP(w, r)
		is.True(w.Code != http.StatusNotFound)
		is.True(w.Code != http.StatusMethodNotAllowed)
	}
}

func TestRoutesDisallowedMethods(t *testing.T) {
	is := is.New(t)
	srv := hello.NewServer(nil)
	routes := map[string][8]string{
		"/": {
			http.MethodHead,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodConnect,
			http.MethodOptions,
			http.MethodTrace,
			http.MethodDelete,
		},
		"/health": {
			http.MethodHead,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodConnect,
			http.MethodOptions,
			http.MethodTrace,
			http.MethodDelete,
		},
		"/shutdown": {
			http.MethodHead,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodConnect,
			http.MethodOptions,
			http.MethodTrace,
			http.MethodGet,
		},
	}
	for route, disallowedMethods := range routes {
		for _, disallowedMethod := range disallowedMethods {
			r := httptest.NewRequest(disallowedMethod, route, nil)
			w := httptest.NewRecorder()
			srv.ServeHTTP(w, r)
			is.Equal(w.Code, http.StatusMethodNotAllowed)
		}
	}
}
