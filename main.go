package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request)

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		fmt.Println(err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	errors := NewErrors()
	errorsAndSyndromes := NewErrorsAndSyndromes()
	syndromes := NewSyndromes()
	results := NewResults()

	errors.Calculate()
	errorsAndSyndromes.Calculate()
	syndromes.Calculate()
	results.Calculate()

	http.HandleFunc("/", mainPageHandler)
	http.HandleFunc("/errors/", errors.GetHandler())
	http.HandleFunc("/errors_and_syndromes/", errorsAndSyndromes.GetHandler())
	http.HandleFunc("/syndromes/", syndromes.GetHandler())
	http.HandleFunc("/results/", results.GetHandler())
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
