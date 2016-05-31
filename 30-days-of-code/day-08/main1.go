package main

import (
	"fmt"
)

type Person struct {
	name, phone string
}

func main() {
	var N int
	fmt.Scan(&N)

	persons := map[string]Person{}

	for i := 0; i < N; i++ {
		var name, phone string
		fmt.Scan(&name)
		fmt.Scan(&phone)

		persons[name] = Person{name, phone}
	}

	for i := 0; i < N; i++ {
		var name string
		fmt.Scan(&name)

		if person, ok := persons[name]; ok {
			fmt.Println(name + "=" + person.phone)
		} else {
			fmt.Println("Not found")
		}
	}

}
