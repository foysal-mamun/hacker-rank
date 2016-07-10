package primeNumber

import "math"

type PrimeNumber struct {
}

func New() *PrimeNumber {
	return &PrimeNumber{}
}

// Basic solution
func (p PrimeNumber) IsPrimeNumberV1(n int) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// Optimized little bit
func (p PrimeNumber) IsPrimeNumberV2(n int) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// 6k ± 1
func (p PrimeNumber) IsPrimeNumberV3(n int) bool {
	if n <= 1 {
		return false
	} else if n <= 3 {
		return true
	} else if n%2 == 0 && n%3 == 0 {
		return false
	}

	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}

	return true
}

func isExist(list []int, n int) bool {
	for _, v := range list {
		if v == n {
			return true
		}

		// Since it is sorted list
		if v > n {
			break
		}
	}

	return false
}

// func (p PrimeNumber) IsPrimeNumberV2(n int) bool {
// 	if n <= 1 {
// 		return true
// 	}

// 	sr := int(math.Sqrt(float64(n)))
// 	for i = 2; i < sr; i++ {
// 		if n % i {

// 		}
// 	}

// }
