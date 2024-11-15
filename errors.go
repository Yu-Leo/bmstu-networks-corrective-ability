package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Errors struct {
}

func NewErrors() *Errors {
	return &Errors{}
}

func (e *Errors) Calculate() {
	// TODO
}

func (e *Errors) GetHandler() Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./templates/errors.html")
		if err != nil {
			fmt.Println(err)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			fmt.Println(err)
		}
	}
}
