package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"pilot/internal/application"
	"pilot/internal/logic"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	fmt.Println("Starting HTTP server...")
	var httpServer *http.Server
	mouseController, keyboardController, httpHandler, err := application.ResolveDependencies()
	if err != nil {
		println(err.Error())
	} else {
		httpServer = &http.Server{
			Addr:    ":8080",
			Handler: httpHandler,
		}
		go func() {
			err = httpServer.ListenAndServe()
			if err != nil {
				println(err.Error())
			}
		}()
	}
    fmt.Println("HTTP server started")
	<-ctx.Done()
	shutdown(httpServer, mouseController, keyboardController)
}

func shutdown(
	httpServer *http.Server,
	mouseController *logic.MouseController,
	keyboardController *logic.KeyboardController,
) {
	fmt.Println("Shutting down...")
	if httpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(ctx); err != nil {
			println(err.Error())
		}
	}
	if mouseController != nil {
		if err := mouseController.Stop(); err != nil {
			println(err.Error())
		}
	}
	if keyboardController != nil {
		if err := keyboardController.Stop(); err != nil {
			println(err.Error())
		}
	}
}
