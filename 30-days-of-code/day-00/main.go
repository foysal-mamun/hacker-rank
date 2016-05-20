package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputString, _, _ := bufio.NewReader(os.Stdin).ReadLine()

	fmt.Println("Hello, World.")
	fmt.Println(string(inputString[:]))
}
