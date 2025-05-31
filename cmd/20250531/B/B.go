package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	set := map[int]struct{}{}
	array := make([]int, 0, N)

	for i := 1; i <= N; i++ {
		var A int
		fmt.Scan(&A)

		if _, ok := set[A]; !ok {
			array = append(array, A)
			set[A] = struct{}{}
		}
	}

	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})

	fmt.Println(len(array))
	for i, v := range array {
		fmt.Print(v)
		if i != len(array)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
