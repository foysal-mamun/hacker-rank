package main

import (
	"fmt"
)

func main() {
	var T int
	var S string

	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&S)
		lenS := len(S)

		if 2 <= lenS && lenS <= 10000 {

			evenS, oddS := "", ""

			for j, s := range S {

				if j%2 != 0 {
					oddS += string(s)
				} else {
					evenS += string(s)
				}
			}

			fmt.Println(evenS, oddS)
		}
	}

}
