package showallusers

import (
	"net/http"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	// db, erro := database.Connect()
	// if erro != nil {
	// 	responses.Erro(w, http.StatusInternalServerError, erro)
	// 	return
	// }
	// defer db.Close()

	// repository := userrepository.NewUserRepository(db)
	// users, erro := repository.FindAllUsers()
	// if erro != nil {
	// 	responses.Erro(w, http.StatusInternalServerError, erro)
	// 	return
	// }

	// responses.JSON(w, http.StatusOK, users)
}
