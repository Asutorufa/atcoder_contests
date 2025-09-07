package main

import (
	"fmt"
)

/*D反而比较简单，基本思路就是直接平分，如果有余数再把余数平分到所有的上面，直到分完*/
func main() {
	var N, K int
	fmt.Scan(&N, &K)

	ans := []int{K}
	for n := 0; n < N; n++ {
		var nxt []int
		for _, a := range ans {
			nxt = append(nxt, a/2)
			nxt = append(nxt, a-a/2)
		}
		ans = nxt
	}

	maxVal, minVal := ans[0], ans[0]
	for _, a := range ans {
		if a > maxVal {
			maxVal = a
		}
		if a < minVal {
			minVal = a
		}
	}

	fmt.Println(maxVal - minVal)
	for i, a := range ans {
		if i+1 == len(ans) {
			fmt.Println(a)
		} else {
			fmt.Print(a, " ")
		}
	}
}
