package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	arrays := make([]int, 0, N)

	for i := 0; i < N; i++ {
		var A int
		fmt.Scan(&A)

		arrays = append(arrays, A)
	}

	deleted := map[int]bool{}

	for i := 0; i < M; i++ {
		var B int
		fmt.Scan(&B)

		findIndex(arrays, B, deleted)
	}

	for i, v := range arrays {
		if !deleted[i] {
			fmt.Print(v, " ")
		}
	}

	fmt.Print("\n")
}

func findIndex(arrays []int, target int, deleted map[int]bool) {
	for i, v := range arrays {
		if v == target && !deleted[i] {
			deleted[i] = true
			return
		}
	}
}
