package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Results struct {
	cfg *Config
}

func NewResults(cfg *Config) *Results {
	return &Results{
		cfg: cfg,
	}
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
