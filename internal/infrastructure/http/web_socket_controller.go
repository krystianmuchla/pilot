package http

import (
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	mouseMoveInput        byte = 0
	mouseButton1Input     byte = 1
	scrollInput           byte = 2
	keyboardInput         byte = 3
	mouseButton1DownInput byte = 4
	mouseButton1UpInput   byte = 5
	mouseButton2DownInput byte = 6
	mouseButton2UpInput   byte = 7
	mouseButton3DownInput byte = 8
	mouseButton3UpInput   byte = 9
)

type WebSocketController struct {
	mouseAdapter    MouseAdapter
	keyboardAdapter KeyboardAdapter
}

func NewWebSocketController(mouseAdapter MouseAdapter, keyboardAdapter KeyboardAdapter) *WebSocketController {
	return &WebSocketController{mouseAdapter: mouseAdapter, keyboardAdapter: keyboardAdapter}
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
				err = controller.mouseAdapter.Move(x, y)
				if err != nil {
					println(err.Error())
					break
				}
			} else if payload[0] == mouseButton1Input {
				err = controller.mouseAdapter.ClickButton1()
				if err != nil {
					println(err.Error())
					break
				}
			} else if payload[0] == scrollInput {
				x := int(payload[1]) - 128
				y := int(payload[2]) - 128
				err = controller.mouseAdapter.Scroll(x, y)
				if err != nil {
					println(err.Error())
					break
				}
			} else if payload[0] == keyboardInput {
				keyCode := payload[1]
				err = controller.keyboardAdapter.EnterKey(keyCode)
				if err != nil {
					println(err.Error())
					break
				}
			} else if payload[0] == mouseButton1DownInput {
				err = controller.mouseAdapter.PressButton1()
				if err != nil {
					println(err.Error())
					break
				}
			} else if payload[0] == mouseButton1UpInput {
				err = controller.mouseAdapter.ReleaseButton1()
				if err != nil {
					println(err.Error())
					break
				}
			} else if payload[0] == mouseButton2DownInput {
				err = controller.mouseAdapter.PressButton2()
				if err != nil {
					println(err.Error())
					break
				}
			} else if payload[0] == mouseButton2UpInput {
				err = controller.mouseAdapter.ReleaseButton2()
				if err != nil {
					println(err.Error())
					break
				}
			} else if payload[0] == mouseButton3DownInput {
				err = controller.mouseAdapter.PressButton3()
				if err != nil {
					println(err.Error())
					break
				}
			} else if payload[0] == mouseButton3UpInput {
				err = controller.mouseAdapter.ReleaseButton3()
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
