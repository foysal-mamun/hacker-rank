package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	a, numSwp := bubbleSwap(a, n)

	fmt.Printf("Array is sorted in %v swaps.\n", numSwp)
	fmt.Printf("First Element: %v\n", a[0])
	fmt.Printf("Last Element: %v", a[n-1])
}

func bubbleSwap(a []int, n int) ([]int, int) {

	numSwp := 0

	for i := 0; i < n; i++ {

		ns := 0

		for j := 0; j < n-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				ns++
			}
		}

		numSwp += ns

		if ns == 0 {
			break
		}
	}

	return a, numSwp
}
