package main

import (
	"fmt"
)

func main() {
	var N, a int
	fmt.Scan(&N)

	var A []int

	for i := 0; i < N; i++ {
		fmt.Scan(&a)
		A = append(A, a)
	}

	for i := N - 1; i >= 0; i-- {
		fmt.Printf("%d ", A[i])
	}
	fmt.Println("")

}
