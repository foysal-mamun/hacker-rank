package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)

	if 1 <= N && N <= 100000 {
		persons := map[string]string{}

		for i := 0; i < N; i++ {
			var name, phone string
			fmt.Scan(&name)
			fmt.Scan(&phone)

			persons[strings.ToLower(name)] = phone
		}

		for j := 0; j < N; j++ {
			var name string
			fmt.Scan(&name)

			if phone, ok := persons[strings.ToLower(name)]; ok {
				fmt.Println(name + "=" + phone)
			} else {
				fmt.Println("Not found")
			}
		}
	}

}
