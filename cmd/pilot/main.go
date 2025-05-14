package main

import (
	"fmt"
	"net/http"
	"pilot/internal/application"
)

func main() {
	fmt.Println("Starting server")
	handler, err := application.ResolveDependencies()
	if err != nil {
		panic(err.Error())
	}
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(err.Error())
	}
}
