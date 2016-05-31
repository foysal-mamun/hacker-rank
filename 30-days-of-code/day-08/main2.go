package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	persons := map[string]string{}

	for i := 0; i < N; i++ {
		var name, phone string
		fmt.Scan(&name)
		fmt.Scan(&phone)

		persons[name] = phone
	}

	for i := 0; i < N; i++ {
		var name string
		fmt.Scan(&name)

		if phone, ok := persons[name]; ok {
			fmt.Println(name + "=" + phone)
		} else {
			fmt.Println("Not found")
		}
	}

}
