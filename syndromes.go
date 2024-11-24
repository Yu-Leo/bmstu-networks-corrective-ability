package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Syndromes struct {
	cfg           *Config
	SyndromeTable map[uint64]uint64
	errorClasses  [][]uint64
}

func NewSyndromes(cfg *Config, errorClasses [][]uint64) *Syndromes {
	return &Syndromes{
		cfg:          cfg,
		errorClasses: errorClasses,
	}
}

func (s *Syndromes) Calculate() {
	s.SyndromeTable = make(map[uint64]uint64, len(s.errorClasses[1]))
	for _, errVec := range s.errorClasses[1] {
		syndrome := GetDivisionRemainder(s.cfg.codedVector^errVec, s.cfg.genPolynomial)
		s.SyndromeTable[syndrome] = errVec
	}
}

func (s *Syndromes) GetHandler() Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./templates/syndromes.html")
		if err != nil {
			fmt.Println(err)
		}
		SyndromeTableStr := make(map[string]string, len(s.SyndromeTable))
		for syndrome, err := range s.SyndromeTable {
			SyndromeTableStr[fmt.Sprintf("%b", syndrome)] = fmt.Sprintf("%b", err)
		}
		err = tmpl.Execute(w, SyndromeTableStr)
		if err != nil {
			fmt.Println(err)
		}
	}
}
