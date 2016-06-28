package main

import "fmt"

func main() {
	T := 0
	fmt.Scan(&T)

	arr := make([]int, T)

	for i := 0; i < T; i++ {
		fmt.Scan(&arr[i])
	}

	for _, v := range arr {

		h := 1
		for i := 1; i <= v; i++ {
			if i%2 == 0 {
				h++
			} else {
				h *= 2
			}
		}
		fmt.Println(h)

	}

}
