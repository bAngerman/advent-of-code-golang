package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {

	pwd, _ := os.Getwd()

	f, err := os.Open(pwd + "/day3/input.txt")

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

	travelXArr := []int{1, 3, 5, 7, 1}
	travelYArr := []int{1, 1, 1, 1, 2}

	treeHitCounts := []int{}

	for solutionIdx := range travelXArr {

		lineCharCount := len(rows[0])
		treeHitCounter := 0
		currentPosX := 0

		// Loop over outer lines
		for lineIdx, line := range rows {

			// Split line space to get chars
			lineParts := strings.Split(line, "")
			// Find positions
			posStart := (currentPosX % lineCharCount)

			if (lineIdx % travelYArr[solutionIdx]) != 0 {
				continue // Skip this iteration as we moved down more than 1
			}
			// If not on first line, and the first char in line is the #..
			if lineIdx != 0 && lineParts[posStart] == "#" {
				treeHitCounter++
			}
			currentPosX = currentPosX + travelXArr[solutionIdx] // Move right
		}
		treeHitCounts = append(treeHitCounts, treeHitCounter)

	}

	multipliedTreeHits := 1
	for _, num := range treeHitCounts {
		multipliedTreeHits = multipliedTreeHits * num
	}

	log.Println("Trees Hit Counts: ", treeHitCounts)
	log.Println("Tree Hits Multipied: ", multipliedTreeHits)
}
