package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

// Save calculated values
// to make less reqursive calls
var calculatedArrangements []int

func getArrangements(arr []int, i int) int {
	// If value was already calculated - return it
	if calculatedArrangements[i] != 0 {
		return calculatedArrangements[i]
	}

	res := 0
	if i == len(arr)-1 || i == len(arr)-2 || i == len(arr)-3 {
		res = 1
	} else if arr[i+1]-arr[i] < 3 && arr[i+2]-arr[i] < 3 && arr[i+3]-arr[i] == 3 {
		res = getArrangements(arr, i+1) + getArrangements(arr, i+2) + getArrangements(arr, i+3)
	} else if arr[i+1]-arr[i] < 3 && arr[i+3]-arr[i] == 3 {
		res = getArrangements(arr, i+1) + getArrangements(arr, i+3)
	} else if arr[i+2]-arr[i] < 3 && arr[i+3]-arr[i] == 3 {
		res = getArrangements(arr, i+2) + getArrangements(arr, i+3)
	} else if arr[i+1]-arr[i] < 3 && arr[i+2]-arr[i] <= 3 {
		res = getArrangements(arr, i+1) + getArrangements(arr, i+2)
	} else {
		res = getArrangements(arr, i+1)
	}
	// Save calculated value
	calculatedArrangements[i] = res
	return res
}

func main() {
	input, _ := os.Open("input.txt")

	// Init array with [0]
	// for the case where first input jolts would be like 1, 2, 3
	jolts := []int{0}
	for {
		var jolt int
		_, err := fmt.Fscanf(input, "%d", &jolt)
		if err == io.EOF {
			break
		}
		jolts = append(jolts, jolt)
	}

	sort.Ints(jolts)
	// Append last jolt value from puzzle definition
	jolts = append(jolts, jolts[len(jolts)-1]+3)
	// Allocate array to save calculated values
	calculatedArrangements = make([]int, len(jolts))

	// Recursively calculate number of the arrangements
	res := getArrangements(jolts, 0)
	fmt.Println(res)
}
