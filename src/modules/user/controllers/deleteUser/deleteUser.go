package deleteuser

import "net/http"

func Controller(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}
