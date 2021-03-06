package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type cube struct {
	Active       bool
	initialState bool
}

func main() {

	rules := parseRules("/day17/input.txt")
	fmt.Printf("Part 1: %d\n", part1(rules))
	// fmt.Printf("Part 2: %d\n", part2(rules))
}

// parseRules returns a map of mainColors mapped to the rest of the rules colors
func parseRules(path string) [][][]cube {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + path)
	if err != nil {
		panic(err)
	}

	rules := make([][][]cube, 0)
	scanner := bufio.NewScanner(file)

	row := 0

	for scanner.Scan() {
		cubes := strings.Split(scanner.Text(), "")

		for _, c := range cubes {
			active := c == "#"
			rules[0][row] = append(rules[0][row], cube{
				Active:       active,
				initialState: active,
			})
		}

		row = row + 1
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return rules
}

func countActiveCubes(cubeMap [][][]cube) int {
	activeCount := 0
	for _, cubeZ := range cubeMap {
		for _, cubeRow := range cubeZ {
			for _, cube := range cubeRow {
				if cube.Active {
					activeCount = activeCount + 1
				}
			}
		}
	}
	return activeCount
}

func setStates(cubeMap [][][]cube) [][][]cube {
	for _, cubeZ := range cubeMap {
		for _, cubeRow := range cubeZ {
			for colIdx, cube := range cubeRow {
				cubeRow[colIdx].initialState = cube.Active
			}
		}
	}
	return cubeMap
}

func cycle(cubeMap [][][]cube) [][][]cube {
	for rowIdx, cubeRow := range cubeMap {
		for colIdx, cube := range cubeRow {
			activeNeighborCount := 0

			// spew.Dump(cube)
			// spew.Dump(rowIdx)
			// spew.Dump(colIdx)

			// Check Right
			if colIdx != len(cubeRow)-1 {
				if cubeMap[rowIdx][colIdx+1].initialState {
					activeNeighborCount++
				}
			}

			// Check Left
			if colIdx != 0 {
				if cubeMap[rowIdx][colIdx-1].initialState {
					activeNeighborCount++
				}
			}

			// Check Down
			if rowIdx != len(cubeMap)-1 {
				if cubeMap[rowIdx+1][colIdx].initialState {
					activeNeighborCount++
				}
			}

			// Check Up
			if rowIdx != 0 {
				if cubeMap[rowIdx-1][colIdx].initialState {
					activeNeighborCount++
				}
			}

			spew.Dump(activeNeighborCount)

			if cube.initialState {
				if activeNeighborCount == 1 || activeNeighborCount == 4 {
					cubeRow[rowIdx].Active = false
				}
			} else {
				if activeNeighborCount == 3 {
					cubeRow[rowIdx].Active = true
				}
			}
		}
	}

	return cubeMap
}

func part1(cubeMap [][][]cube) int {
	activeCubes := 0
	iterations := 6

	for i := 0; i < iterations; i++ {
		cubeMap = cycle(cubeMap)

		// Set initial states for cubes now that they have changed.
		cubeMap = setStates(cubeMap)
		fmt.Printf("After transform: \n")
		spew.Dump(cubeMap)
	}

	activeCubes = countActiveCubes(cubeMap)

	return activeCubes
}
