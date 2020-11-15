package hello

import (
	"encoding/json"
	"github.com/matryer/is"
	"github.com/rafaelzig/go-rest/internal/app/hello"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleShutdownResponse(t *testing.T) {
	is := is.New(t)
	srv := hello.NewServer(nil)
	r := httptest.NewRequest(http.MethodDelete, "/shutdown", nil)
	r.SetBasicAuth("admin", "password")
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, r)
	is.Equal(w.Code, http.StatusAccepted)
	is.Equal(w.Header().Get("Content-Type"), "application/json")
	type response = struct {
		Status string `json:"status"`
	}
	expected := response{
		Status: "shutdown initiated",
	}
	var actual response
	err := json.NewDecoder(w.Body).Decode(&actual)
	is.NoErr(err)
	is.Equal(actual, expected)
}
