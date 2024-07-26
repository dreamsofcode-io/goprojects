package middleware

import (
	"net/http"
)

func NoCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		r.Header.Add("Cache-Control", "max-age=0")
	})
}
