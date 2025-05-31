package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var T int
	fmt.Fscanln(br, &T)

	for i := 0; i < T; i++ {
		var N int
		fmt.Fscanln(br, &N)

		var S string = readString(br, N)

		solve(S)
	}

}

func readString(br *bufio.Reader, n int) string {
	var str []byte
	for i := 0; i < n; i++ {
		b, _ := br.ReadByte()
		str = append(str, b)
	}
	br.ReadByte()
	return string(str)
}

func solve(str string) int {
	sum := 0
	c := make([]int, len(str)+1)
	for i, v := range str {
		if v == '0' {
			c[i+1] = c[i] + 1
		} else {
			sum++
			c[i+1] = c[i] - 1
		}
	}

	/*
		此处为 Kadane 算法的变种
			Kadane 算法 是一种非常高效的动态规划算法，专门用于求解一维数组中 最大子数组和（Maximum Subarray Sum）的问题，时间复杂度为 O(n)。
	*/
	ma := 0
	res := 0 // In C++, res is initialized to 0.
	for i := 0; i <= len(str); i++ {
		// 找的是这个“跌落最深”的地方，也就是最小的 c[i] - ma
		res = min(res, c[i]-ma)
		ma = max(ma, c[i])
	}
	/*
			例如： 11111111001
			前缀数组为: [0 -1 -2 -3 -4 -5 -6 -7 -8 -7 -6 -7]

			res最小为 -8, 但1的总数却有9个, 所以 -8+9 = 1


			例如： 1111111111
			前缀数组为： [0 -1 -2 -3 -4 -5 -6 -7 -8 -9 -10]

			res最小为 -10, 1的总数有10个, 所以 -10+10 = 0

			例如： 0000000
			前缀数组为： [0 1 2 3 4 5 6 7]
			res最小为 0（根本没有0,不会加1）, 1的总数有0个, 所以 0+0 = 0

			例如： 010101
			前缀数组为： [0 1 0 1 0 1 0]
			res最小为 -1, 1的总数有3个, 所以 -1+3 = 2

			例如： 01010101
			前缀数组为： [0 1 0 1 0 1 0 1 0]
			res最小为 -1, 1的总数有4个, 所以 -1+4 = 3

			例如： 0110110
			前缀数组为： [0 1 0 -1 0 -1 -2 -1]
			res = min(res, c[i]-ma)
			ma = max(ma, c[i])

			ma = 0, res = 0
			index = 0 [0], res = min(0, 0 - 0) = 0, ma = max(0, 0) = 0
			index = 1 [1], res = min(0, 1 - 0) = 0, ma = max(0, 1) = 1
			index = 2 [0], res = min(0, 0 - 1) = -1, ma = max(1, 0) = 1
			index = 3 [-1], res = min(-1, - 1 - 1) = -2, ma = max(1, 1) = 1
			index = 4 [0], res = min(-2, 0 - 1) = -2, ma = max(1, 0) = 1
			index = 5 [-1], res = min(-2, - 1 - 1) = -2, ma = max(1, 1) = 1
			index = 6 [-2], res = min(-2, - 2 - 1) = -3, ma = max(1, 1) = 1
			index = 7 [-1], res = min(-3, - 1 - 1) = -3, ma = max(1, 1) = 1

			res最小为 -3, 1的总数有4个, 所以 -3+4 = 1


		 由上可知，res只会越来越小，ma只会越来越大
		 而1的数量越多，res越大形成连续的1的代价就越大，res越小形成连续1的代价越小，当与sum完全相反时则没有任何代价
	*/

	fmt.Println(sum + res)

	return sum + res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
