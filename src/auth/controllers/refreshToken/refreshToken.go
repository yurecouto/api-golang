package refreshtoken

import (
	"api-golang/src/utils"
	"fmt"
	"net/http"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	requestToken := utils.ExtractToken(r)

	fmt.Println(requestToken)
}
