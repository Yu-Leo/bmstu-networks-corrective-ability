package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Syndromes struct {
	cfg           *Config
	syndromeTable map[uint64]uint64
	errorVectors  []uint64
}

func NewSyndromes(cfg *Config, errorVectors []uint64) *Syndromes {
	return &Syndromes{
		cfg:          cfg,
		errorVectors: errorVectors,
	}
}

func (s *Syndromes) Calculate() {
	s.syndromeTable = make(map[uint64]uint64, len(s.errorVectors))
	for _, errVec := range s.errorVectors {
		_, syndrome := OperationO(errVec, s.cfg.genPolynomial)
		s.syndromeTable[syndrome] = errVec
	}
}

func (s *Syndromes) GetHandler() Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./templates/syndromes.html")
		if err != nil {
			fmt.Println(err)
		}
		syndromeTableStr := make(map[string]string, len(s.syndromeTable))
		for syndrome, err := range s.syndromeTable {
			syndromeTableStr[fmt.Sprintf("%b", syndrome)] = fmt.Sprintf("%b", err)
		}
		err = tmpl.Execute(w, syndromeTableStr)
		if err != nil {
			fmt.Println(err)
		}
	}
}
