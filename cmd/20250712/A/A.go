package main

import "fmt"

func main() {
	var N, L, R int
	fmt.Scan(&N, &L, &R)

	count := 0
	for i := 0; i < N; i++ {
		var X, Y int
		fmt.Scan(&X, &Y)

		if X <= L && Y >= R {
			count++
		}
	}

	fmt.Println(count)
}
