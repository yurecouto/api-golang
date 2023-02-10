package main

import (
	"api-golang/src/config"
	"api-golang/src/router"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	config.Load()
	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
