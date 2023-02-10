package user

import (
	"api-golang/src/interfaces"
	createuser "api-golang/src/modules/user/controllers/createUser"
	deleteuser "api-golang/src/modules/user/controllers/deleteUser"
	showuser "api-golang/src/modules/user/controllers/showUser"
	showusers "api-golang/src/modules/user/controllers/showUsers"
	updateuser "api-golang/src/modules/user/controllers/updateUser"
	"net/http"
)

var Routes = []interfaces.Route{
	{
		Uri:        "/user",
		Method:     http.MethodPost,
		Controller: createuser.Controller,
		Auth:       false,
	},
	{
		Uri:        "/user",
		Method:     http.MethodGet,
		Controller: showusers.Controller,
		Auth:       false,
	},
	{
		Uri:        "/user/{userId}",
		Method:     http.MethodGet,
		Controller: showuser.Controller,
		Auth:       false,
	},
	{
		Uri:        "/user/{userId}",
		Method:     http.MethodPatch,
		Controller: updateuser.Controller,
		Auth:       false,
	},
	{
		Uri:        "/user/{userId}",
		Method:     http.MethodDelete,
		Controller: deleteuser.Controller,
		Auth:       false,
	},
}
