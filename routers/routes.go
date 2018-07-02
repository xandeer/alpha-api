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
		"Get",
		"/",
		AllItems,
	},
	Route{
		"Item",
		"Get",
		"/items",
		AllItems,
	},
	Route{
		"FindItem",
		"Get",
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
		"Put",
		"/items",
		UpdateItem,
	},
	Route{
		"DeleteItem",
		"Delete",
		"/items/{id}",
		DeleteItem,
	},
	Route{
		"GetOperateVersion",
		"Get",
		"/version",
		OperateVersion,
	},
	Route{
		"PullOperates",
		"Get",
		"/operates/{version}",
		PullOperates,
	},
	// Route{
	// 	"Signin",
	// 	"Post",
	// 	"/users",
	// 	Signin,
	// },
}
