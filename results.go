package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ResultRow struct {
	CorrectiveAbility string
	Count             uint64
	DetectedCount     uint64
	ClassSize         int
}

type Results struct {
	cfg           *Config
	result        []ResultRow
	errorClasses  [][]uint64
	SyndromeTable map[uint64]uint64
}

func NewResults(cfg *Config, errorClasses [][]uint64, SyndromeTable map[uint64]uint64) *Results {
	return &Results{
		cfg:           cfg,
		errorClasses:  errorClasses,
		SyndromeTable: SyndromeTable,
	}
}

func (s *Results) Calculate() {
	s.result = make([]ResultRow, s.cfg.N+1)

	for class, errorClass := range s.errorClasses {
		var correctedCounter, detectedCounter uint64
		for _, errorVector := range errorClass {
			transferredVector := s.cfg.codedVector ^ errorVector
			if s.cfg.debug && class == 1 {
				fmt.Printf("\ntransferredVector: %b\n", transferredVector)
			}
			syndrome := GetDivisionRemainder(transferredVector, s.cfg.genPolynomial)
			if s.cfg.debug && class == 1 {
				fmt.Printf("syndrome: %b | error: %b\n", syndrome, s.SyndromeTable[syndrome])
			}
			if syndrome == 0 {
				continue
			}
			detectedCounter++
			correctedVector := transferredVector ^ s.SyndromeTable[syndrome]

			if s.cfg.debug && class == 1 {
				fmt.Printf("correctedVector: %b\n", correctedVector)
			}
			if correctedVector == s.cfg.codedVector {
				correctedCounter++
				if s.cfg.debug && class == 1 {
					fmt.Printf("eror corrected successfully | counter: %d\n", correctedCounter)
				}
			}
		}
		s.result[class] = ResultRow{
			fmt.Sprintf("%.2f", float64(correctedCounter)*100/float64(len(errorClass))),
			correctedCounter,
			detectedCounter,
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
