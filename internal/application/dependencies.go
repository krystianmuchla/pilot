package application

import (
	"pilot/internal/infrastructure/http"
	"pilot/internal/infrastructure/linux"
)

func ResolveDependencies() (*linux.Mouse, *linux.Keyboard, *http.Handler, error) {
	mouse, err := linux.NewMouse("pilot-mouse")
	if err != nil {
		return nil, nil, nil, err
	}
	keyboard, err := linux.NewKeyboard("pilot-keyboard")
	if err != nil {
		return mouse, nil, nil, err
	}
	mouseAdapter := http.NewMouseAdapter(mouse)
	keyboardAdapter := http.NewKeyboardAdapter(keyboard)
	webSocketController := http.NewWebSocketController(mouseAdapter, keyboardAdapter)
	mainController, err := http.NewMainController()
	if err != nil {
		return mouse, keyboard, nil, err
	}
	httpHandler := http.NewHandler(webSocketController, mainController)
	return mouse, keyboard, httpHandler, nil
}
