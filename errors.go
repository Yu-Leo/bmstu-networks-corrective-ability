package main

import (
	"fmt"
	"html/template"
	"math/bits"
	"net/http"
)

type Errors struct {
	ErrorClasses [][]uint64
	cfg          *Config
}

func NewErrors(cfg *Config) *Errors {
	return &Errors{
		cfg: cfg,
	}
}

func (e *Errors) Calculate() {
	e.ErrorClasses = make([][]uint64, e.cfg.N+1)
	for i := uint64(1); i <= e.cfg.N; i++ {
		size := factorial(e.cfg.N) / factorial(e.cfg.N-i) / factorial(i)
		e.ErrorClasses[i] = make([]uint64, 0, size)
	}

	for i := uint64(1); i < powBinary(e.cfg.N); i++ {
		class := bits.OnesCount64(i)
		e.ErrorClasses[class] = append(e.ErrorClasses[class], i)
	}
}

func (e *Errors) GetHandler() Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./templates/errors.html")
		if err != nil {
			fmt.Println(err)
		}

		errorsCountView := make([]int, len(e.ErrorClasses))

		errorsRawView := make([][]string, len(e.ErrorClasses))
		for class, errorClass := range e.ErrorClasses {
			errorsRawView[class] = make([]string, len(errorClass))
			for i, err := range errorClass {
				errorsRawView[class][i] = fmt.Sprintf("%b", err)
			}
			errorsCountView[class] = len(errorsRawView[class])
		}
		data := map[string]any{
			"errorsCount": errorsCountView,
			"errorsRaw":   errorsRawView,
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			fmt.Println(err)
		}
	}
}
