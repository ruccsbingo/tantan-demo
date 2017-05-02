package router

import (
	"net/http"
	"tantan-demo/controller"
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
		"ListUsers",
		"GET",
		"/users",
		controller.ListUsers,
	},
	Route{
		"CreateUsers",
		"POST",
		"/users",
		controller.CreateUsers,
	},
	Route{
		"ListRelations",
		"GET",
		"/users/{user_id}/relationships",
		controller.ListRelations,
	},
	Route{
		"UpdateRelations",
		"PUT",
		"/users/{user_id}/relationships/{other_user_id}",
		controller.UpdateRelations,
	},
}
