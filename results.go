package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Results struct {
}

func NewResults() *Results {
	return &Results{}
}

func (s *Results) Calculate() {
	// TODO
}

func (s *Results) GetHandler() Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./templates/results.html")
		if err != nil {
			fmt.Println(err)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			fmt.Println(err)
		}
	}
}
