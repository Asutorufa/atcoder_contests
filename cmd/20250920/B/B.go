package main

import "fmt"

func main() {
	var N, M, K int
	fmt.Scan(&N, &M, &K)

	ret := make([]int, 0, N)
	maps := map[int]map[int]struct{}{}

	for i := 1; i <= N; i++ {
		maps[i] = map[int]struct{}{}

		for j := 1; j <= M; j++ {
			maps[i][j] = struct{}{}
		}
	}

	for i := 0; i < K; i++ {
		var a, b int
		fmt.Scan(&a, &b)

		delete(maps[a], b)

		if len(maps[a]) == 0 {
			if !exist(ret, a) {
				ret = append(ret, a)
			}
		}
	}

	for _, v := range ret {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func exist(m []int, v int) bool {
	for _, vv := range m {
		if vv == v {
			return true
		}
	}
	return false
}
