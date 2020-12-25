package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	input, _ := os.Open("input.txt")

	values := []uint64{}

	// offset is the index of the element in values array we start searching for pair
	// step is length of the slice in which elements are searched
	offset := 0
	step := 25

	for i := 0; i < step; i++ {
		var value uint64
		fmt.Fscanf(input, "%d", &value)
		values = append(values, value)
	}

	// Loooking for desired value
	var desiredValue uint64
	for {
		var value uint64
		_, err := fmt.Fscanf(input, "%d", &value)
		if err == io.EOF {
			break
		}

		values = append(values, value)
		matchFound := false
		for i := offset; i < offset+step; i++ {
			for j := i + 1; j < offset+step; j++ {
				if values[i]+values[j] == value {
					matchFound = true
					break
				}
			}
			if matchFound {
				break
			}
		}
		if !matchFound {
			desiredValue = value
			break
		}

		offset++
	}

	// Starting from element with index i
	// we then trying to find the sequence with length from j and up
	// summing all items in values[i:i+j]
	// checking for oversum
	var sequenceItems []uint64
	for i := 0; i < len(values); i++ {
		solutionFound := false
		for j := 2; j < len(values)-i; j++ {
			oversum := false
			sequenceItems = make([]uint64, j)
			sum := uint64(0)
			for k := 0; k < j; k++ {
				sequenceItems[k] = values[i+k]
				sum += values[i+k]
				// If sum is already more than desired value
				// we continue from next i
				if sum > desiredValue {
					oversum = true
					break
				}
				if sum == desiredValue {
					solutionFound = true
				}
			}
			if oversum || solutionFound {
				break
			}
		}
		if solutionFound {
			break
		}
	}

	min := sequenceItems[0]
	max := sequenceItems[0]
	for _, value := range sequenceItems {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	fmt.Println(min + max)
}
