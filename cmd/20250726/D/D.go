package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	br := bufio.NewReaderSize(os.Stdin, 65535)

	var T int
	fmt.Fscan(br, &T)

	result := make([]int, 0, T)
	for i := 0; i < T; i++ {
		var N, M int
		fmt.Fscan(br, &N, &M)

		var A = make([]int, 0, N)
		var B = make([]int, 0, N)
		for i := 0; i < N; i++ {
			var NUM int
			fmt.Fscan(br, &NUM)

			A = append(A, NUM)
		}

		for i := 0; i < N; i++ {
			var NUM int
			fmt.Fscan(br, &NUM)

			B = append(B, NUM)
		}

		result = append(result, maxSumMod(A, B, M))
	}

	for _, r := range result {
		fmt.Println(r)
	}
}

/*
此处除法为编程中的除法，取整，而不是小数

**其次题目中有个重要条件 0 <= Ai, Bi < M, 可得 Ai + Bi 不可能大于两倍的M**

x mod M = x - M * (x / M)

(Ai + Bi) mod M = (Ai + Bi) - M * ( (Ai + Bi) / M )

if  (Ai + Bi) < M

	then (Ai + Bi) mod M = Ai + Bi

if  (Ai + Bi) >= M

	then (Ai + Bi) mod M = (Ai + Bi) - M

so sum((Ai + Bi) mod M) = sum ((Ai+Bi)- C * M)

所以根本就不用单个单个的匹配，直接贪心算法，找出所有Ai+Bi > M 的个数即可
*/
func maxSumMod(A, B []int, M int) int {
	idx := 0
	c := 0

	sort.Slice(A, func(i, j int) bool {
		return A[i] > A[j]
	})

	sort.Ints(B)

	sum := 0

	for _, v := range A {
		for idx < len(B) && B[idx]+v < M {
			idx++
		}

		if idx >= len(B) {
			break
		}

		c++
		idx++
	}

	for i, v := range A {
		sum += v
		sum += B[i]
	}

	return sum - c*M
}

/*
3
3 6
3 1 4
2 0 1
1 1000000000
999999999
999999999
10 201
144 150 176 154 110 187 38 136 111 46
96 109 73 63 85 1 156 7 13 171
*/
