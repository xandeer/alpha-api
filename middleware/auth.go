package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"

	"github.com/xandeer/alpha-api/auth"
)

type AuthMiddleware struct {
}

type methods []string

var Methods = methods{
	"GET",
	"OPTIONS",
}

func (mw *AuthMiddleware) Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if Methods.contains(r.Method) {
			h.ServeHTTP(w, r)
			return
		}

		token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
			// since we only use the one private key to sign the tokens,
			// we also only use its public counter part to verify
			return auth.VerifyKey, nil
		})

		if err != nil {
			if r.RequestURI == "/signin" {
				h.ServeHTTP(w, r)
				return
			}
			response, _ := json.Marshal(map[string]string{"error": "Invalid token"})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(response)
			return
		}
		if r.RequestURI == "/signin" {
			response, _ := json.Marshal(token)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(response)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func (arr methods) contains(str string) bool {
	for _, it := range arr {
		if it == str {
			return true
		}
	}
	return false
}
