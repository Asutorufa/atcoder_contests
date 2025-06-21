package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscanln(br, &N)

	buildsMap := map[int][]int{}

	readIntFunc(N, br, func(x int, index int) {
		buildsMap[x] = append(buildsMap[x], index)
	})

	maxL := 1
	for _, v := range buildsMap {
		n := len(v)
		if n < maxL {
			continue
		}

		if n == 1 {
			continue
		}

		if n == 2 && maxL < 2 {
			maxL = 2
			continue
		}

		dp := make([]map[int]int, n)
		for i := range dp {
			dp[i] = make(map[int]int, n)
		}

		for i := 0; i < n; i++ {
			for j := 1; j < n; j++ {
				sub := v[j] - v[i]

				prev := dp[i][sub]
				if prev == 0 {
					prev = 1
				}

				dp[j][sub] = prev + 1

				if dp[j][sub] > maxL {
					maxL = dp[j][sub]
				}
			}
		}
	}

	fmt.Println(maxL)
}

func readIntFunc(n int, r *bufio.Reader, f func(_ int, index int)) {
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(r, &x)
		f(x, i)
	}
}
