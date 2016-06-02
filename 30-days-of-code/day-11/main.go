package main

import (
	"fmt"
)

func main() {

	A := [6][6]int{}

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Scan(&A[i][j])
		}
	}

	max := -70

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {

			// fmt.Println(A[i][j], A[i][j+1], A[i][j+2])
			// fmt.Println(" ", A[i+1][j+1], "")
			// fmt.Println(A[i+2][j], A[i+2][j+1], A[i+2][j+2])
			// fmt.Println("")

			a := A[i][j] + A[i][j+1] + A[i][j+2] + A[i+1][j+1] + A[i+2][j] + A[i+2][j+1] + A[i+2][j+2]
			fmt.Println(a, max)
			if max < a {
				max = a
			}
		}
	}

	fmt.Println(max)
}
