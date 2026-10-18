package http

import (
	"net/http"
	"pilot/web"
	"strings"
)

type ResourceController struct{}

func NewResourceController() *ResourceController {
	return &ResourceController{}
}

func (controller ResourceController) Handle(responseWriter http.ResponseWriter, request *http.Request) error {
	path := strings.TrimPrefix(request.URL.Path, "/")
	http.ServeFileFS(responseWriter, request, web.Resources, path)
	return nil
}
