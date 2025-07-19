package main

import "fmt"

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	cache := map[int]struct{}{}

	for i := 0; i < N; i++ {
		var A int
		_, _ = fmt.Scan(&A)

		cache[A] = struct{}{}
	}

	var X int
	_, _ = fmt.Scan(&X)

	if _, ok := cache[X]; ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
