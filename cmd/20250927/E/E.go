package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// 阶乘
func factorial(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

// 去重排列数
func uniquePermutationCount(nums []int) int {
	countMap := make(map[int]int)
	for _, v := range nums {
		countMap[v]++
	}

	n := len(nums)
	res := factorial(n)
	for _, c := range countMap {
		res /= factorial(c)
	}
	return res
}

func main() {
	br := bufio.NewReader(os.Stdin)

	var T, M int
	fmt.Fscan(br, &T, &M)

	for i := 0; i < T; i++ {
		var N int
		fmt.Fscan(br, &N)

		array := make([]int, 0, N)
		for j := 0; j < N; j++ {
			var C int
			fmt.Fscan(br, &C)

			for k := 0; k < C; k++ {
				array = append(array, j+1)
			}
		}

		count := uniquePermutationCount(array)

		if count < 0 {
			count += math.MaxInt
		}

		fmt.Println(count % M)
	}
}

/*
3 998244353
1
1
3
4 2 5
*/
