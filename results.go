package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ResultRow struct {
	CorrectiveAbility string
	Count             uint64
	ClassSize         int
}

type Results struct {
	cfg           *Config
	result        []ResultRow
	errorClasses  [][]uint64
	SyndromeTable map[int]map[uint64]uint64
}

func NewResults(cfg *Config, errorClasses [][]uint64, SyndromeTable map[int]map[uint64]uint64) *Results {
	return &Results{
		cfg:           cfg,
		errorClasses:  errorClasses,
		SyndromeTable: SyndromeTable,
	}
}

func (s *Results) Calculate() {
	s.result = make([]ResultRow, s.cfg.N+1)

	for class, errorClass := range s.errorClasses {
		var correctedCounter uint64
		for _, errorVector := range errorClass {
			transferredVector := ImposeError(s.cfg.vector, errorVector)
			if s.cfg.debug && class == 1 {
				fmt.Printf("\ntransferredVector: %b\n", transferredVector)
			}
			_, syndrome := OperationO(transferredVector, s.cfg.genPolynomial)
			if s.cfg.debug && class == 1 {
				fmt.Printf("syndrome: %b\n", syndrome)
			}
			if syndrome == 0 {
				continue
			}
			correctedVector := ImposeError(transferredVector, s.SyndromeTable[class][syndrome])
			if s.cfg.debug && class == 1 {
				fmt.Printf("correctedVector: %b\n", correctedVector)
			}
			if correctedVector == s.cfg.vector {
				correctedCounter++
				if s.cfg.debug && class == 1 {
					fmt.Printf("eror corrected successfully | counter: %d\n", correctedCounter)
				}
			}
		}
		s.result[class] = ResultRow{
			fmt.Sprintf("%.2f", float64(correctedCounter)*100/float64(len(errorClass))),
			correctedCounter,
			len(errorClass),
		}
	}
}

func (s *Results) GetHandler() Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./templates/results.html")
		if err != nil {
			fmt.Println(err)
		}
		err = tmpl.Execute(w, s.result)
		if err != nil {
			fmt.Println(err)
		}
	}
}
