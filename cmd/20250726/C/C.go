package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, K, X int
	_, _ = fmt.Scan(&N, &K, &X)

	arrays := make([]string, 0, N)

	for i := 0; i < N; i++ {
		var A string
		_, _ = fmt.Scan(&A)

		arrays = append(arrays, A)
	}

	dfs(arrays, "", 0, K)

	sort.Strings(result)

	fmt.Println(result[X-1])
}

var result = make([]string, 0, 100000)

func dfs(arrays []string, str string, current int, max int) {
	if current == max {
		result = append(result, str)
		return
	}

	for i := range arrays {
		dfs(arrays, str+arrays[i], current+1, max)
	}
}
