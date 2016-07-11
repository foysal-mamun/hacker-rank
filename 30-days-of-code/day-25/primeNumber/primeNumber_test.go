package primeNumber

import "testing"

func TestIsPrimeNumberV1(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149}
	notPrimes := []int{0, 1, 4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20, 21, 22, 24, 25, 26, 27, 28, 30, 32, 33, 34, 35, 36, 38, 39, 40, 42, 44, 45, 46, 48, 49, 50, 51, 52, 54, 55, 56, 57, 58, 60, 62, 63}
	pn := New()
	for _, i := range primes {
		if !pn.IsPrimeNumberV1(i) {
			t.Errorf("%d is not prime number", i)
		}
	}
	for _, i := range notPrimes {
		if pn.IsPrimeNumberV1(i) {
			t.Errorf("%d is prime number, but we are tesing non prime number", i)
		}
	}

	// t.Errorf("test %v", pn.IsPrimeNumberV3(2147483647))
}

func TestIsPrimeNumberV2(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149}
	notPrimes := []int{0, 1, 4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20, 21, 22, 24, 25, 26, 27, 28, 30, 32, 33, 34, 35, 36, 38, 39, 40, 42, 44, 45, 46, 48, 49, 50, 51, 52, 54, 55, 56, 57, 58, 60, 62, 63}
	pn := New()
	for _, i := range primes {
		if !pn.IsPrimeNumberV2(i) {
			t.Errorf("%d is not prime number", i)
		}
	}
	for _, i := range notPrimes {
		if pn.IsPrimeNumberV2(i) {
			t.Errorf("%d is prime number, but we are tesing non prime number", i)
		}
	}
}

func TestIsPrimeNumberV3(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149}
	notPrimes := []int{0, 1, 4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20, 21, 22, 24, 25, 26, 27, 28, 30, 32, 33, 34, 35, 36, 38, 39, 40, 42, 44, 45, 46, 48, 49, 50, 51, 52, 54, 55, 56, 57, 58, 60, 62, 63}
	pn := New()
	for _, i := range primes {
		if !pn.IsPrimeNumberV2(i) {
			t.Errorf("%d is not prime number", i)
		}
	}
	for _, i := range notPrimes {
		if pn.IsPrimeNumberV2(i) {
			t.Errorf("%d is prime number, but we are tesing non prime number", i)
		}
	}
}

func BenchmarkIsPrimeNumberV1(b *testing.B) {
	pn := New()
	pn.IsPrimeNumberV1(2147483647)
}

func BenchmarkIsPrimeNumberV2(b *testing.B) {
	pn := New()
	pn.IsPrimeNumberV1(2147483647)
}

func BenchmarkIsPrimeNumberV3(b *testing.B) {
	pn := New()
	pn.IsPrimeNumberV3(2147483647)
}
