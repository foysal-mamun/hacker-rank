package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)

	if 1 <= N && N <= 100000 {
		persons := map[string]string{}

		for i := 0; i < N; i++ {
			// var name, phone string
			name_phone, _, _ := bufio.NewReader(os.Stdin).ReadLine()
			// fmt.Scan(&name_phone)
			name_phones := strings.Fields(string(name_phone))
			// fmt.Println(name_phones)
			fmt.Println(len(name_phones))
			if len(name_phones) > 1 {
				persons[name_phones[0]] = name_phones[1]
			}

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
