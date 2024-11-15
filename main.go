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

type Config struct {
	N             uint64
	K             uint64
	vector        uint64
	genPolynomial uint64
	debug         bool
}

func main() {
	var config = &Config{
		N:             15,
		K:             11,
		vector:        32050, // 111.1101.0011.0010
		genPolynomial: 19,    // 10011
		debug:         true,
	}

	errors := NewErrors(config)
	errors.Calculate()

	errorsAndSyndromes := NewErrorsAndSyndromes(config)
	errorsAndSyndromes.Calculate()

	syndromes := NewSyndromes(config, errors.ErrorClasses)
	syndromes.Calculate()

	results := NewResults(config, errors.ErrorClasses, syndromes.SyndromeTable)
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
