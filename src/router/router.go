package router

import (
	auth "api-golang/src/auth/routes"
	user "api-golang/src/modules/user/routes"

	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	r.Route("/auth", auth.Routes)
	r.Route("/user", user.Routes)
}
