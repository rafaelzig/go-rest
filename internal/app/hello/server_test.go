package hello

import (
	"encoding/json"
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerIntegration(t *testing.T) {
	is := is.New(t)
	srv := httptest.NewServer(NewServer(nil))
	defer srv.Close()
	resp, err := http.Get(srv.URL + "/health")
	is.NoErr(err)
	is.Equal(resp.StatusCode, http.StatusOK)
	is.Equal(resp.Header.Get("Content-Type"), "application/json")
	type response = struct {
		Status string `json:"status"`
	}
	expected := response{
		Status: "ready",
	}
	var actual response
	err = json.NewDecoder(resp.Body).Decode(&actual)
	is.NoErr(err)
	is.Equal(actual, expected)
}
