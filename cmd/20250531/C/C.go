package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(br, &N, &M)

	// 范围 前缀和相加
	jh := make([]int, N+2)

	for i := 0; i < M; i++ {
		var x, y int
		fmt.Fscan(br, &x, &y)

		jh[x-1]++
		jh[y]--
	}

	min := math.MaxInt
	curr := 0
	for _, v := range jh[:N] {
		curr += v
		if curr < min {
			min = curr
		}
	}

	fmt.Println(min)
}
