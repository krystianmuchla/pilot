package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"pilot/internal/application"
	"pilot/internal/infrastructure/linux"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	fmt.Println("Starting HTTP server...")
	var httpServer *http.Server
	mouse, keyboard, httpHandler, err := application.ResolveDependencies()
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
	shutdown(httpServer, mouse, keyboard)
}

func shutdown(httpServer *http.Server, mouse *linux.Mouse, keyboard *linux.Keyboard) {
	fmt.Println("Shutting down...")
	if httpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(ctx); err != nil {
			println(err.Error())
		}
	}
	if mouse != nil {
		if err := mouse.Close(); err != nil {
			println(err.Error())
		}
	}
	if keyboard != nil {
		if err := keyboard.Close(); err != nil {
			println(err.Error())
		}
	}
}
