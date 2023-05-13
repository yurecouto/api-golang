package main

import (
	"api-golang/src/config"
	"api-golang/src/database"
	"api-golang/src/router"
	"api-golang/src/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	config.Load()
	utils.Log("Golang API Starting...")

	r := chi.NewRouter()
	r.Route("/", router.Routes)

	_, erro := database.Connect()
	if erro != nil {
		utils.Error(erro)
		return
	}
	utils.Log("Golang API Conected to Database.")

	utils.Log("Golang API Listening and Serving!")
	http.ListenAndServe(config.ApiPort, r)
}
