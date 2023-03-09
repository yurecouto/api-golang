package auth

import (
	"api-golang/src/auth/controllers/login"
	refreshtoken "api-golang/src/auth/controllers/refreshToken"
	createuser "api-golang/src/modules/user/controllers/createUser"

	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	r.Post("/login", login.Controller)
	r.Post("/refresh/token", refreshtoken.Controller)
	r.Post("/change/password", createuser.Controller)
}
