package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
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

func (pp passport) isValidByr() bool {
	validity := false
	i, err := strconv.Atoi(pp.byr)
	if err != nil {
		return validity
	}

	if (i >= 1920) && (i <= 2002) {
		validity = true
	}

	return validity
}

func (pp passport) isValidIyr() bool {
	validity := false
	i, err := strconv.Atoi(pp.iyr)
	if err != nil {
		return validity
	}

	if (i >= 2010) && (i <= 2020) {
		validity = true
	}

	return validity
}

func (pp passport) isValidEyr() bool {
	validity := false
	i, err := strconv.Atoi(pp.eyr)
	if err != nil {
		return validity
	}

	if (i >= 2020) && (i <= 2030) {
		validity = true
	}

	return validity
}

func (pp passport) isValidHgt() bool {
	validity := false

	isCm := strings.Contains(pp.hgt, "cm")
	isIn := strings.Contains(pp.hgt, "in")

	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	height := reg.ReplaceAllString(pp.hgt, "")

	i, err := strconv.Atoi(height)

	if err != nil {
		return validity
	}

	if isCm {
		if (i >= 150) && (i <= 193) {
			validity = true
		}
	} else if isIn {
		if (i >= 59) && (i <= 76) {
			validity = true
		}
	}

	return validity
}

func (pp passport) isValidHcl() bool {
	validity := false

	reg, err := regexp.Compile("^#[a-z0-9]{6}$")
	if err != nil {
		log.Fatal(err)
	}

	validity = reg.MatchString(pp.hcl)

	return validity
}

func (pp passport) isValidEcl() bool {
	validity := false
	validEyeColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	validity = stringInSlice(pp.ecl, validEyeColors)

	return validity
}

func (pp passport) isValidPid() bool {
	validity := false

	reg, err := regexp.Compile("^[0-9]{9}$")
	if err != nil {
		log.Fatal(err)
	}

	validity = reg.MatchString(pp.pid)

	return validity
}

func (pp passport) IsValidPassport() bool {
	validity := false

	if pp.isValidByr() &&
		pp.isValidIyr() &&
		pp.isValidEyr() &&
		pp.isValidHgt() &&
		pp.isValidHcl() &&
		pp.isValidEcl() &&
		pp.isValidPid() {
		validity = true
	}

	// cid is optional

	return validity
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
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
