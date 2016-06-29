package main

import (
	"fmt"
	"sort"
)

func main() {
	N := 0
	fmt.Scan(&N)

	P := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&P[i])
	}

	if sort.IntsAreSorted(P) {
		fmt.Printf("%.6f\n", 0.0)
	} else {

		dupFreArr := dupCount(P)
		factSum := 1
		for _, v := range dupFreArr {
			factSum *= factorial(v)
		}

		fmt.Printf("%.6f\n", float64(factorial(N)/factSum))

	}
}

func dupCount(arr []int) []int {

	dupFre := make(map[int]int)

	for _, v := range arr {
		if _, ex := dupFre[v]; ex {
			dupFre[v] += 1
		} else {
			dupFre[v] = 1
		}
	}

	dupFreArr := []int{}
	for _, v := range dupFre {
		if v != 1 {
			dupFreArr = append(dupFreArr, v)
		}
	}

	return dupFreArr

}

func factorial(n int) int {
	if n <= 0 {
		return 1
	}

	return n * factorial(n-1)
}
