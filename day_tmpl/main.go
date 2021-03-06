package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// type color struct {
// 	Amt   uint
// 	Color string
// }

func main() {

	rules := parseRules("/day10/input.txt")
	fmt.Printf("Part 1: %d\n", part1(rules))
	// fmt.Printf("Part 2: %d\n", part2(rules))
}

// parseRules returns a map of mainColors mapped to the rest of the rules colors
func parseRules(path string) []int {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + path)
	if err != nil {
		panic(err)
	}

	rules := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())

		if err != nil {
			panic(err)
		}

		rules = append(rules, num)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return rules
}

func part1(rules []int) int {
	return 0
}
