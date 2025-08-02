package http

import "net/http"

type Handler struct {
	webSocketController *WebSocketController
	mainController      *MainController
}

func NewHandler(webSocketController *WebSocketController, mainController *MainController) *Handler {
	return &Handler{webSocketController: webSocketController, mainController: mainController}
}

func (handler *Handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	var err error
	if isWebSocketHandshake(request) {
		err = handler.webSocketController.Handle(responseWriter, request)
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
