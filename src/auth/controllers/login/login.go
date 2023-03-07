package login

import (
	"api-golang/src/database"
	"api-golang/src/models"
	"api-golang/src/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"

	userrepository "api-golang/src/modules/user/repository"
	responses "api-golang/src/utils"
)

type Login struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginResponse struct {
	user  models.User
	token string
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
	defer db.Close()

	repository := userrepository.NewUserRepository(db)
	user, erro := repository.FindByEmail(login.Email)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	token, erro := utils.GenerateToken(user.ID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = utils.CheckPassword(login.Password, user.Password); erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, []byte(token))
}
