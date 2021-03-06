package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type passport struct {
	byr string // (Birth Year)
	iyr string // (Issue Year)
	eyr string // (Expiration Year)
	hgt string // (Height)
	hcl string // (Hair Color)
	ecl string // (Eye Color)
	pid string // (Passport ID)
	cid string // (Country ID)
}

func (pp passport) IsValidPassport() bool {
	validity := ""

	if pp.byr == "" {
		validity = "Invalid due to byr"
	}
	if pp.iyr == "" {
		validity = "Invalid due to iyr"
	}
	if pp.eyr == "" {
		validity = "Invalid due to eyr"
	}
	if pp.hgt == "" {
		validity = "Invalid due to hgt"
	}
	if pp.hcl == "" {
		validity = "Invalid due to hcl"
	}
	if pp.ecl == "" {
		validity = "Invalid due to ecl"
	}
	if pp.pid == "" {
		validity = "Invalid due to pid"
	}

	// cid is optional

	if len(validity) == 0 {
		return true
	}

	// spew.Dump(pp)
	// spew.Dump(validity)
	// spew.Dump("------------------")

	return false
}

func main() {

	pwd, _ := os.Getwd()

	f, err := os.Open(pwd + "/day4/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	rows := make([]string, 0)

	for scanner.Scan() {
		if err != nil {
			log.Fatal(err)
		}
		rows = append(rows, scanner.Text())
	}

	validPassportCount := 0

	// Passport object
	pp := new(passport)
	ppIndex := 0

	// Iterate over input lines
	for idx, line := range rows {

		if line != "" {
			lineParts := strings.Split(line, " ")

			for _, partStrings := range lineParts {
				keyVals := strings.Split(partStrings, ":")

				switch keyVals[0] {
				case "byr":
					pp.byr = keyVals[1]
				case "iyr":
					pp.iyr = keyVals[1]
				case "eyr":
					pp.eyr = keyVals[1]
				case "hgt":
					pp.hgt = keyVals[1]
				case "hcl":
					pp.hcl = keyVals[1]
				case "ecl":
					pp.ecl = keyVals[1]
				case "pid":
					pp.pid = keyVals[1]
				case "cid":
					pp.cid = keyVals[1]
				}
			}

		}

		if line == "" || idx == len(rows)-1 {
			// Check if it's valid
			if pp.IsValidPassport() {
				validPassportCount++
			}

			// Reset passport
			pp = new(passport)
			ppIndex++
			continue
		}
	}

	log.Println("Total Passport count: ", ppIndex+1)
	log.Println("Valid Passport count: ", validPassportCount)
}
