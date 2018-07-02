package routers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/xandeer/alpha-api/middleware"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	cors := &middleware.CORSOriginsMiddleware{}
	logger := &middleware.LoggerMiddleware{}

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(cors.Middleware)
	router.Use(logger.Middleware)

	return router
}
