package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	N := readInt()

	ret := make([]int, N)
	for i := range N {
		ret[i] = solve()
	}

	for i := range N {
		fmt.Println(ret[i])
	}
}

func solve() int {
	var N, W int
	fmt.Fscan(br, &N, &W)

	M := 2 * W

	// 聚合到 mod 2W
	mods := make([]int, M)
	for i := range N {
		c := readInt()
		r := (i + 1) % M
		mods[r] += c
	}

	// 拉直成 4W，处理环
	arr := make([]int, 2*M)
	for i := range 2 * M {
		arr[i] = mods[i%M]
	}

	// 前缀和
	prefix := make([]int, 2*M+1)
	for i := range 2 * M {
		prefix[i+1] = prefix[i] + arr[i]
	}

	// 滑动窗口：起点 0..2W-1
	ret := math.MaxInt64
	for i := range M {
		sum := prefix[i+W] - prefix[i]
		if sum < ret {
			ret = sum
		}
	}

	return ret
}

var br = bufio.NewReader(os.Stdin)

func readInt() int {
	var x int
	fmt.Fscan(br, &x)
	return x
}

func readString() string {
	var x string
	fmt.Fscan(br, &x)
	return x
}
