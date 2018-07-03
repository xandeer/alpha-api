package routers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/xandeer/alpha-api/middleware"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	cors := &middleware.CORSMiddleware{}
	logger := &middleware.LoggerMiddleware{}
	auth := &middleware.AuthMiddleware{}

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		router.
			Methods(route.Method, "OPTIONS").
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	// router.Use(mux.CORSMethodMiddleware(router))
	router.Use(logger.Middleware)
	router.Use(cors.Middleware)
	router.Use(auth.Middleware)

	return router
}
