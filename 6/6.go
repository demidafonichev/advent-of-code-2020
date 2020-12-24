package main

import (
	"fmt"
	"io"
	"os"

	"6/set"
)

func main() {
	input, _ := os.Open("input.txt")
	countPositive := 0

	fileEnded := false
	for !fileEnded {

		groupAllPositiveAnswers := set.SetFromString("abcdefghijklmnopqrstuvwxyz")
		for {
			var personAnswers string
			_, err := fmt.Fscanln(input, &personAnswers)
			// Check for end of file first
			if err == io.EOF {
				fileEnded = true
				break
			}
			// Check for end of person answers
			if personAnswers == "" {
				break
			}
			// Collect group positive answers
			// without duplicates with Intersection
			personPositiveAnswers := set.SetFromString(personAnswers)
			groupAllPositiveAnswers = groupAllPositiveAnswers.Intersection(personPositiveAnswers)
		}

		countPositive += groupAllPositiveAnswers.Size()
	}

	fmt.Println(countPositive)
}
