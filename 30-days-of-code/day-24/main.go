package main

import (
	"fmt"

	"github.com/foysal-mamun/hacker-rank/30-days-of-code/day-24/linkedList"
)

func main() {
	N := 0
	fmt.Scan(&N)

	nums := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&nums[i])
	}

	ll := linkedList.NewNode(nil)
	for _, n := range nums {
		ll.Insert(linkedList.Int(n))
	}

	ll.RemoveDuplicates()
	fmt.Println(ll.Traverse())

}
