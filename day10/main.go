package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	rules := parseRules("/day10/input.txt")
	fmt.Printf("Part 1: %d\n", part1(rules))
	fmt.Printf("Part 2: %d\n", part2(rules))
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

func inSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func minMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func part1(rules []int) int {
	jDiffOne := 0
	jDiffThree := 0
	usedAdapters := make([]int, 0)
	joltage := 0

	for i := 0; i < len(rules); i++ {

		validAdapters := make([]int, 0)

		for _, adapter := range rules {

			if inSlice(adapter, usedAdapters) {
				continue
			}

			if (adapter-joltage) > 0 && (adapter-joltage) <= 3 {
				validAdapters = append(validAdapters, adapter)
			}
		}

		if len(validAdapters) != 0 {
			minAdapter, _ := minMax(validAdapters)

			// fmt.Printf("Joltage: %v\t\tAdapater: %v\t\tDiff: %v\n", joltage, minAdapter, (minAdapter - joltage))

			usedAdapters = append(usedAdapters, minAdapter)

			if (minAdapter - joltage) == 3 {
				jDiffThree = jDiffThree + 1
			} else if minAdapter-joltage == 1 {
				jDiffOne = jDiffOne + 1
			}

			// Set joltage to new minAdapter value
			joltage = minAdapter
		} else {
			// fmt.Printf("Joltage: %v has no adapter.\n", joltage)
		}

	}

	jDiffThree = jDiffThree + 1

	// fmt.Printf("1 diff: %v\n3 diff: %v\n", jDiffOne, jDiffThree)

	return (jDiffOne * jDiffThree)
}

func part2(rules []int) int {
	arrangements := make([][]int, 0)

	for joltage := 0; joltage < len(rules); joltage++ {
		validAdapters := make([]int, 0)

		for _, adapter := range rules {
			// fmt.Printf("adapter: %v,\t\t joltage: %v\n", adapter, joltage)
			if (adapter-joltage) > 0 && (adapter-joltage) <= 3 {
				validAdapters = append(validAdapters, adapter)
			}
		}

		if len(validAdapters) != 0 {
			for _, vAdapter := range validAdapters {

				// We need to add the variations onto the arrangements array.
				if len(arrangements) == 0 {
					arrangements = append(arrangements, []int{vAdapter})
				} else {
					for aIndex, arrangement := range arrangements {
						minAdapter, _ := minMax(arrangement)
						fmt.Printf("vAdapter: %v,\t\t minAdapter: %v\n", vAdapter, minAdapter)
						// If this one works... merge into arrangement
						if vAdapter > minAdapter {
							arrangements[aIndex] = append(arrangements[aIndex], minAdapter)
						}
					}
				}

				fmt.Printf("Arrangements: %v\n", arrangements)

			}

			fmt.Printf("Arrangements: %v\n", arrangements)

			// // minAdapter, _ := minMax(validAdapters)

			// // fmt.Printf("Joltage: %v\t\tAdapater: %v\t\tDiff: %v\n", joltage, minAdapter, (minAdapter - joltage))

			// // 	usedAdapters = append(usedAdapters, minAdapter)

			// // 	if (minAdapter - joltage) == 3 {
			// // 		jDiffThree = jDiffThree + 1
			// // 	} else if minAdapter-joltage == 1 {
			// // 		jDiffOne = jDiffOne + 1
			// // 	}
		}
	}

	return len(arrangements)
}
