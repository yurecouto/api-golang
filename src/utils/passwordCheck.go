package utils

import "golang.org/x/crypto/bcrypt"

func PasswordCheck(passwordString, passwordHash string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(passwordHash),
		[]byte(passwordString),
	)
}
