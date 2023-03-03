package utils

import "golang.org/x/crypto/bcrypt"

func VerifyPassword(passwordString, passwordHash string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(passwordHash),
		[]byte(passwordString),
	)
}
