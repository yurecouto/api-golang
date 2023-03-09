package utils

import (
	"api-golang/src/config"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Metodo errado %v", token.Header["alg"])
	}

	return config.AccessKey, nil
}

func ValidateToken(tokenString string) error {
	token, erro := jwt.Parse(tokenString, returnVerificationKey)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Invalid Token")
}
