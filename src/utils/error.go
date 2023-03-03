package utils

import (
	"net/http"
)

func Erro(response http.ResponseWriter, statusCode int, erro error) {
	JSON(response, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
