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

	lineCharCount := len(rows[0])
	treeHitCounter := 0
	currentPosX := 0

	travelX := 3
	travelY := 1

	// Loop over outer lines
	for lineIdx, line := range rows {

		// Split line space to get chars..
		lineParts := strings.Split(line, "")

		// Set this to maybe change it if index goes off array
		posStart := (currentPosX % lineCharCount)

		if (lineIdx % travelY) != 0 {
			continue // Skip this iteration as we moved down more than 1
		}

		// If not on first line, and the first char in line is the #..
		if lineIdx != 0 && lineParts[posStart] == "#" {
			treeHitCounter++
		}

		currentPosX = currentPosX + travelX // Move right
	}

	log.Println("Trees Hit: ", treeHitCounter)
}
