package hello

import (
	"encoding/json"
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var srv *Server
var request *http.Request
var responseWriter *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	srv = NewServer(nil)
	request = httptest.NewRequest("GET", "/health", nil)
	responseWriter = httptest.NewRecorder()
	srv.ServeHTTP(responseWriter, request)
}

func teardown() {

}

func TestHandleHealthResponseBody(t *testing.T) {
	is := is.New(t)
	is.Equal(responseWriter.Code, http.StatusOK)
}

func TestHandleHealthResponseCode(t *testing.T) {
	is := is.New(t)
	type response = struct {
		Status string `json:"status"`
	}
	expected := response{
		Status: "ready",
	}
	var actual response
	err := json.NewDecoder(responseWriter.Body).Decode(&actual)
	is.NoErr(err)
	is.Equal(actual, expected)
}
