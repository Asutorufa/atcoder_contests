package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	type Pair struct {
		a, b int
	}

	AB := make([]Pair, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &AB[i].a, &AB[i].b)
	}

	// 排序规则：
	// a 升序
	// a 相同则 b 降序
	sort.Slice(AB, func(i, j int) bool {
		if AB[i].a == AB[j].a {
			return AB[i].b > AB[j].b
		}
		return AB[i].a < AB[j].a
	})

	dp := make([]int, 0)

	for _, p := range AB {
		b := p.b
		// 找到第一个大于等于 b 的元素
		// 如果不存在则添加，否则替换
		i := sort.SearchInts(dp, b)
		if i == len(dp) {
			dp = append(dp, b)
		} else {
			dp[i] = b
		}
	}

	fmt.Println(len(dp))
}
