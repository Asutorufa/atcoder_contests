package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(br, &n)

	preA := make([]int, n+1)
	preB := make([]int, n+1)
	preC := make([]int, n+1)

	for i := range n {
		var a int
		fmt.Fscan(br, &a)
		preA[i+1] = preA[i] + a
	}
	for i := range n {
		var b int
		fmt.Fscan(br, &b)
		preB[i+1] = preB[i] + b
	}
	for i := range n {
		var c int
		fmt.Fscan(br, &c)
		preC[i+1] = preC[i] + c
	}

	best := preA[1] - preB[1]
	ans := -1 << 60

	for j := 2; j <= n-1; j++ {

		/*
			sum := preA[i] + (preB[j] - preB[i]) + (preC[n] - preC[j])
			==> preA[i] - preB[i] + preB[j] - preC[j] + preC[n]
			==> (preA[i] - preB[i]) + (preB[j] - preC[j]) + preC[n]
			==> best + (preB[j] - preC[j]) + preC[n]
		*/
		ans = max(ans, best+(preB[j]-preC[j])+preC[n])
		best = max(best, preA[j]-preB[j])
	}

	fmt.Println(ans)
}
