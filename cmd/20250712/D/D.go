package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(br, &N, &M)

	var points = make([]int, N)
	for i := range points {
		fmt.Fscan(br, &points[i])
	}

	// 真蠢了，直接合并最小的区间，直到规定的数量
	result := minTotalSpan(points, M)

	fmt.Println(result)
}

func minTotalSpan(points []int, k int) int {
	sort.Ints(points)

	cost := make([]int, 0, len(points))

	for i := range points {
		if i == 0 {
			continue
		}
		cost = append(cost, points[i]-points[i-1])
	}

	sort.Ints(cost)

	n := len(points)
	result := 0
	for n > k {
		result += cost[0]
		cost = cost[1:]
		n--
	}

	return result
}
