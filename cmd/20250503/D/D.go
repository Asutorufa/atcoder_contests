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

	ne = make([]int, N)

	for i := range ne {
		var C int
		fmt.Fscan(br, &C)

		ne[i] = C
	}

	de = make([][]int, N)
	remain := make([]int, M)

	for i := 0; i < M; i++ {
		var K int
		fmt.Fscan(br, &K)

		for j := 0; j < K; j++ {
			var A int
			fmt.Fscan(br, &A)

			de[A-1] = append(de[A-1], i)
		}
	}

	var minCost = math.MaxInt
	for pattern := 0; pattern < int(math.Pow(3, float64(N))); pattern++ {
		zooCount := make([]int, N)
		cost := 0
		tmp := pattern
		for i := 0; i < N; i++ {
			zooCount[i] = tmp % 3       // 每个动物园去几次，0,1,2
			cost += zooCount[i] * ne[i] // 总费用
			tmp /= 3                    // 去掉已经处理的一位
		}

		remain := make([]int, M)
		for i := 0; i < N; i++ {
			for _, animal := range de[i] {
				remain[animal] += zooCount[i]
			}
		}

		if allAccess(remain) { // 判断是否 可以访问所有动物园
			if minCost > cost {
				minCost = cost
			}
		}
	}

	fmt.Println(minCost)
	return

	/*
		下面是dfs解法，会超时
	*/
	var min = math.MaxInt
	history := make([]int, N)

	for i, v := range de {
		for _, v := range v {
			remain[v] += 1
		}

		history[i] += 1
		x := find(ne[i], remain, history)
		history[i] -= 1
		if min > x {
			min = x
		}

		for _, v := range v {
			remain[v] -= 1
		}
	}

	fmt.Println(min)
}

var de = make([][]int, 0)
var ne = make([]int, 0)

func find(now int, remain []int, history []int) int {
	if allAccess(remain) {
		return now
	}

	min := math.MaxInt

	for i, v := range de {
		if history[i] >= 2 {
			continue
		}

		for _, v := range v {
			remain[v] += 1
		}

		if allAccess(remain) {
			if min > now+ne[i] {
				min = now + ne[i]
			}
		} else {
			history[i] += 1
			x := find(now+ne[i], remain, history)
			history[i] -= 1
			if min > x {
				min = x
			}
		}

		for _, v := range v {
			remain[v] -= 1
		}
	}

	return min
}

func allAccess(remain []int) bool {
	for _, v := range remain {
		if v < 2 {
			return false
		}
	}

	return true
}
