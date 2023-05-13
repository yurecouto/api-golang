package showallusers

import (
	"api-golang/src/database"
	userrepository "api-golang/src/modules/user/repository"
	responses "api-golang/src/utils"

	"net/http"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repository := userrepository.NewUserRepository(db)
	users, erro := repository.FindAll()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}
