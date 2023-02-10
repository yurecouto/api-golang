package middlewares

import (
	"api-golang/src/auth"
	"api-golang/src/responses"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		log.Printf("\n %s %s %s", request.Method, request.RequestURI, request.Host)
		next(response, request)
	}
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if erro := auth.ValidateToken(request); erro != nil {
			responses.Erro(response, http.StatusUnauthorized, erro)
			return
		}
		next(response, request)
	}
}
