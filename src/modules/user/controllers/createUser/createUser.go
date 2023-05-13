package createuser

import (
	"api-golang/src/database"
	"api-golang/src/models"
	userrepository "api-golang/src/modules/user/repository"
	"api-golang/src/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	// userId := r.Context().Value("userId")
	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		utils.ResponseError(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user *models.User
	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		utils.ResponseError(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare("register"); erro != nil {
		utils.ResponseError(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	repository := userrepository.NewUserRepository(db)
	erro = repository.Create(user)
	if erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	utils.ResponseJson(w, http.StatusCreated, user)
}
