package login

import (
	"api-golang/src/database"
	userrepository "api-golang/src/modules/user/repository"
	"api-golang/src/utils"
	responses "api-golang/src/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Login struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func Controller(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var login Login
	if erro = json.Unmarshal(requestBody, &login); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repository := userrepository.NewUserRepository(db)
	user, erro := repository.FindByEmail(login.Email)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = utils.CheckPassword(login.Password, user.Password); erro != nil {
		responses.Erro(w, http.StatusUnauthorized, fmt.Errorf("Invalid Password."))
		return
	}

	accessToken, erro := utils.GenerateToken(user.ID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	refreshToken, erro := utils.GenerateRefeshToken(user.ID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	type Response struct {
		Name         string `json:"name"`
		Email        string `json:"email"`
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	}

	response := &Response{
		Name:         user.Name,
		Email:        user.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	erro = repository.SaveRefreshToken(user.ID, refreshToken)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}
