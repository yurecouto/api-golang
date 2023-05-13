package utils

import (
	"api-golang/src/config"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 1).Unix()
	permissions["userId"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(config.AccessKey))
}

func GenerateRefeshToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 48).Unix()
	permissions["userId"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(config.RefreshKey))
}

func ReturnVerificationKey(token *jwt.Token, isRefresh bool) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Wrong method %v", token.Header["alg"])
	}

	if isRefresh == true {
		return config.RefreshKey, nil
	}

	return config.AccessKey, nil
}

func ValidateToken(tokenString string, isRefresh bool) error {
	token, erro := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return ReturnVerificationKey(token, isRefresh)
	})
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Invalid Token")
}
