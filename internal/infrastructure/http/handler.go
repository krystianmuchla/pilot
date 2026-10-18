package http

import (
	"net/http"
	"strings"
)

type Handler struct {
	webSocketController *WebSocketController
	mainController      *MainController
	resourceController  *ResourceController
}

func NewHandler(
	webSocketController *WebSocketController,
	mainController *MainController,
	resourceController *ResourceController,
) *Handler {
	return &Handler{
		webSocketController: webSocketController,
		mainController:      mainController,
		resourceController:  resourceController,
	}
}

func (handler *Handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	var err error
	if isWebSocketHandshake(request) {
		err = handler.webSocketController.Handle(responseWriter, request)
	} else if strings.HasPrefix(request.URL.Path, "/resources") {
		err = handler.resourceController.Handle(responseWriter, request)
	} else {
		err = handler.mainController.Handle(responseWriter, request)
	}
	if err != nil {
		println(err.Error())
		responseWriter.WriteHeader(http.StatusInternalServerError)
	}
}

func isWebSocketHandshake(request *http.Request) bool {
	upgrade := request.Header.Get("Upgrade")
	connection := request.Header.Get("Connection")

	return upgrade == "websocket" && connection == "Upgrade"
}
