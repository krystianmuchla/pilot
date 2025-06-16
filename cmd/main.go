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
	fmt.Println("Starting server")
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
	shutdown := shutdownGracefully(httpServer, mouseController, keyboardController)
	<-shutdown
}

func shutdownGracefully(
	httpServer *http.Server,
	mouseController *logic.MouseController,
	keyboardController *logic.KeyboardController) <-chan struct{} {
	shutdown := make(chan struct{})
	go func() {
		osSignal := make(chan os.Signal, 1)
		signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)
		<-osSignal
		fmt.Println("Shutting down gracefully...")
		if httpServer != nil {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			err := httpServer.Shutdown(ctx)
			if err != nil {
				println(err.Error())
			}
		}
		if mouseController != nil {
			err := mouseController.Stop()
			if err != nil {
				println(err.Error())
			}
		}
		if keyboardController != nil {
			err := keyboardController.Stop()
			if err != nil {
				println(err.Error())
			}
		}
		close(shutdown)
	}()
	return shutdown
}
