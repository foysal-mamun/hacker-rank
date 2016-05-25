package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	if 2 <= N && N <= 20 {
		for i := 1; i <= 10; i++ {
			fmt.Printf("%d x %d = %d\n", N, i, N*i)
		}
	}
}
