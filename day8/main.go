package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type command struct {
	name string
	num  int
	ran  bool
}

func (c command) SetRan(r bool) command {
	c.ran = r
	return c
}

func (c command) Ran() bool {
	return c.ran
}

func main() {

	rules := parseRules("/day8/input.txt")
	fmt.Printf("Part 1: %d\n", part1(rules))
	fmt.Printf("Part 2: %d\n", part2(rules))
}

// parseRules returns a map of mainColors mapped to the rest of the rules colors
func parseRules(path string) []command {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + path)
	if err != nil {
		panic(err)
	}

	commands := make([]command, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		name := parts[0]
		amount, err := strconv.Atoi(parts[1])

		if err != nil {
			panic(err)
		}

		commands = append(commands, command{
			name: name,
			num:  amount,
			ran:  false,
		})

	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return commands
}

func part1(commands []command) int {
	accumulator := 0
	commandIndex := 0

	// Infinite loop
	for {
		if commands[commandIndex].Ran() {
			// Exit and return
			break
		}

		// Mark ran as true.
		commands[commandIndex] = commands[commandIndex].SetRan(true)

		switch commands[commandIndex].name {
		case "nop":
			commandIndex = commandIndex + 1 // No OP, just go to next command
		case "acc":
			accumulator = accumulator + commands[commandIndex].num // Increment accumulator by the number

			// Do this step AFTER we increment the accumulator, DOH
			commandIndex = commandIndex + 1 // Go to next  command
		case "jmp":
			commandIndex = commandIndex + commands[commandIndex].num // Modify the index by the num value
		}
	}

	return accumulator
}

func isInfiniteLoop(commands []command) (bool, int) {
	accumulator := 0
	commandIndex := 0
	isInfinite := false

	// Infinite loop
	for {

		if commandIndex == len(commands) {
			break
		}

		if commands[commandIndex].Ran() {
			// Exit and return
			isInfinite = true
			break
		}

		// Mark ran as true.
		commands[commandIndex] = commands[commandIndex].SetRan(true)

		switch commands[commandIndex].name {
		case "nop":
			commandIndex = commandIndex + 1 // No OP, just go to next command
		case "acc":
			accumulator = accumulator + commands[commandIndex].num // Increment accumulator by the number

			// Do this step AFTER we increment the accumulator, DOH
			commandIndex = commandIndex + 1 // Go to next  command
		case "jmp":
			commandIndex = commandIndex + commands[commandIndex].num // Modify the index by the num value
		}
	}

	return isInfinite, accumulator
}

func part2(commands []command) int {
	accumulator := 0
	for idx := range commands {

		newCommands := parseRules("/day8/input.txt")

		switch newCommands[idx].name {

		case "nop":
			newCommands[idx].name = "jmp"
		case "jmp":
			newCommands[idx].name = "nop"
		}

		isInfinite, acc := isInfiniteLoop(newCommands)

		if !isInfinite {
			accumulator = acc
		}
	}

	return accumulator
}
