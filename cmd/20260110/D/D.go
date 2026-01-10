package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func solve(X int, A []int, Y int) int {
	n := len(A)

	start := sort.Search(n, func(i int) bool {
		return A[i] >= X
	})

	// X 没有在数组中时
	if start == n {
		return X + Y - 1
	}

	idx := sort.Search(n-start,
		func(k int) bool {
			// k 表示第k个数(也可以说是已经被占的数量)，因为是已经排序的数组了
			// A[start+k]-X 表示从 X 到 A[start+k] 中间有的位置，包括空位和已经被占的数
			// 所以 位置 - 已有的k个数 ==> A[start+k]-X-k 就是空位
			//
			return (A[start+k]-X)-k >= Y
		})

	// Y 不在数组的空位中时
	if idx+start == n {
		last := n - 1
		// last - start 表示数组中已有的数量
		// A[last] - X 表示从 X 到 A[last] 中间有的位置，包括空位和已经被占的数
		// 数组中的空位=总位置-已经被占的
		totalGaps := (A[last] - X) - (last - start)
		// Y - totalGaps 表示超出数组的空位
		// 所以答案是数组的最后一个数+超出数组的空位
		return A[last] + (Y - totalGaps)
	}

	// Y 在数组的空位中时
	// idx 表示已经被占的数量，因为是数组且已排序
	return X + Y + idx - 1
}

func main() {
	var N, Q int
	fmt.Fscan(br, &N, &Q)

	A := make([]int, 0, N)
	for i := 0; i < N; i++ {
		A = append(A, readInt())
	}

	sort.Ints(A)

	ret := make([]int, Q)

	for i := range Q {
		ret[i] = solve(readInt(), A, readInt())
	}

	for i := range Q {
		fmt.Println(ret[i])
	}
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
