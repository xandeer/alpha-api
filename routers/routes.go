package routers

import (
	"net/http"

	. "github.com/xandeer/alpha-api/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Item",
		"GET",
		"/",
		AllItems,
	},
	Route{
		"Item",
		"GET",
		"/items",
		AllItems,
	},
	Route{
		"FindItem",
		"GET",
		"/items/{id}",
		FindItem,
	},
	Route{
		"CreateItem",
		"POST",
		"/items",
		CreateItem,
	},
	Route{
		"UpdateItem",
		"PUT",
		"/items",
		UpdateItem,
	},
	Route{
		"DeleteItem",
		"DELETE",
		"/items/{id}",
		DeleteItem,
	},
	Route{
		"GetOperateVersion",
		"GET",
		"/version",
		OperateVersion,
	},
	Route{
		"PullOperates",
		"GET",
		"/operates/{version}",
		PullOperates,
	},
	Route{
		"Signin",
		"POST",
		"/signin",
		Signin,
	},
}
