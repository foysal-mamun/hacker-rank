package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := ""
	fmt.Scan(&s)

	fmt.Println(ConvertStoI(s))
}

func ConvertStoI(s string) interface{} {
	i, err := strconv.Atoi(s)
	if err != nil {
		return "Bad String"
	}

	return i
}
