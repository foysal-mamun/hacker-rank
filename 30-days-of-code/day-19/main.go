package main

import (
	"fmt"
	"math"
)

type advancedArithmetic interface {
	DivisorSum(int) int
}

type Calculator struct {
}

func (c *Calculator) DivisorSum(n int) int {

	sumOfDiv := 0

	sqr := int(math.Sqrt(float64(n)))
	for i := 1; i <= sqr; i++ {
		if n%i == 0 {
			sumOfDiv += i

			if div := n / i; i != div {
				sumOfDiv += div
			}
		}

	}
	return sumOfDiv
}

func main() {
	n := 0
	fmt.Scan(&n)

	var a advancedArithmetic

	c := &Calculator{}
	a = c
	fmt.Println(a.DivisorSum(n))

}
