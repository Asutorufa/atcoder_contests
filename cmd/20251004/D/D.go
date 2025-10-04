package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	var T int
	fmt.Fscan(br, &T)

	for i := 0; i < T; i++ {
		var N int
		var S string
		fmt.Fscan(br, &N, &S)
		fmt.Println(solve(S))
	}
}

// 正反都要来一遍，差一点，太可惜了
func solve(S string) int {
	ret, max := arrays(S)

	ans := 2 * len(S)
	for i := 0; i < len(ret); i++ {
		m := max[i]

		tmp := (ret[i]-m)*2 + ret[i^1]

		if tmp < ans {
			ans = tmp
		}
	}

	return ans
}

func arrays(S string) ([2]int, [2]int) {
	ret := []int{}
	last := '_'
	max := [2]int{}
	count := [2]int{}
	for i, v := range S {
		if v == '0' {
			count[0]++
		} else {
			count[1]++
		}
		if v == last {
			ret = append(ret, ret[i-1]+1)
		} else {
			ret = append(ret, 1)
		}

		if ret[i] > max[v-'0'] {
			max[v-'0'] = ret[i]
		}
		last = v
	}
	return count, max
}
