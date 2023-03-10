package showuser

import (
	"api-golang/src/database"
	userrepository "api-golang/src/modules/user/repository"
	responses "api-golang/src/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Controller(w http.ResponseWriter, r *http.Request) {
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

	user, erro := repo.FindById(userID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}
