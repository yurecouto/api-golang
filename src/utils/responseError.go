package utils

import "net/http"

func ResponseError(response http.ResponseWriter, statusCode int, erro error) {
	ResponseJson(response, statusCode, struct {
		Erro string `json:"error"`
	}{
		Erro: erro.Error(),
	})
}
