package main

import (
	"api-golang/src/config"
	"api-golang/src/router"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	config.Load()
	r := router.Generate()

	log.Fatal(http.ListenAndServe(":5000", r))
}
