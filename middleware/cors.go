package middleware

import (
	"net/http"
	"strings"
)

type CORSMiddleware struct {
}

type origins []string

var Origins = origins{
	"http://xandeer.top",
	"http://localhost:1234",
}
var allowedMethods = []string{
	"GET",
	"PUT",
	"POST",
	"DELETE",
	"OPTIONS",
}

func (mw *CORSMiddleware) Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if Origins.contains(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Access-Control-Allow-Methods", strings.Join(allowedMethods, ","))
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func (arr origins) contains(str string) bool {
	for _, it := range arr {
		if it == str {
			return true
		}
	}
	return false
}
