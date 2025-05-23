package application

import (
	"pilot/internal/infrastructure/http"
	"pilot/internal/logic"
)

func ResolveDependencies() (*http.Handler, error) {
	mouseController, err := logic.NewMouseController()
	if err != nil {
		return nil, err
	}
	keyboardController, err := logic.NewKeyboardController()
	if err != nil {
		return nil, err
	}
	webSocketController := http.NewWebSocketController(mouseController, keyboardController)
	httpController, err := http.NewHttpController()
	if err != nil {
		return nil, err
	}
	handler := http.NewHandler(webSocketController, httpController)
	return &handler, nil
}
