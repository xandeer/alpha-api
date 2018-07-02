package middleware

import (
	"net/http"
)

type CORSOriginsMiddleware struct {
}

type origins []string

var Origins = origins{
	"http://xandeer.top",
	"http://localhost:1234",
}

func (mw *CORSOriginsMiddleware) Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if Origins.contains(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token,X-Custom-Header")
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
