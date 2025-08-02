package http

import (
	"html/template"
	"net/http"
)

type MainController struct {
	template *template.Template
}

func NewMainController() (*MainController, error) {
	temp, err := template.ParseFiles("web/index.html")
	if err != nil {
		return nil, err
	}
	return &MainController{template: temp}, nil
}

func (controller MainController) Handle(responseWriter http.ResponseWriter, request *http.Request) error {
	if request.RequestURI != "/" {
		responseWriter.Header().Set("Location", "/")
		responseWriter.WriteHeader(http.StatusFound)
	} else {
		err := controller.template.Execute(responseWriter, nil)
		if err != nil {
			return err
		}
	}
	return nil
}
