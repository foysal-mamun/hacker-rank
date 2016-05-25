package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	n, _ := strconv.ParseInt(string(input[:]), 0, 0)

	if 1 <= n && n <= 100 {

		if n%2 != 0 || (6 <= n && n <= 20) {
			fmt.Println("Weird")
		} else if (2 <= n && n <= 5) || n > 20 {
			fmt.Println("Not Weird")
		}
	}

}
