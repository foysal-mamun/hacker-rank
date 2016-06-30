package main

import (
	"fmt"
	"reflect"
	"sort"
)

var mem = map[int]int{}
var set = map[int][]int{}

func main() {

	N := 1
	M := 0

	fmt.Scan(&N)
	fmt.Scan(&M)

	money := make([]int, N+1)
	roads := make([][]int, M)

	for i := 1; i <= N; i++ {
		fmt.Scan(&money[i])
	}
	for i := 0; i < M; i++ {
		r1, r2 := 0, 0
		fmt.Scan(&r1)
		fmt.Scan(&r2)
		roads[i] = []int{r1, r2}
	}

	// sort.Ints(money)
	maxMoneyWay := make(map[int]int)

	for i := 1; i <= N; i++ {
		mem = make(map[int]int)

		lenRoads := len(roads)
		for j := 0; j < lenRoads; j++ {

			r1, r2 := roads[j][0], roads[j][1]

			if i == r1 {
				mem[r2] = 1
			} else if i == r2 {
				mem[r1] = 1
			}
		}

		set[i] = append(set[i], money[i])
		maxMoneyWay[i] = getMoney(money, roads, i)
		//fmt.Println(mem)
		for k := 1; k < N+1; k++ {
			if _, ok := mem[k]; !ok {
				set[i] = append(set[i], money[k])
				maxMoneyWay[i] += getMoney(money, roads, k)
			}
		}

	}

	// i := 4
	// lenRoads := len(roads)
	// for j := 0; j < lenRoads; j++ {
	// 	r1, r2 := roads[j][0], roads[j][1]
	// 	if i == r1 {
	// 		mem[r2] = 1
	// 	} else if i == r2 {
	// 		mem[r1] = 1
	// 	}
	// }
	// maxMoneyWay[i] = getMoney(money, roads, i)

	maxMonye := 0
	maxMoneyIdx := []int{}

	for i, v := range maxMoneyWay {

		if maxMonye < v {
			maxMonye = v
			maxMoneyIdx = []int{}
			maxMoneyIdx = append(maxMoneyIdx, i)
		} else if maxMonye == v {
			maxMonye = v
			maxMoneyIdx = append(maxMoneyIdx, i)
		}

	}

	Max2Set := make(map[int][]int)
	for _, v := range maxMoneyIdx {
		Max2Set[v] = set[v]
	}

	Max2Set1 := make(map[int][]int)

	for i, v := range Max2Set {

		if len(Max2Set1) == 0 {
			Max2Set1[i] = v
			continue
		}

		notFound := false
		for _, p := range Max2Set1 {
			sort.Ints(v)
			sort.Ints(p)
			if !reflect.DeepEqual(v, p) {
				notFound = true
			} else {
				notFound = false
				break
			}

		}

		if notFound {
			Max2Set1[i] = v
		}

	}

	totalSet := len(Max2Set1)
	if totalSet == 0 {
		totalSet = 1
	}
	if maxMonye == 0 {
		totalSet = N + 2
	}

	// fmt.Println(maxMoneyWay)
	// fmt.Println(maxMoneyIdx)
	// fmt.Println(Max2Set)
	// fmt.Println()
	fmt.Println(maxMonye, totalSet)
	// fmt.Println(roads)
}

func getMoney(money []int, roads [][]int, curNode int) int {

	mem[curNode] = 1

	next := getNotConnectedNode(money, roads, curNode)
	if next == 0 {

		return money[curNode]
	}
	set[curNode] = append(set[curNode], money[next])
	return money[curNode] + getMoney(money, roads, next)
}

func getNotConnectedNode(money []int, roads [][]int, curNode int) int {

	lenRoads := len(roads)
	for j := 0; j < lenRoads; j++ {
		r1, r2 := roads[j][0], roads[j][1]

		if curNode == r1 {
			mem[r2] = 1
		} else if curNode == r2 {
			mem[r1] = 1
		}

		_, ok1 := mem[r1]
		_, ok2 := mem[r2]
		//fmt.Println(r1, money[r1], mem[r1], r2, money[r2], mem[r2])
		if !ok1 && !ok2 {
			if money[r1] > money[r1] {
				//fmt.Println("r1", r1)
				return r1
			} else {
				//fmt.Println("r1", r2)
				return r2
			}
		} else if !ok1 {
			//fmt.Println("r1", r1)
			return r1
		} else if !ok2 {
			//fmt.Println("r2", r2)
			return r2
		}

	}

	return 0

}
