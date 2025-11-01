package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, A, B int
	fmt.Fscan(br, &N, &A, &B)

	var S string
	fmt.Fscan(br, &S)

	fmt.Println(countSubstrings(S, A, B))
}

func countSubstrings(S string, A, B int) int {
	// 前缀和
	n := len(S)

	prefixa := make([]int, n+1)
	prefixb := make([]int, n+1)

	for i := 1; i <= n; i++ {
		prefixa[i] = prefixa[i-1]
		prefixb[i] = prefixb[i-1]

		if S[i-1] == 'a' {
			prefixa[i]++
		} else {
			prefixb[i]++
		}
	}

	var ans int

	r1, r2 := 1, 1

	// index每次右移一个，遍历N次
	for l := 1; l <= n; l++ {

		r1 = max(r1, l)
		r2 = max(r2, l)

		// 找到第一个满足A数量的位置
		for r1 <= n && prefixa[r1]-prefixa[l-1] < A {
			r1++
		}
		// 找到第一个满足B数量的位置
		for r2 <= n && prefixb[r2]-prefixb[l-1] < B {
			r2++
		}

		// 如果B的位置大于A，那么中间的数字都是多余的，所以直接计算
		if r2 > r1 {
			ans += r2 - r1
		}
	}

	return ans
}
