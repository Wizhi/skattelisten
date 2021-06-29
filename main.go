package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"os"
	"strconv"
	"strings"
)

type company struct {
	CVR           *string
	Name          *string
	SE            *string
	IncomeYear    *int64
	CompanyType   *string
	TaxableIncome *int64
	Deficit       *int64
	CorporateTax  *int64
}

func main() {
	var src, dst string

	flag.StringVar(&src, "src", "", "")
	flag.StringVar(&dst, "dst", "", "")

	flag.Parse()

	if src == "" {
		os.Exit(1)
	}

	input, err := os.Open(src)

	if err != nil {
		os.Exit(2)
	}

	defer input.Close()

	output, err := os.Create(dst)

	if err != nil {
		os.Exit(2)
	}

	defer output.Close()

	s := bufio.NewScanner(input)
	w := bufio.NewWriter(output)

	// Get rid of headers
	s.Scan()

	for s.Scan() {
		cols := strings.Split(s.Text(), ",")

		var incomeYear, taxableIncome, deficit, corporateTax *int64

		if a, err := strconv.ParseInt(cols[3], 10, 64); err == nil {
			incomeYear = &a
		}

		if a, err := strconv.ParseInt(cols[8], 10, 64); err == nil {
			taxableIncome = &a
		}

		if a, err := strconv.ParseInt(cols[9], 10, 64); err == nil {
			deficit = &a
		}

		if a, err := strconv.ParseInt(cols[10], 10, 64); err == nil {
			corporateTax = &a
		}

		comp := company{
			CVR:           &cols[0],
			Name:          &cols[1],
			SE:            &cols[2],
			IncomeYear:    incomeYear,
			CompanyType:   &cols[6],
			TaxableIncome: taxableIncome,
			Deficit:       deficit,
			CorporateTax:  corporateTax,
		}

		json.NewEncoder(w).Encode(&comp)
	}
}
