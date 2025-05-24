package http

import (
	"github.com/gorilla/websocket"
	"net/http"
	"pilot/internal/logic"
)

const (
	mouseMoveInput    byte = 0
	mouseButton1Input byte = 1
	scrollInput       byte = 4
	keyboardInput     byte = 5
)

type WebSocketController struct {
	mouseController    *logic.MouseController
	keyboardController *logic.KeyboardController
}

func NewWebSocketController(mouseController *logic.MouseController, keyboardController *logic.KeyboardController) *WebSocketController {
	return &WebSocketController{mouseController, keyboardController}
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
				err = controller.mouseController.Move(x, y)
				if err != nil {
					println(err.Error())
					break
				}
			} else if payload[0] == mouseButton1Input {
				err = controller.mouseController.ClickButton1()
				if err != nil {
					println(err.Error())
					break
				}
			} else if payload[0] == scrollInput {
				x := int(payload[1]) - 128
				y := int(payload[2]) - 128
				err = controller.mouseController.Scroll(x, y)
				if err != nil {
					println(err.Error())
					break
				}
			} else if payload[0] == keyboardInput {
				keyCode := payload[1]
				err = controller.keyboardController.EnterKey(keyCode)
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
	CheckOrigin: func(request *http.Request) bool {
		return true
	},
}
