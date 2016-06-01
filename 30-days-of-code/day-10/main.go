package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	bi := getBinary(n)

	s := strings.Split(bi, "0")
	fmt.Println(getConsecutive(s))

}

func getConsecutive(s []string) int {
	max := 0
	for _, v := range s {
		if l := len(v); max < l {
			max = l
		}
	}

	return max
}

func getBinary(n int) string {
	if n <= 1 {
		return "1"
	}

	return getBinary(n/2) + strconv.Itoa(n%2)
}
