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

	repository := userrepository.NewUserRepository(db)
	savedToken, erro := repository.FindRefreshToken(requestToken)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	var userID uint64 = uint64(savedToken.ID)

	erro = utils.ValidateToken(savedToken.Token)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
	}

	accessToken, erro := utils.GenerateToken(userID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	refreshToken, erro := utils.GenerateRefeshToken(userID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	type Response struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	}

	response := &Response{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	erro = repository.SaveRefreshToken(userID, refreshToken)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	erro = repository.DeleteRefreshToken(requestToken)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, response)
}
