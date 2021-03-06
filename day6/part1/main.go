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
	groupScore := 0
	alreadyAnsweredVals := make([]string, 0)

	for idx, line := range rows {

		// if idx == 20 {
		// 	break
		// }

		lineChars := strings.Split(line, "")

		// log.Println(lineChars)

		for _, char := range lineChars {
			// If the element is not in the array we can add to counts
			if !contains(alreadyAnsweredVals, char) {
				groupScore++
				alreadyAnsweredVals = append(alreadyAnsweredVals, char)
				// log.Println("Already answered vals: ", alreadyAnsweredVals)
			} else {
				// s := fmt.Sprintf("Skipped %s as it was already in answers", char)
				// log.Println(s)
			}
		}

		// Check if new group, if it is then reset counters.
		if line == "" || idx == len(rows)-1 {

			// Add cumulative group score to total
			totalScore = totalScore + groupScore

			// Reset already answered and other vals.
			alreadyAnsweredVals = make([]string, 0)
			// log.Println("Group index: ", groupIdx)
			// log.Println("Group Score: ", groupScore)
			groupScore = 0
			groupIdx++
			continue
		}
	}

	log.Println("Total Answers: ", totalScore)
}
