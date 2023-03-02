package main

import (
	"api-golang/src/config"
	"context"
	"fmt"
	"net/http"

	user "api-golang/src/modules/user/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("Golang API Starting.")
	config.Load()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(MeuMiddleware)
	r.Route("/user", user.Router)

	fmt.Println("Golang API Listening and Serving.")
	http.ListenAndServe(config.Port, r)
}

func MeuMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user", "123")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
