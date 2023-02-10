package routes

import (
	"api-golang/src/middlewares"
	user "api-golang/src/modules/user/routes"

	"github.com/gorilla/mux"
)

func ConfigureRoutes(r *mux.Router) *mux.Router {
	routes := user.Routes
	// routes = append(routes, loginRoute)
	// routes = append(routes, postRoutes...)

	for _, route := range routes {
		if route.Auth {
			r.HandleFunc(route.Uri,
				middlewares.Logger(middlewares.Auth(route.Controller)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.Uri, route.Controller).Methods(route.Method)
		}
	}
	return r
}
