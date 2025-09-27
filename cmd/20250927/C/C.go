package main

import (
	"bufio"
	"fmt"
	"os"
)

type CircularPrefixSum struct {
	n      int
	prefix []int
	offset int
}

func NewCircularPrefixSum(a []int) *CircularPrefixSum {
	n := len(a)
	b := make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		b[i] = a[i%n]
	}
	prefix := make([]int, 2*n+1)
	for i := 0; i < 2*n; i++ {
		prefix[i+1] = prefix[i] + b[i]
	}
	return &CircularPrefixSum{n: n, prefix: prefix, offset: 0}
}

// rotateLeft 把第一个元素移到最后，相当于偏移量+1
func (c *CircularPrefixSum) rotateLeft(i int) {
	c.offset = (c.offset + i) % c.n
}

// rangeSum 计算当前旋转数组中 [l, r] 区间的和 (0-indexed, inclusive)
func (c *CircularPrefixSum) rangeSum(l, r int) int {
	l += c.offset
	r += c.offset
	return c.prefix[r+1] - c.prefix[l]
}

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, Q int
	fmt.Fscan(br, &N, &Q)

	As := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(br, &As[i])
	}

	c := NewCircularPrefixSum(As)

	ans := make([]int, 0, Q)
	for i := 0; i < Q; i++ {
		var T int
		fmt.Fscan(br, &T)

		if T == 1 {
			var X int
			fmt.Fscan(br, &X)
			c.rotateLeft(X)
		} else {
			var L, R int
			fmt.Fscan(br, &L, &R)
			ans = append(ans, c.rangeSum(L-1, R-1))
		}
	}

	for _, v := range ans {
		fmt.Println(v)
	}
}
