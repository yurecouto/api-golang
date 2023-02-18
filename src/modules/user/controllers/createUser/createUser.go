package createuser

import (
	"api-golang/src/database"
	"api-golang/src/models"
	"api-golang/src/modules/user/repository"
	"api-golang/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Controller(response http.ResponseWriter, request *http.Request) {
	requestBody, erro := ioutil.ReadAll(request.Body)
	if erro != nil {
		responses.Erro(response, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		responses.Erro(response, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare("cadastro"); erro != nil {
		responses.Erro(response, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NewUserRepository(db)
	user.ID, erro = repository.Create(user)
	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(response, http.StatusCreated, user)
}
