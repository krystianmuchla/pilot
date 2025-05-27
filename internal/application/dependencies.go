package application

import (
	"pilot/internal/infrastructure/http"
	"pilot/internal/logic"
)

func ResolveDependencies() (*logic.MouseController, *logic.KeyboardController, *http.Handler, error) {
	mouseController, err := logic.NewMouseController()
	if err != nil {
		return nil, nil, nil, err
	}
	keyboardController, err := logic.NewKeyboardController()
	if err != nil {
		return mouseController, nil, nil, err
	}
	webSocketController := http.NewWebSocketController(mouseController, keyboardController)
	httpController, err := http.NewHttpController()
	if err != nil {
		return mouseController, keyboardController, nil, err
	}
	httpHandler := http.NewHandler(webSocketController, httpController)
	return mouseController, keyboardController, httpHandler, nil
}
