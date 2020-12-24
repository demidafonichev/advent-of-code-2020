package main

import (
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open("input.txt")
	const n = 1000

	res := 0
	for i := 0; i < n; i++ {
		var imin, imax int
		var symbol byte
		var password string
		fmt.Fscanf(input, "%d-%d %c: %s", &imin, &imax, &symbol, &password)
		if (password[imin-1] == symbol) != (password[imax-1] == symbol) {
			res++
		}
	}

	fmt.Println(res)
}
