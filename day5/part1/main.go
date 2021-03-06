package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func makeArr(len int) []int {
	a := []int{}
	for i := 0; i < len; i++ {
		a = append(a, i)
	}
	return a
}

func sliceLow(arr []int) []int {
	l := len(arr)
	middle := l / 2

	for len(arr) > middle {
		arr = removeIndex(arr, len(arr)-1)
	}

	return arr
}

func sliceHigh(arr []int) []int {
	l := len(arr)
	middle := l / 2

	for len(arr) > middle {
		arr = removeIndex(arr, 0)
	}

	return arr
}

func findMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func main() {

	pwd, _ := os.Getwd()

	f, err := os.Open(pwd + "/day5/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := make([]string, 0)

	for scanner.Scan() {
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, scanner.Text())
	}

	seatIds := make([]int, 0)

	for _, seatString := range lines {
		rowTarget := makeArr(128)
		columnTarget := makeArr(8)

		// log.Println("RowTarget: ", rowTarget)
		// log.Println("Columntarget: ", columnTarget)

		seatChars := strings.Split(seatString, "")

		for charIdx, char := range seatChars {
			if charIdx < 7 {
				switch char {
				case "F":
					rowTarget = sliceLow(rowTarget)
				case "B":
					rowTarget = sliceHigh(rowTarget)
				}
			} else {
				switch char {
				case "R":
					columnTarget = sliceHigh(columnTarget)
				case "L":
					columnTarget = sliceLow(columnTarget)
				}

			}
		}

		rowTargetFinal := rowTarget[0]
		columnTargetFinal := columnTarget[0]

		seatIds = append(seatIds, ((rowTargetFinal * 8) + columnTargetFinal))
	}

	log.Println("Seat Ids: ", seatIds)
	min, max := findMinAndMax(seatIds)

	log.Println("Max is: ", max)
	log.Println("Min is: ", min)
}
