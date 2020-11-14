package hello

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"
)

func (s *handler) checkAuthorization(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isAuthorized(r) {
			h(w, r)
		} else {
			s.respond(w, r, struct{}{}, http.StatusUnauthorized)
		}
	}
}

func isAuthorized(r *http.Request) bool {
	token := parseAuthorizationToken(r)
	return token != nil && token[0] == "admin" && token[1] == "password"
}

func parseAuthorizationToken(r *http.Request) []string {
	split := strings.Split(r.Header.Get("Authorization"), "Basic ")
	if len(split) != 2 {
		return nil
	}
	return decodeAuthorizationToken(split[1])
}

func decodeAuthorizationToken(encoded string) []string {
	resultBuffer := make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	length, err := base64.StdEncoding.Decode(resultBuffer, []byte(encoded))
	if err != nil {
		return nil
	}

	split := strings.Split(string(resultBuffer[:length]), ":")
	if len(split) != 2 {
		return nil
	}

	return split
}

func (s *handler) logAccess(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s request to %s\n", r.RequestURI, r.Method)
		h(w, r)
	}
}
