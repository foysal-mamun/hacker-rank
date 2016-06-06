package main

import (
	"fmt"
	"sort"
)

func main() {
	N := 0
	fmt.Scan(&N)

	a := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
	}

	sort.Ints(a)

	d := NewDifference(a)
	d.ComputeDifference()
	fmt.Println(d.maximumDifference)
}

type Difference struct {
	elements          []int
	maximumDifference int
}

func NewDifference(elements []int) Difference {
	d := Difference{}
	d.elements = elements

	return d
}

func (d *Difference) ComputeDifference() {
	d.maximumDifference = d.elements[len(d.elements)-1] - d.elements[0]
}
