package showuser

import (
	"api-golang/src/database"
	userrepository "api-golang/src/modules/user/repository"
	"api-golang/src/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Controller(w http.ResponseWriter, r *http.Request) {
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

	user, erro := repo.FindById(userID)
	if erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	utils.ResponseJson(w, http.StatusOK, user)
}
