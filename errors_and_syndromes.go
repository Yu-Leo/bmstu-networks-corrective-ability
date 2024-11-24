package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ErrorsAndSyndromes struct {
	cfg      *Config
	errorMap map[string]string
}

func NewErrorsAndSyndromes(cfg *Config) *ErrorsAndSyndromes {
	return &ErrorsAndSyndromes{
		cfg: cfg,
	}
}

func (e *ErrorsAndSyndromes) Calculate() {
	e.errorMap = make(map[string]string, powBinary(e.cfg.N))
	for i := uint64(1); i < powBinary(e.cfg.N); i++ {
		syndrome := GetDivisionRemainder(e.cfg.codedVector^i, e.cfg.genPolynomial)
		e.errorMap[fmt.Sprintf("%b", i)] = fmt.Sprintf("%b", syndrome)
	}
}

func (e *ErrorsAndSyndromes) GetHandler() Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./templates/errors_and_syndromes.html")
		if err != nil {
			fmt.Println(err)
		}
		err = tmpl.Execute(w, e.errorMap)
		if err != nil {
			fmt.Println(err)
		}
	}
}
