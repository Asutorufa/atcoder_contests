package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	// 两个数两两结合的总数
	ans := M * (M - 1) / 2

	x := map[int]int{}

	for i := 0; i < M; i++ {
		var A, B int
		fmt.Scan(&A, &B)

		// (A+B)%N
		// 斜率相同的直线必定平行
		x[(A+B)%N] = x[(A+B)%N] + 1
	}

	for _, v := range x {
		// 用总数减去多余的平行线的数量
		ans -= v * (v - 1) / 2
	}

	fmt.Println(ans)
}
