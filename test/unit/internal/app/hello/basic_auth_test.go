package hello

import (
	"github.com/matryer/is"
	"github.com/rafaelzig/go-rest/internal/app/hello"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBasicAuthPass(t *testing.T) {
	is := is.New(t)
	srv := hello.NewServer(nil)
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	user := "user"
	pass := "pass"
	r.SetBasicAuth(user, pass)
	is.True(srv.BasicAuth(user, pass)(r))
}

func TestBasicAuthFail(t *testing.T) {
	is := is.New(t)
	srv := hello.NewServer(nil)
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.SetBasicAuth("user", "pass")
	is.True(!srv.BasicAuth("another", "password")(r))
}
