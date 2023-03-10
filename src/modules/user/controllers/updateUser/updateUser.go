package updateuser

import (
	"api-golang/src/database"
	"api-golang/src/models"
	userrepository "api-golang/src/modules/user/repository"
	responses "api-golang/src/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	Id := chi.URLParam(r, "id")

	userID, erro := strconv.ParseUint(Id, 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := userrepository.NewUserRepository(db)

	updatedUder, erro := repo.FindByIdAndUpdate(userID, user)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, updatedUder)
}
