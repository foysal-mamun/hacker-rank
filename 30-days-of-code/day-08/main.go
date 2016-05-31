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

	persons := []Person{}

	for i := 0; i < N; i++ {
		var name, phone string
		fmt.Scan(&name)
		fmt.Scan(&phone)

		persons = append(persons, Person{name, phone})
	}

	for i := 0; i < N; i++ {
		var name string
		fmt.Scan(&name)

		found := false

		for _, person := range persons {
			if person.name == name {
				fmt.Println(person.name + "=" + person.phone)
				found = true
			}
		}

		if !found {
			fmt.Println("Not found")
		}
	}

}
