package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// getRatesDec will convert bits in string to their decimal
// value
func getRatesDec(gam, eps *string) (int, int) {
	nGB, err := strconv.ParseInt(*gam, 2, 64)
	if err != nil {
		log.Println(err)
	}
	nEB, err := strconv.ParseInt(*eps, 2, 64)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(nGB, nEB)
	return int(nGB), int(nEB)
}

// getRatesInBits will determine the most common bit in each
// index location and generate gamma (most used) and epsilon (least used)
// and return a string for each.
func getRatesInBits(r []string) (string, string) {
	var gammaBits, epsilonBits string
	var one, zero int

	for k := 1; k <= len(r[0]); k++ {
		for _, v := range r {
			b := v[k-1 : k]
			if b == "0" {
				zero++
			} else {
				one++
			}
		}
		if one >= zero {
			gammaBits = gammaBits + "1"
			epsilonBits = epsilonBits + "0"
		} else {
			gammaBits = gammaBits + "0"
			epsilonBits = epsilonBits + "1"
		}
		// reset one and zero bit counts
		one, zero = 0, 0
	}
	return gammaBits, epsilonBits
}

func main() {
	report := strings.Split(rawInput, "\n")
	gammaBits, epsilonBits := getRatesInBits(report)
	gammaRate, epsilonRate := getRatesDec(&gammaBits, &epsilonBits)

	fmt.Printf("Gamma bits:'%s', Epsilon bits:'%s'\n", gammaBits, epsilonBits)
	fmt.Printf("Epsilon rate:'%d', Epsilon rate:'%d'\n", gammaRate, epsilonRate)
	fmt.Printf("Power consumption (gama rate * epsilon rate):'%d'\n", gammaRate*epsilonRate)
}
