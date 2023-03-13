package main

import (
	"api-golang/src/config"
	"api-golang/src/database"
	"api-golang/src/router"
	"api-golang/src/utils"
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
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
