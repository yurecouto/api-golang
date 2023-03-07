package utils

import "golang.org/x/crypto/bcrypt"

func CheckPassword(passwordString, passwordHash string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(passwordHash),
		[]byte(passwordString),
	)
}
