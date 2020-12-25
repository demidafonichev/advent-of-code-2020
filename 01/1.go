package main

import (
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open("input.txt")
	const N = 200

	numbers := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscanf(input, "%d", &numbers[i])
	}

	for i, n := range numbers {
		for j, m := range numbers[i+1:] {
			for t, k := range numbers[i+j+2:] {
				if i != j && j != t && i != t && n+m+k == 2020 {
					fmt.Println(n * m * k)
					return
				}
			}
		}
	}
}
