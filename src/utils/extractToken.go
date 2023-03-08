package utils

import (
	"net/http"
	"strings"
)

func ExtractToken(r *http.Request) string {
	token := r.Header.Get("x-refresh-token")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}
