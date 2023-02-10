package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(response http.ResponseWriter, statusCode int, data interface{}) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)

	if data != nil {
		if erro := json.NewEncoder(response).Encode(data); erro != nil {
			log.Fatal(erro)
		}
	}
}

func Erro(response http.ResponseWriter, statusCode int, erro error) {
	JSON(response, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
