package hello

import (
	"encoding/json"
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleIndexResponseCode(t *testing.T) {
	is := is.New(t)
	srv := Server{}
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	srv.handleIndex()(w, r)
	is.Equal(w.Code, http.StatusOK)
	is.Equal(w.Header().Get("Content-Type"), "application/json")
	type response = struct {
		Message string `json:"message"`
	}
	expected := response{
		Message: "Hello World",
	}
	var actual response
	err := json.NewDecoder(w.Body).Decode(&actual)
	is.NoErr(err)
	is.Equal(actual, expected)
}
