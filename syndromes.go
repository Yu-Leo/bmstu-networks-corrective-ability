package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Syndromes struct {
	cfg *Config
}

func NewSyndromes(cfg *Config) *Syndromes {
	return &Syndromes{
		cfg: cfg,
	}
}

func (s *Syndromes) Calculate() {
	// TODO
}

func (s *Syndromes) GetHandler() Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./templates/syndromes.html")
		if err != nil {
			fmt.Println(err)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			fmt.Println(err)
		}
	}
}
