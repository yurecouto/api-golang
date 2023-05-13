package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseJson(response http.ResponseWriter, statusCode int, data interface{}) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)

	if data != nil {
		if erro := json.NewEncoder(response).Encode(data); erro != nil {
			log.Fatal(erro)
		}
	}
}
