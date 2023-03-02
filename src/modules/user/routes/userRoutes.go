package user

import (
	createuser "api-golang/src/modules/user/controllers/createUser"

	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Get("/", createuser.Controller)
	r.Get("/{id}", createuser.Controller)
	r.Post("/", createuser.Controller)
	r.Put("/{id}", createuser.Controller)
	r.Patch("/{id}", createuser.Controller)
	r.Delete("/{id}", createuser.Controller)
}
