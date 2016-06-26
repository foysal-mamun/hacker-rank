package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	m, n := 0, 0
	fmt.Scan(&m)
	fmt.Scan(&n)

	arr := make([][]string, m)

	for i := 0; i < m; i++ {

		s := ""
		fmt.Scan(&s)
		arr[i] = strings.Split(s, "")

	}

	max := float64(m)
	if n > m {
		max = float64(n)
	}

	c1, c2, r1, r2 := "", "", "", ""
	f1, f2 := 0, 0
	for i := 0; i < int(math.Sqrt(max)); i++ {
		c1 = checkCol(i, i, arr)
		r1 = checkRow(i, i, arr)
		c2 = checkRCol(m-1-i, n-1-i, arr)
		r2 = checkRRow(m-1-i, n-1-i, arr)

		if c1 != "x" && c2 != "x" && r1 != "x" && r2 != "x" {
			f1, f2 = m-1-i, n-1-i
			break
		}
	}

	if r := f1*2 + f2*2; r > 0 {
		fmt.Println(r)
	} else {
		fmt.Println("impossible")
	}

}

func checkCol(x, y int, arr [][]string) string {

	l := len(arr)
	a := ""
	for i, j := x, y; i < l-x; i++ {
		//fmt.Println(i, j, arr[i][j])
		if arr[i][j] == "x" {
			a = "x"
			break
		}
	}
	return a
}

func checkRow(x, y int, arr [][]string) string {

	l := len(arr[x])
	a := ""
	for i, j := x, y; j < l-y; j++ {
		//fmt.Println(i, j, arr[i][j])
		if arr[i][j] == "x" {
			a = "x"
			break
		}
	}
	return a
}
func checkRCol(x, y int, arr [][]string) string {

	l := len(arr)
	a := ""
	//fmt.Println(x, y, l)
	for i, j := x, y; i >= l-x; i-- {
		//fmt.Println(i, j, arr[i][j], arr[i][j] == "x")
		if arr[i][j] == "x" {
			a = "x"
			break
		}
	}
	return a
}

func checkRRow(x, y int, arr [][]string) string {

	l := len(arr[x])
	a := ""
	//fmt.Println(x, y, l)
	for i, j := x, y; j >= l-y; j-- {
		//fmt.Println(i, j, arr[i][j])
		if arr[i][j] == "x" {
			a = "x"
			break
		}
	}
	return a
}
