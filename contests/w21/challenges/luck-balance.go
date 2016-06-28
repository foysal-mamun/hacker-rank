package main

import (
	"fmt"
	"sort"
)

func main() {

	N, K := 1, 0
	fmt.Scan(&N)
	fmt.Scan(&K)

	if 1 <= N && N <= 100 && 0 <= K && K <= N {

		notImpTask, veryImpTask := []int{}, []int{}

		for i := 1; i <= N; i++ {
			t, r := 1, 0
			fmt.Scan(&t)
			fmt.Scan(&r)

			if r == 0 {
				notImpTask = append(notImpTask, t)
			} else {
				veryImpTask = append(veryImpTask, t)
			}
		}

		luckBalance := 0
		for _, v := range notImpTask {
			luckBalance += v
		}

		sort.Ints(veryImpTask)
		l := len(veryImpTask)

		for i := l - 1; i >= 0; i-- {
			m := -1
			if K > 0 {
				m = 1
				K--
			}

			luckBalance = luckBalance + (veryImpTask[i] * m)
		}

		fmt.Println(luckBalance)

	}

}
