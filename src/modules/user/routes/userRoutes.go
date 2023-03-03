package user

import (
	createuser "api-golang/src/modules/user/controllers/createUser"
	deleteuser "api-golang/src/modules/user/controllers/deleteUser"
	showallusers "api-golang/src/modules/user/controllers/showAllUsers"
	showuser "api-golang/src/modules/user/controllers/showUser"
	updateuser "api-golang/src/modules/user/controllers/updateUser"

	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Get("/", showallusers.Controller)
	r.Get("/{id}", showuser.Controller)
	r.Post("/", createuser.Controller)
	r.Patch("/{id}", updateuser.Controller)
	r.Delete("/{id}", deleteuser.Controller)
}
