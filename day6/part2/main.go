package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func contains(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}

	return false
}

func removeIndex(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func main() {

	pwd, _ := os.Getwd()

	f, err := os.Open(pwd + "/day6/input.txt")

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

	totalScore := 0

	groupIdx := 0
	personIdx := 0
	mustHaveVals := make([]string, 0)

	for idx, line := range rows {

		lineChars := strings.Split(line, "")

		// This is the first person in the group, any answer
		// that they have the other people must also have.
		if personIdx == 0 {
			for _, char := range lineChars {
				mustHaveVals = append(mustHaveVals, char)
			}
		}

		// Compare what's in this line vs what's in the first
		// person, remove from the mustHaveVals arr if dont have
		if line != "" && personIdx != 0 {
			// Loop over must haves
			for _, char := range mustHaveVals {
				// If we dont have the char in the line, we gotta remove it Sadge
				if !contains(lineChars, char) {

					// log.Println(fmt.Sprintf("We seem to be missing %s in %s", char, mustHaveVals))

					for idx, mustHaveChar := range mustHaveVals {
						if char == mustHaveChar {
							// log.Println("Removing char ", char)
							mustHaveVals = removeIndex(mustHaveVals, idx)
						}
					}
				}
			}
		}

		// Check if new group, if it is then reset counters.
		if line == "" || idx == len(rows)-1 {

			// Add cumulative group score to total
			totalScore = totalScore + len(mustHaveVals)

			// Reset vals
			mustHaveVals = make([]string, 0)
			// log.Println("Group index: ", groupIdx)
			// log.Println("Group Score: ", groupScore)
			personIdx = 0
			groupIdx++
			continue
		}

		personIdx++
	}

	log.Println("Total Answers: ", totalScore)
}
