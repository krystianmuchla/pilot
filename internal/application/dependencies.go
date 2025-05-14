package application

import (
	http2 "pilot/internal/infrastructure/http"
	"pilot/internal/logic"
)

func ResolveDependencies() (*http2.Handler, error) {
	mouseController, err := logic.NewMouseController()
	if err != nil {
		return nil, err
	}
	webSocketController := http2.NewWebSocketController(mouseController)
	httpController, err := http2.NewHttpController()
	if err != nil {
		return nil, err
	}
	handler := http2.NewHandler(webSocketController, httpController)
	return &handler, nil
}
