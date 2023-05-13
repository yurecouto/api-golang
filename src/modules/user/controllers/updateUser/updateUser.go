package updateuser

import (
	"api-golang/src/database"
	"api-golang/src/models"
	userrepository "api-golang/src/modules/user/repository"
	"api-golang/src/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Controller(w http.ResponseWriter, r *http.Request) {
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

	Id := chi.URLParam(r, "id")

	userID, erro := strconv.ParseUint(Id, 10, 64)
	if erro != nil {
		utils.ResponseError(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	repo := userrepository.NewUserRepository(db)

	updatedUser, erro := repo.Update(userID, user)
	if erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	utils.ResponseJson(w, http.StatusOK, updatedUser)
}
