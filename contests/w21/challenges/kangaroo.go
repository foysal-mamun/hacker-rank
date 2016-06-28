package main

import "fmt"

func main() {
	x1, v1, x2, v2 := 0, 0, 0, 0
	fmt.Scan(&x1)
	fmt.Scan(&v1)
	fmt.Scan(&x2)
	fmt.Scan(&v2)

	x := x1 - x2
	v := v2 - v1
	if x%v == 0 && x/v >= 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
