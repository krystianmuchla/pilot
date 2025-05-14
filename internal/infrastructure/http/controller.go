package http

import "net/http"

type Controller interface {
	Handle(responseWriter http.ResponseWriter, request *http.Request) error
}
