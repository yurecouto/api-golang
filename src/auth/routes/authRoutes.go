package auth

import (
	"api-golang/src/auth/controllers/login"
	createuser "api-golang/src/modules/user/controllers/createUser"
	showuser "api-golang/src/modules/user/controllers/showUser"

	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	r.Post("/login", login.Controller)
	r.Post("/refresh/token", showuser.Controller)
	r.Post("/change/password", createuser.Controller)
}
