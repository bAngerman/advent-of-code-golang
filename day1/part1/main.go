package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {

	pwd, _ := os.Getwd()

	f, err := os.Open(pwd + "/day1/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	numbers := make([]int, 0)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, num)
	}

	targetSum := 2020
	solved := false

	for _, val1 := range numbers {
		if solved {
			break
		}
		// fmt.Println(idx, val)
		for _, val2 := range numbers {
			// log.Println(val1)
			if val1 == val2 {
				continue
			}

			if val1+val2 == targetSum {
				log.Println("Sum Reached")
				log.Println(val1)
				log.Println(val2)

				log.Println("Multiplied value:")
				log.Println(val1 * val2)
				solved = true
				break
			}

		}
	}
}
