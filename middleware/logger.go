package middleware

import (
	"net/http"
	"log"
	"time"
)

type LoggerMiddleware struct {
}

func (mw *LoggerMiddleware) Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
		h.ServeHTTP(w, r)
	})
}
