package auth

import (
	"api-golang/src/auth/controllers/login"
	refreshtoken "api-golang/src/auth/controllers/refreshToken"

	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	r.Post("/login", login.Controller)
	r.Post("/refresh/token", refreshtoken.Controller)
}
