package interfaces

import "net/http"

type Route struct {
	Uri        string
	Method     string
	Controller func(http.ResponseWriter, *http.Request)
	Auth       bool
}
