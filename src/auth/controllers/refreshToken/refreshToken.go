package refreshtoken

import (
	"api-golang/src/database"
	userrepository "api-golang/src/modules/user/repository"
	"api-golang/src/utils"
	responses "api-golang/src/utils"
	"net/http"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	requestToken := r.Header.Get("x-refresh-token")

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := userrepository.NewUserRepository(db)
	savedToken, erro := repository.FindRefreshToken(requestToken)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	erro = utils.ValidateToken(savedToken.Token)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
	}
}
