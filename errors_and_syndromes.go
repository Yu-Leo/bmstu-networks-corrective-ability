package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ErrorsAndSyndromes struct {
	cfg *Config
}

func NewErrorsAndSyndromes(cfg *Config) *ErrorsAndSyndromes {
	return &ErrorsAndSyndromes{
		cfg: cfg,
	}
}

func (e *ErrorsAndSyndromes) Calculate() {
	// TODO
}

func (e *ErrorsAndSyndromes) GetHandler() Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./templates/errors_and_syndromes.html")
		if err != nil {
			fmt.Println(err)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			fmt.Println(err)
		}
	}
}
