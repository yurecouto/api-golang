package models

import (
	"api-golang/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdat,omitempty"`
}

func (user *User) Prepare(step string) error {
	if erro := user.validate(step); erro != nil {
		return erro
	}

	if erro := user.format(step); erro != nil {
		return erro
	}

	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("The name is required.")
	}

	if user.Email == "" {
		return errors.New("The email is required.")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("The email is invalid.")
	}

	if user.Password == "" && step == "register" {
		return errors.New("The password is required.")
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {
		passwordHash, erro := security.Hash(user.Password)
		if erro != nil {
			return erro
		}

		user.Password = string(passwordHash)
	}

	return nil
}
