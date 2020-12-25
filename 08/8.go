package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	input, _ := os.Open("input.txt")

	result := 0
	var (
		instructionCommands []string
		instructionValues   []int
	)

	for {
		var command string
		var value int
		_, err := fmt.Fscanf(input, "%s %d", &command, &value)
		if err == io.EOF {
			break
		}

		instructionCommands = append(instructionCommands, command)
		instructionValues = append(instructionValues, value)
	}
	instructionsExecuted := make([]bool, len(instructionCommands))

	i := 0
	executeCommand := map[string]func(int){
		"acc": func(v int) {
			result += v
			i++
			return
		},
		"nop": func(v int) {
			i++
			return
		},
		"jmp": func(v int) {
			i += v
			return
		},
	}

	// Find instructions that create loop
	// Those instructions would be instructions from the beginning
	// to the instruction previous to one that repeats
	var instructionsInLoopIndexes []int
	for i < len(instructionCommands) {
		if instructionsExecuted[i] {
			break
		}
		instructionsExecuted[i] = true
		instructionsInLoopIndexes = append(instructionsInLoopIndexes, i)

		command := instructionCommands[i]
		value := instructionValues[i]
		executeCommand[command](value)
	}

	// Find the instruction to change and compute the result
	executionEnded := false
	// Trying to change each instruction with nop/jmp command
	for _, j := range instructionsInLoopIndexes {
		if instructionCommands[j] == "acc" {
			continue
		}

		// Making deep copy of the array to change instruction
		// (to keep initial array of instructions unchanged)
		instructionCommandsCopy := make([]string, len(instructionCommands))
		copy(instructionCommandsCopy, instructionCommands)

		// Changing the instruction command
		if instructionCommandsCopy[j] == "nop" {
			instructionCommandsCopy[j] = "jmp"
		} else if instructionCommandsCopy[j] == "jmp" {
			instructionCommandsCopy[j] = "nop"
		}

		// Executing instructions checking to loop
		i, result = 0, 0
		instructionsExecuted := make([]bool, len(instructionCommandsCopy))
		for i < len(instructionCommandsCopy) {
			if instructionsExecuted[i] {
				break
			}
			instructionsExecuted[i] = true

			command := instructionCommandsCopy[i]
			value := instructionValues[i]
			executeCommand[command](value)
			// Execution will end if we reach
			// index more than length of the array of the instructions
			if i >= len(instructionCommandsCopy) {
				executionEnded = true
			}
		}
		// Execution ended means loop is fixed
		// and result is computed
		if executionEnded {
			break
		}
	}

	fmt.Println(result)
}
