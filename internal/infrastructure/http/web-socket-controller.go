package http

import (
	"github.com/gorilla/websocket"
	"net/http"
	"pilot/internal/logic"
)

const (
	mouseMoveInput byte = 0
)

type WebSocketController struct {
	mouseController *logic.MouseController
}

func NewWebSocketController(mouseController *logic.MouseController) *WebSocketController {
	return &WebSocketController{mouseController: mouseController}
}

func (controller WebSocketController) Handle(responseWriter http.ResponseWriter, request *http.Request) error {
	connection, err := webSocketUpgrader.Upgrade(responseWriter, request, nil)
	if err != nil {
		return err
	}
	for {
		messageType, payload, err := connection.ReadMessage()
		if err != nil {
			println(err.Error())
			break
		}
		if messageType == websocket.BinaryMessage {
			if payload[0] == mouseMoveInput {
				x := int(payload[1]) - 128
				y := int(payload[2]) - 128
				err := controller.mouseController.Move(x, y)
				if err != nil {
					println(err.Error())
					break
				}
			}
		}
	}
	err = connection.Close()
	if err != nil {
		return err
	}
	return nil
}

var webSocketUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
