package main

import "fmt"

func main() {
	intArray := []int{1, 2, 3}
	stringArray := []string{"Hello", "World"}

	vals := make([]interface{}, len(intArray))
	for i, v := range intArray {
		vals[i] = v
	}
	printArray(vals)

	vals = make([]interface{}, len(stringArray))
	for i, v := range stringArray {
		vals[i] = v
	}
	printArray(vals)
}

func printArray(a []interface{}) {
	for _, element := range a {
		fmt.Println(element)
	}
}
