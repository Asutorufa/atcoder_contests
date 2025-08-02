package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Gift struct {
	p, a, b uint
}

func maxPA(gifts []Gift) uint {
	max := uint(0)
	for _, g := range gifts {
		if g.p+g.a > max {
			max = g.p + g.a
		}
	}
	return max
}

func min(x, y uint) uint {
	if x < y {
		return x
	}
	return y
}

func main() {
	sc := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(sc, &N)

	// 构造 b 的前缀和
	// 所有下降心情的前缀和
	sumB := make([]uint, N)
	gifts := make([]Gift, N)
	for i := range gifts {
		var P, A, B uint
		fmt.Fscan(sc, &P, &A, &B)
		gifts[i] = Gift{P, A, B}

		if i == 0 {
			sumB[i] = B
		} else {
			sumB[i] = sumB[i-1] + B
		}
	}

	// 预估dp范围
	// 当前为： 礼物价值+增加值 最大的那个
	M := maxPA(gifts)
	if M > 1000 {
		M = 1000
	}

	// dp[i][j]: i个礼物后，当前心情为j，最终心情是多少
	dp := make([][]uint, N+1)
	for i := range dp {
		dp[i] = make([]uint, M+1)
	}
	// 终止状态
	for j := range dp[N] {
		dp[N][j] = uint(j)
	}

	// 从后往前做DP
	// 范围是 0 - M
	for i := int(N) - 1; i >= 0; i-- {
		p, a, b := gifts[i].p, gifts[i].a, gifts[i].b
		for j := uint(0); j <= M; j++ {
			if j <= p {
				dp[i][j] = dp[i+1][j+a]
			} else {
				dp[i][j] = dp[i+1][j-min(j, b)]
			}
		}
	}

	// access 函数：给一个初始心情值 x，返回最终心情值
	access := func(x uint) uint {
		if x <= M {
			return dp[0][x]
		}
		// 处理超出dp范围的心情
		//
		// 找到能把心情降到dp范围内的那个前缀和
		//
		// ! 此处必须二分查找， 不然会超时， 没有任何容错度
		idx := lowerBound(sumB, x-M)
		// 如果范围还是超出dp， 直接返回 当前心情减去所有前缀和
		if idx == -1 || idx >= N {
			return x - sumB[N-1]
		}

		return dp[idx+1][x-min(x, sumB[idx])]
	}

	var Q uint
	fmt.Fscan(sc, &Q)

	for i := uint(0); i < Q; i++ {
		var P uint
		fmt.Fscan(sc, &P)

		fmt.Println(access(P))
	}
}

func lowerBound(a []uint, x uint) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
