package main

import (
	"fmt"
)

func factorial(N int) int {
	if N <= 1 {
		return 1
	}

	return N * factorial(N-1)
}

func main() {
	var N int
	fmt.Scan(&N)

	if 2 <= N && N <= 12 {
		fmt.Println(factorial(N))
	}
}
