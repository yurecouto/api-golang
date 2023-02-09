package user

import (
	"net/http"
)

type Route struct {
	Uri        string
	Method     string
	Controller func(http.ResponseWriter, *http.Request)
	Auth       bool
}

var userRoutes = []Route{
	{
		Uri:        "/user",
		Method:     http.MethodPost,
		Controller: CreateUserController,
		Auth:       false,
	},
	{
		Uri:        "/user",
		Method:     http.MethodGet,
		Controller: func(w http.ResponseWriter, r *http.Request) {},
		Auth:       false,
	},
	{
		Uri:        "/user/{userId}",
		Method:     http.MethodGet,
		Controller: func(w http.ResponseWriter, r *http.Request) {},
		Auth:       false,
	},
	{
		Uri:        "/user/{userId}",
		Method:     http.MethodPatch,
		Controller: func(w http.ResponseWriter, r *http.Request) {},
		Auth:       false,
	},
	{
		Uri:        "/user/{userId}",
		Method:     http.MethodDelete,
		Controller: func(w http.ResponseWriter, r *http.Request) {},
		Auth:       false,
	},
}
