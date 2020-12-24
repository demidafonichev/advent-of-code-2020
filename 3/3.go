package main

import (
	"fmt"
	"math"
	"os"
)

func CountTreesOnPath(field [][]bool, i, j, down, right int) (res int) {
	if i == len(field)-1 {
		if field[i][j] {
			res = 1
		} else {
			res = 0
		}
	} else {
		if field[i][j] {
			res = CountTreesOnPath(field, i+down, j+right, down, right) + 1
		} else {
			res = CountTreesOnPath(field, i+down, j+right, down, right)
		}
	}
	return
}

func main() {
	const n = 323
	const m = 2261

	var repeats = int(math.Ceil(m / 31.0)) // width of pattern
	field := make([][]bool, n)

	input, _ := os.Open("input.txt")
	for i := 0; i < n; i++ {
		var line string
		fmt.Fscanf(input, "%s", &line)
		field[i] = make([]bool, m)
		for r := 0; r < repeats; r++ {
			l := len(line)
			for j := 0; j < l && j+r*l < m; j++ {
				if line[j] == '#' {
					field[i][j+r*l] = true
				}
			}
		}
	}

	res1 := CountTreesOnPath(field, 0, 0, 1, 1)
	res2 := CountTreesOnPath(field, 0, 0, 1, 3)
	res3 := CountTreesOnPath(field, 0, 0, 1, 5)
	res4 := CountTreesOnPath(field, 0, 0, 1, 7)
	res5 := CountTreesOnPath(field, 0, 0, 2, 1)

	fmt.Println(res1 * res2 * res3 * res4 * res5)
}
