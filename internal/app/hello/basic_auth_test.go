package hello

import (
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBasicAuthPass(t *testing.T) {
	is := is.New(t)
	srv := Server{}
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	user := "user"
	pass := "pass"
	r.SetBasicAuth(user, pass)
	is.True(srv.basicAuth(user, pass)(r))
}

func TestBasicAuthFail(t *testing.T) {
	is := is.New(t)
	srv := Server{}
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.SetBasicAuth("user", "pass")
	is.True(!srv.basicAuth("another", "password")(r))
}
