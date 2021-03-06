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

	// Split on colon to divide password and policy
	for _, val := range fullStrings {
		passwordAndPolicy := strings.Split(val, ":")
		policy, password := passwordAndPolicy[0], passwordAndPolicy[1]

		// Split policy on white space to get chars
		policyNumsAndChar := strings.Split(policy, " ")
		policyNums, policyChar := policyNumsAndChar[0], policyNumsAndChar[1]

		// Split policy on dash to get position requirements
		policyVals := strings.Split(policyNums, "-")
		posOne, _ := strconv.Atoi(policyVals[0])
		posTwo, _ := strconv.Atoi(policyVals[1])
		validPassword := false

		password = strings.Trim(password, " ")

		// log.Println("Char: " + policyChar)

		// log.Println("Char position one:")
		// log.Println(posOne)

		// log.Println("Char position two:")
		// log.Println(posTwo)

		for idx, char := range strings.Split(strings.Trim(password, " "), "") {
			// log.Println("Idx:")
			// log.Println(idx)

			if idx+1 == posOne {
				if char == policyChar {
					log.Println(password + ": position 1 set")
					validPassword = true
				}
			}

			if idx+1 == posTwo {
				if char == policyChar {
					if validPassword != true {
						log.Println(password + ": position 2 set")
						validPassword = true
					} else {
						log.Println(password + ": both positions, BAD")
						validPassword = false
					}
				}
			}
		}

		if validPassword == true {
			// log.Println("Password Valid")
			validPasswordCount++
		} else {
			// log.Println(password + " is invalid")
		}
	}

	log.Println("Valid Password Count:")
	log.Println(validPasswordCount)
}
