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

	ret := make([]int, 0, T)
	for i := 0; i < T; i++ {
		var a, b, c int
		fmt.Fscan(br, &a, &b, &c)

		ret = append(ret, calc(a, b, c))
	}

	for i := range ret {
		fmt.Println(ret[i])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calc(a, b, c int) int {
	// 先尽量做 abc
	x := min(a, min(b, c))
	a -= x
	b -= x
	c -= x

	if b > 0 {
		// b 只能用在 abc，不可能继续
		return x
	}

	// 被坑了，这个地方不能单纯的a或c除以2再用剩余的计算，反而会变少
	// 可以这样考虑 aac+acc = (a+c)/3, 因为加起来 a和c都是三个
	return x + maxComb(a, c)
}

func maxComb(a, c int) int {
	sum := a + c
	// 最多数量受三个条件限制：
	// 1. 不能超过 a
	// 2. 不能超过 c
	// 3. 不能超过 (a+c)/3
	return min(min(a, c), sum/3)
}
