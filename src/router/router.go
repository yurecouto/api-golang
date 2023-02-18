package router

import (
	"api-golang/src/router/routes"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()

	return routes.ConfigureRoutes(r)
}
