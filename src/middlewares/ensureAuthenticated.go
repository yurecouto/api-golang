package middlewares

import (
	"api-golang/src/utils"
	responses "api-golang/src/utils"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
)

func EnsureAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("x-access-token")
		token, erro := jwt.Parse(tokenString, utils.ReturnVerificationKey)
		if erro != nil {
			responses.Erro(w, http.StatusUnauthorized, erro)
			return
		}

		if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
			if erro != nil {
				responses.Erro(w, http.StatusUnauthorized, erro)
				return
			}

			if erro := utils.ValidateToken(tokenString); erro != nil {
				responses.Erro(w, http.StatusUnauthorized, erro)
				return
			}

			ctx := context.WithValue(r.Context(), "userId", userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
