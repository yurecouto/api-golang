package models

import (
	"api-golang/src/utils"
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
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type UserToken struct {
	ID     uint64 `json:"id,omitempty"`
	Token  string `json:"token,omitempty"`
	UserId uint64 `json:"userId,omitempty"`
}

func (user *User) Prepare(stage string) error {
	if erro := user.validate(stage); erro != nil {
		return erro
	}

	if erro := user.format(stage); erro != nil {
		return erro
	}

	return nil
}

func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New("Name is required.")
	}

	if user.Email == "" {
		return errors.New("E-mail is required.")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("A valid E-mail is required.")
	}

	if user.Password == "" && stage == "register" {
		return errors.New("Password is required.")
	}

	return nil
}

func (user *User) format(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "register" {
		passwordHash, erro := utils.HashPassword(user.Password)
		if erro != nil {
			return erro
		}

		user.Password = string(passwordHash)
	}

	return nil
}
