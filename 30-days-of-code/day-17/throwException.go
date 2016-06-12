package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	T := 0
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		n, p := 0.0, 0.0
		fmt.Scan(&n)
		fmt.Scan(&p)

		cal := Calculator{}
		if p, err := cal.Power(n, p); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(p)
		}

	}
}

type Calculator struct {
}

func (c Calculator) Power(n, p float64) (float64, error) {

	if n < 0 || p < 0 {
		return 0, errors.New("n and p should be non-negative")
	}

	return math.Pow(n, p), nil
}
