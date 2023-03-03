package createuser

import (
	"api-golang/src/models"
	responses "api-golang/src/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

	fmt.Println("TESTE =>", user)
}
