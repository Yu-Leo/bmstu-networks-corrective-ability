package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Syndromes struct {
	cfg           *Config
	SyndromeTable map[int]map[uint64]uint64
	errorClasses  [][]uint64
}

func NewSyndromes(cfg *Config, errorClasses [][]uint64) *Syndromes {
	return &Syndromes{
		cfg:          cfg,
		errorClasses: errorClasses,
	}
}

func (s *Syndromes) Calculate() {

	s.SyndromeTable = make(map[int]map[uint64]uint64, len(s.errorClasses))
	for i := range s.errorClasses {
		s.SyndromeTable[i] = make(map[uint64]uint64, len(s.errorClasses[i]))
		for _, errVec := range s.errorClasses[i] {
			_, syndrome := OperationO(errVec, s.cfg.genPolynomial)
			s.SyndromeTable[i][syndrome] = errVec
		}
	}
}

func (s *Syndromes) GetHandler() Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./templates/syndromes.html")
		if err != nil {
			fmt.Println(err)
		}
		SyndromeTableStr := make(map[int]map[string]string, len(s.SyndromeTable))
		for i := range s.SyndromeTable {
			SyndromeTableStr[i] = make(map[string]string, len(s.SyndromeTable[i]))
			for syndrome, err := range s.SyndromeTable[i] {
				SyndromeTableStr[i][fmt.Sprintf("%b", syndrome)] = fmt.Sprintf("%b", err)
			}
		}
		err = tmpl.Execute(w, SyndromeTableStr)
		if err != nil {
			fmt.Println(err)
		}
	}
}
