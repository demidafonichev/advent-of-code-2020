package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	input, _ := os.Open("input.txt")
	const n = 770
	seatIDs := make([]int, n)

	for i := 0; ; i++ {
		var line string
		_, err := fmt.Fscanln(input, &line)

		if err == io.EOF {
			break
		}

		rowMin, rowMax, colMin, colMax := 0, 127, 0, 7
		for _, symbol := range line {
			switch symbol {
			case 'F':
				rowMax -= (rowMax - rowMin + 1) / 2
			case 'B':
				rowMin += (rowMax - rowMin + 1) / 2
			case 'L':
				colMax -= (colMax - colMin + 1) / 2
			case 'R':
				colMin += (colMax - colMin + 1) / 2
			}
		}

		seatID := rowMin*8 + colMin
		seatIDs[i] = seatID
	}

	sort.Ints(seatIDs)
	for i := 0; i < n-1; i++ {
		if seatIDs[i]+1 != seatIDs[i+1] {
			fmt.Println(seatIDs[i] + 1)
			break
		}
	}
}
