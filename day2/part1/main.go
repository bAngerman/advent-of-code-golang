package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	pwd, _ := os.Getwd()

	f, err := os.Open(pwd + "/day2/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	fullStrings := make([]string, 0)

	for scanner.Scan() {
		if err != nil {
			log.Fatal(err)
		}
		fullStrings = append(fullStrings, scanner.Text())
	}

	validPasswordCount := 0
	invalidPasswordCount := 0

	for _, val := range fullStrings {
		// Split on colon to divide password and policy
		passwordAndPolicy := strings.Split(val, ":")
		policy, password := passwordAndPolicy[0], passwordAndPolicy[1]

		// Split policy on white space to get chars
		policyNumsAndChar := strings.Split(policy, " ")
		policyNums, policyChar := policyNumsAndChar[0], policyNumsAndChar[1]

		// Split policy on dash to get high and low requirements
		policyVals := strings.Split(policyNums, "-")
		policyLow, _ := strconv.Atoi(policyVals[0])
		policyHigh, _ := strconv.Atoi(policyVals[1])

		passwordCharCount := 0
		validPassword := true

		for _, char := range strings.Split(strings.Trim(password, " "), "") {
			if char == policyChar {
				passwordCharCount++
			}
		}

		// log.Println("Char:")
		// log.Println(policyChar)

		// log.Println("Password:")
		// log.Println(password)

		// log.Println("Low and High:")
		// log.Println(policyLow)
		// log.Println(policyHigh)

		if passwordCharCount < policyLow {
			validPassword = false
		}

		if passwordCharCount > policyHigh {
			validPassword = false
		}

		if validPassword == true {
			validPasswordCount++
		} else {
			invalidPasswordCount++
		}
	}

	log.Println("Valid Password Count:")
	log.Println(validPasswordCount)

	log.Println("Invalid Password Count:")
	log.Println(invalidPasswordCount)
}
