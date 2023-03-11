package main

import (
	"api-golang/src/config"
	"api-golang/src/database"
	"api-golang/src/router"
	"api-golang/src/utils"
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	//"github.com/go-chi/chi/v5/middleware"
)

func main() {
	utils.Log("Golang API Starting...")

	config.Load()
	r := chi.NewRouter()

	r.Route("/", router.Routes)

	db, erro := database.Connect()
	if erro != nil {
		utils.Error(erro)
		return
	}
	defer db.Close()
	utils.Log("Golang API Conected to Database.")

	utils.Log("Golang API Listening and Serving!")
	http.ListenAndServe(config.Port, r)
}

func MeuMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user", "123")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
