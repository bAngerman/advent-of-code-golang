package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	rules := parseRules("/day9/input.txt")
	part1 := part1(rules)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2(rules, part1))
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

func part1(numbers []int) int {
	invalidNum := 0
	offset := 25

	for idx, val := range numbers {
		if idx < offset {
			continue
		} else {
			// Check if number is any of the 25 previous numbers
			subset := numbers[(idx - offset):idx]
			sum := val
			valid := false

			for idx1, num1 := range subset {
				for idx2, num2 := range subset {
					if idx1 != idx2 {
						if (num1 + num2) == sum {

							// fmt.Printf("%d + %d = %d\n", num1, num2, sum)
							valid = true
						}
					}
				}
			}

			if valid == false {
				return sum
			}
		}
	}

	return invalidNum
}

func part2(numbers []int, target int) int {
	runningSumNums := make([]int, 0)

	for idx := range numbers {

		subset := numbers[idx : len(numbers)-1]
		runningSum := 0
		runningSumNums = make([]int, 0)
		solved := false

		for _, innerNum := range subset {
			runningSum = runningSum + innerNum
			runningSumNums = append(runningSumNums, innerNum)

			// fmt.Printf("Running sum is: %v\n", runningSum)
			// fmt.Printf("Nums: %v\n", runningSumNums)

			if runningSum == target {
				solved = true
				break
			} else if runningSum > target {
				break
			}
		}

		if solved == true {
			break
		}
	}

	min, max := minMax(runningSumNums)

	return min + max
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
